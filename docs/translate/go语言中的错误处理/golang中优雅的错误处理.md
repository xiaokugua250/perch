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
&emsp;上面代码中出现的部分函数定义需要读者去自行补充，但函数要实现的功能有以下几点。
- 从URL中提取ID参数
- 使用该ID进行数据库查询
- 检查客户端是否有权限查看该响应结果
- 返回结果给客户端  
&emsp;这个功能可能会在其他资源类型上也需要实现，所以很快就会出现重复需求。想象一下，如果对foo、bar、account等资源也做同样的代码实现。而Django中要实现上面的功能，可以选择这样实现代码

```
def handle_thing(request):
    id = id_or_bad_request(request)
    thing = thing_or_404(id)

    account = account_or_forbidden(request)

    if not thing.has_permission(account):
        raise Forbidden()

    return JsonResponse(...)
```
&emsp;以上的python代码利用了id_or_bad_request, thing_or_404和account_or_forbidden函数来响应请求或在出错是直接抛出错误，并且对这些错误的处理由上层逻辑来实现，具体错误处理方法比如可选用正确的状态代码来响应结果的同时记录错误信息。

## 开始简化go语言代码
&emsp;我们先记下上面的python代码，并想一想参考上述的python代码，可以怎样优化我们的go语言代码，来简化go语言中的错误处理逻辑。

- 当代码中出现错误时，我们只想把错误抛出到上层某个地方。这些错误通常是客户端错误，但是不一定。也许上层逻辑能解析和处理这些错误
- 如果错误不是客户端错误，那就需要将错误进行日志记录
- 上层逻辑应该能响应和处理http.Error的调用。

&emsp;Golang的[错误处理](https://blog.golang.org/error-handling-and-go)中阐述了一些关于在http handler中处理错误的常见做法，并且给出了下面的例子。
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
&emsp;上面的代码段通过`http.Error`的函数调用解决了我们刚提出的问题。但有时我们不想把错误的详细信息暴露给客户端，所以我会用一个通用的内部服务器错误信息来代替实际信息。另外，记录服务器内部错误的原因也很重要，否则就无法排查出具体的错误原因。

## 利用WithError函数封装
&emsp;我们希望从实际的`http.Handler`调用中返回错误，需求是在错误出现时能够向客户端优化后的错误结果，并且在返回客户端错误的同时能够记录下错误堆栈信息。代码实现如下。
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
&emsp;嗯，那些 "其他特殊场景 "的需求情况可能会多种多样，而且对于一些处理程序来说可能会变得很特殊。另外，我们还需要写那些is404err和isBadRequest处理程序，以及后续的其他需求。我们通过定义错误接口的形式来进行抽象。
```
type ErrorResponder interface {
    // RespondError writes an error message to w. If it doesn't know what to
    // respond, it returns false.
	RespondError(w http.ResponseWriter, r *http.Request) bool
}
```
&emsp;有了这个接口，我们可以做很多事情，比如，现在我们的WithError变成了这样。

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
&emsp;注意到前一节中的“特殊场景”是如何消失的吗？它们现在只是`ErrorResponder`的另一个接口实现。现在`Not found`和`Bad request`错误形式如下。
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
&emsp;通过参考上面的示例代码，可以很轻松的写出诸如处理权限错误的`ErrorResponders`或其他错误处理函数。

## 目前进度
&emsp;通过`ErrorResponder`接口和`WithError`函数，我们可以将之前的`handleThing`处理程序优化如下 。
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
&emsp;现在的代码看起来就好多了! 我将把认证和权限检查结合起来的需求作为留给读者的练习。另一个练习是在WithError函数中实现更好的日志记录功能，而不仅仅是 "Something went wrong: error" 信息。新的日志记录功能可以尝试去记录请求路径和请求发起人，或者使用跟踪id？
&emsp;有了这些，我们现在可以
- 把错误 返回到上层代码逻辑中
- 上层调用者对http.Error的调用更加清晰
- 记录非客户端错误

## 文末总结
&emsp;有时候我很惊讶于go的错误类型是如此的简单却又如此强大。其他时候，我会因为想不出如何使用这种简单性而郁闷不已。在这篇博文中提出的解决方案真的很简单，但我（对我来说）花了很尴尬的时间才 "弄明白"。  

&emsp;我也写过不少关于`rust`的文章,我有时候会想也许我会在`rust`语言中实现比`go`中更好的错误处理方案。但`rust`中的解决方案可能会是次优的，因为会 "premature"。我写过很多那种 "记录错误（`log error`），调用`http.Error` ，然后返回 (`return`) "的错误处理方法，并且忍受着这种重复的痛苦，足以看到所有的各种用例。对于`rust`，我可能会过早的寻求通用解决方案，但是目前的通用解决方案却很糟糕

&emsp;最后，我注意到，经验较少的开发人员不一定有勇气去写一个可能会在项目的所有代码库中都会用到WithError包装器。他们希望他们使用的工具能提供这样的通用功能，就像django那样。又或者他们 "知道 "go，但不了解go（和它的哲学）？这一切都难以定论，也许我在这里也只是对自己的开发经验做了一个总结反馈而已。


## 参考
[1].https://medium.com/adapptor/sum-types-in-swift-and-kotlin-7d7286e7f40d
[2].https://en.wikipedia.org/wiki/Tagged_union  
[3].https://thingsthatkeepmeupatnight.dev/posts/golang-http-handler-errors/