# Elegance of Go's error handling

&emsp;每隔一段时间，各个论坛上都会出现一些关于`go`语言错误的帖子，并且针对`go`语言的错误处理每个人都有自己的看法。有些人说它们应该更像`throwable exception`（译者注：该观点类似Java语言中`throw exception`异常处理，发表该观点的人推测有Java开发背景），另外一些人则更喜欢`sum types`(可参考[1](https://lispcast.com/what-are-product-and-sum-types/).[2](https://en.wikipedia.org/wiki/Tagged_union).)模式的错误处理，比如类似`rust`语言中的Result<T, E>错误处理模式。虽然我在[typecript中的错误处理方式](https://dev.to/duunitori/mimicing-rust-s-result-type-in-typescript-3pn1)采用了`sum types`的方法，但我还是喜欢`go`处理错误的方式。  

&emsp;话虽如此，要真正弄清楚如何正确的处理错误可能还是需要一些时间（无论使用`sum type`模式还是`exception`模式）。在这篇文章中，我将介绍`go`语言中`http.Handler`中的错误处理方法`http.Handler`.

## 典型示例
&emsp;如果你期望`error`在不做任何处理的情况下以原始形式直接返回，那么直接返回的结果`error`可能让人沮丧。通常的示例是这样的。
```
func copyfile(src, dst) error {
	fsrc, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fsrc.Close()

	fdst, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fdst.Close()

	err := io.Copy(fdst, fsrc)

	return err
}
```
&emsp;我很肯定你们大多数人都见过这样的例子，而且这种错误处理还算好的。现在我们来看看可能会出现在`http.Hanlder`中的类似的情况。

```
func handleThing(w http.ResponseWriter, r *http.Request) {
	// Our path is something like /thing/3
	id, err := idFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusNotFound)
		return
	}

	thing, err := store.GetThingByID(id)
	if err != nil {
		// The error might be sql.NoRows, or it might be something else.
		if store.IsNotFoundErr(err) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		log.Printf("Failed to get a thing: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	acc := AccountFromRequest(r)
	if acc == nil {
		// No account attached to the request's session -> permission denied.
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	has, err := thing.HasPermissionToView(acc)
	if err != nil {
		// For some reason, we failed to check permissions. Better log it.
		log.Printf("Failed to check permissions: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if !has {
		// Permission denied.
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	// All good, send data to the client.
	respond(w, r, decodeThing(thing))
}
```
&emsp;上面代码中出现的函数功能定义可能需要需要做一些想象，但函数要实现的功能有以下几点。
- 从URL中提取ID
- 使用该ID进行数据库查询
- 检查客户端是否有权限查看该响应结果
- 返回结果给客户端
&emsp;这个功能可能会在其他资源类型上重复，所以很快就会重复。想象一下，如果对foo、bar、account等资源也做同样的事情，那该多好啊。同样的功能可以在Django中这样写。

```
def handle_thing(request):
    id = id_or_bad_request(request)
    thing = thing_or_404(id)

    account = account_or_forbidden(request)

    if not thing.has_permission(account):
        raise Forbidden()

    return JsonResponse(...)
```
&emsp;id_or_bad_request, thing_or_404和account_or_forbidden都会抛出一个错误，有人会在更高的地方发现并做相应的事情，比如用正确的状态代码来响应，并记录任何错误。

## Trying to simplify it
&emsp;牢记这段python代码，让我们想一想，我们可以在go代码中做些什么，让它更简洁一些。

- 当一个错误发生时，我们只想把它 "扔 "到某个地方。通常是客户端错误，但不一定。也许有人能弄明白？
- 如果一个非客户端错误发生，它需要被记录在某个地方
- 应该还有人能够弄清楚http.Error的调用。

&emsp;Golang的错误处理和Go谈了一下http处理程序中的错误处理，并给出了下面的例子。
```
func viewRecord(w http.ResponseWriter, r *http.Request) error {
    c := appengine.NewContext(r)
    key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
    record := new(Record)
    if err := datastore.Get(c, key, record); err != nil {
        return err
    }
    return viewTemplate.Execute(w, record)
}

// NOTE: the following is my adapted version from the example's ServeHTTP to a
// middleware/wrapper

type HandlerE = func(w http.ResponseWriter, r *http.Request) error

func WithError(h HandlerE) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}
```
&emsp;这已经解决了我们的一个问题，那就是http.Error调用。但有时我们不想把详细的错误暴露给客户端，所以我会用一个通用的内部服务器错误信息来代替实际信息。另外，记录服务器内部错误的原因也很重要，这样你就可以弄清楚出了什么问题。

## Improving the WithError wrapper
&emsp;我们希望从实际的 http.Handler 中返回一个错误，但以某种方式指示 WithError 封装函数在客户端收到错误时正确响应，并在一些错误上记录错误。就像这样。
```
func WithError(h HandlerE) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {

			if is404err(err) {
				http.Error(w, "not found", 404)
				return
			}

			if isBadRequest(err) {
				http.Error(w, "bad request", 400)
				return
			}

			// Some other special cases...
			// ...

			log.Printf("Something went wrong: %v", err)

			http.Error(w, "Internal server error", 500)
		}
	}
}
```
&emsp;嗯，那些 "其他特殊 "的情况可能会来来去去，而且对于一些处理程序来说可能会变得很特殊。另外，我们还需要写那些is404err和isBadRequest处理程序，以及后续的什么东西。我们可以用一个接口做得更好。
```
type ErrorResponder interface {
    // RespondError writes an error message to w. If it doesn't know what to
    // respond, it returns false.
	RespondError(w http.ResponseWriter, r *http.Request) bool
}
```
&emsp;有了这个接口，我们可以做相当强大的事情。我们的WithError变成了这样。

```
func WithError(h HandlerE) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if er, ok := err.(ErrorResponder); ok {
				if er.RespondError(w, r) {
					return
				}
			}

			log.Printf("Something went wrong: %v", err)

			http.Error(w, "Internal server error", 500)
		}
	}
}
```
&emsp;注意到我们的特殊情况是如何消失的吗？它们现在只是ErrorResponder的另一个实现。这就是我们的Not found和Bad请求错误的样子。
```

// BadRequest error responds with bad request status code, and optionally with
// a json body.
type BadRequestError struct {
	err  error
	body interface{}
}

func BadRequest(err error) *BadRequestError {
	return &BadRequestError{err: err}
}

func BadRequestWithBody(body interface{}) *BadRequestError {
	return &BadRequestError{body: body}
}

func (e *BadRequestError) RespondError(w http.ResponseWriter, r *http.Request) bool {
	if e.body == nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusBadRequest)

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(e.body)

		if err != nil {
			log.Printf("Failed to encode a response: %v", err)
		}
	}

	return true
}

func (e *BadRequestError) Error() string {
	return e.err.Error()
}

// Maybe404Error responds with not found status code, if its supplied error
// is sql.ErrNoRows.
type Maybe404Error struct {
	err error
}

func Maybe404(err error) *Maybe404Error {
	return &Maybe404Error{err: err}
}

func (e *Maybe404Error) Error() string {
	return fmt.Sprintf("Maybe404: %v", e.err.Error())
}

func (e *Maybe404Error) Is404() bool {
	return errors.Is(e.err, sql.ErrNoRows)
}

func (e *Maybe404Error) RespondError(w http.ResponseWriter, r *http.Request) bool {
	if !e.Is404() {
		return false
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return true
}
```
&emsp;你可以很容易地编写更多的ErrorResponders来处理权限被拒绝的错误等等。

## Where we’re at
&emsp;通过ErrorResponder和WithError，我们可以将之前的handleThing处理程序缩减为这个。
```
func handleThing(w http.ResponseWriter, r *http.Request) error {
	// Our path is something like /thing/3
	id, err := idFromPath(r.URL.Path)
	if err != nil {
		// Literally bad request. We could use BadRequestWithBody to
		// respond with a fancy information for the client.
		return BadRequest(err)
	}

	thing, err := store.GetThingByID(id)
	if err != nil {
		// Likely a not found issue, but something else might have gone wrong.
		// Maybe404Error handles both cases.
		return Maybe404(err)
	}

	acc := AccountFromRequest(r)
	if acc == nil {
		// No account attached to the request. Client needs to authenticate.
		return AuthenticationRequired()
	}

	has, err := thing.HasPermissionToView(acc)
	if err != nil {
		// Something actually went wrong. Error will be logged and 500 message
		// sent to the client.
		return err
	}

	if !has {
		// Client doesn't have permission to view this resource.
		return PermissionDenied()
	}

	// All good, send data to the client.
	respond(w, r, decodeThing(thing))
}

func main() {
	...
	mux.Handle("/thing/", WithError(handleThing))
	...
}
```
&emsp;这样就好多了! 我将把 auth 和权限检查结合起来，作为一个练习留给读者。另一个练习是在WithError函数中做一个更好的日志记录，而不仅仅是 "Something went wrong: error"。也许可以记录路径和请求者，或者使用跟踪id？

&emsp;有了这些，我们现在可以

- 把错误 "扔 "到某个地方
- 还有人搞清楚了http.Error的调用。
- 非客户端错误被记录

## 文末总结
&emsp;有时候我很惊讶于go的错误类型是如此的简单（但却很强大）。其他时候，我会因为想不出如何使用这种简单性而以头撞墙。在这篇博文中提出的解决方案真的很简单，但我（对我来说）花了很尴尬的时间才 "弄明白"。

&emsp;我也写过不少生锈的文章，我喜欢认为我在生锈中会比在围棋中更快地想出类似的解决方案。但那个生锈的解决方案可能会是次优的，因为会 "夭折"。我写过很多那种 "记录错误，调用http.Error，然后返回 "的错误处理方法，并且忍受着这种重复的痛苦，足以看到所有的各种用例。对于rust，我可能会过早的急于求成，而被糟糕的求成所困。

&emsp;最后，我注意到，经验较少的开发人员不一定有勇气去写一个WithError包装器，在项目的所有代码库中都会用到。他们希望他们使用的工具能提供这样的通用功能，就像django那样。又或者他们 "知道 "go，但不了解go（和它的哲学）？不知道，也许我在这里只是反映了自己的经验。


## 参考
[1].https://medium.com/adapptor/sum-types-in-swift-and-kotlin-7d7286e7f40d
[2].https://en.wikipedia.org/wiki/Tagged_union
[3].https://thingsthatkeepmeupatnight.dev/posts/golang-http-handler-errors/