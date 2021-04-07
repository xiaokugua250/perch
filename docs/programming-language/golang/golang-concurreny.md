# golang 并发纪要

1. [errorGroup](https://pkg.go.dev/golang.org/x/sync/errgroup#example-Group-JustErrors) 使用

errorgroup 参考 

使用代码示例
```
func TestWithContext(t *testing.T) {
	errDoom := errors.New("group_test: doomed")

	cases := []struct {
		errs []error
		want error
	}{
		{want: nil},
		{errs: []error{nil}, want: nil},
		{errs: []error{errDoom}, want: errDoom},
		{errs: []error{nil, errDoom}, want: errDoom},
	}

	for _, tc := range cases {
		g, ctx := errgroup.WithContext(context.Background())

		for _, err := range tc.errs {
			err := err
			g.Go(func() error {
				log.Error(err) // 当此时的err = nil 时，g.Go不会将 为nil 的 err 放入g.err中
				return err
			})
		}
		err := g.Wait() // 这里等待所有Go跑完即add==0时，此处的err是g.err的信息。
		log.Error(err)
		log.Error(tc.want)
		if err != tc.want {
			t.Errorf("after %T.Go(func() error { return err }) for err in %v\n"+
				"g.Wait() = %v; want %v",
				g, tc.errs, err, tc.want)
		}

		canceled := false
		select {
		case <-ctx.Done():
		    // 由于上文中内部调用了cancel(),所以会有Done()接受到了消息
		    // returns an error or ctx.Done is closed 
		    // 在当前工作完成或者上下文被取消之后关闭
			canceled = true
		default:
		}
		if !canceled {
			t.Errorf("after %T.Go(func() error { return err }) for err in %v\n"+
				"ctx.Done() was not closed",
				g, tc.errs)
		}
	}
}

```



## 参考
[1].https://pkg.go.dev/golang.org/x/sync/errgroup
[2].https://driverzhang.github.io/post/goroutine%E6%97%A0%E6%B3%95%E6%8A%9B%E9%94%99%E5%B0%B1%E7%94%A8errgroup/
[3].https://ithelp.ithome.com.tw/articles/10249229
[4].https://www.cyningsun.com/subjects/
[5].https://steveazz.xyz/blog/import-context/