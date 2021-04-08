  # golang

  ## golang 要点
  #### golang interface
  - > 首先 interface 是一种类型，从它的定义可以看出来用了 type 关键字，更准确的说 interface 是一种具有一组方法的类型，这些方法定义了 interface 的行为。如果一个类型实现了一个 interface 中所有方法，我们说类型实现了该 interface，所以所有类型都实现了 empty interface，因为任何一种类型至少实现了 0 个方法。go 没有显式的关键字用来实现 interface，只需要实现 interface 包含的方法即可。
  - > interface 变量存储的是实现者的值.不难看出 interface 的变量中存储的是实现了 interface 的类型的对象值，这种能力是 duck typing。在使用 interface 时不需要显式在 struct 上声明要实现哪个 interface ，只需要实现对应 interface 中的方法即可，go 会自动进行 interface 的检查，并在运行时执行从其他类型到 interface 的自动转换，即使实现了多个 interface，go 也会在使用对应 interface 时实现自动转换，这就是 interface 的魔力所在。
  - > 一个 interface 被多种类型实现时，有时候我们需要区分 interface 的变量究竟存储哪种类型的值，go 可以使用 comma, ok 的形式做区分 value, ok := em.(T)：em 是 interface 类型的变量，T代表要断言的类型，value 是 interface 变量存储的值，ok 是 bool 类型表示是否为该断言的类型 T。
  - > interface{} 是一个空的 interface 类型，根据前文的定义：一个类型如果实现了一个 interface 的所有方法就说该类型实现了这个 interface，空的 interface 没有方法，所以可以认为所有的类型都实现了 interface{}。如果定义一个函数参数是 interface{} 类型，这个函数应该可以接受任何类型作为它的参数。go 不会对 类型是interface{} 的 slice 进行转换 。为什么 go 不帮我们自动转换，一开始我也很好奇，最后终于在 go 的 wiki 中找到了答案 https://github.com/golang/go/wiki/InterfaceSlice 大意是 interface{} 会占用两个字长的存储空间，一个是自身的 methods 数据，一个是指向其存储值的指针，也就是 interface 变量存储的值，因而 slice []interface{} 其长度是固定的N*2，但是 []T 的长度是N*sizeof(T)，两种 slice 实际存储值的大小是有区别的(文中只介绍两种 slice 的不同，至于为什么不能转换猜测可能是 runtime 转换代价比较大)。
  - >nterface 定义时并没有严格规定实现者的方法 receiver 是个 value receiver 还是 pointer receiver，上面代码中的 S 的 Set receiver 是 pointer，也就是实现 I 的两个方法的 receiver 一个是 value 一个是 pointer，使用 f(s)的形势调用，传递给 f 的是个 s 的一份拷贝，在进行 s 的拷贝到 I 的转换时，s 的拷贝不满足 Set 方法的 receiver 是个 pointer，也就没有实现 I。go 中函数都是按值传递即 passed by value。 如果是按 pointer 调用，go 会自动进行转换，因为有了指针总是能得到指针指向的值是什么，如果是 value 调用，go 将无从得知 value 的原始值是什么，因为 value 是份拷贝。go 会把指针进行隐式转换得到 value，但反过来则不行。对于 receiver 是 value 的 method，任何在 method 内部对 value 做出的改变都不影响调用者看到的 value，这就是按值传递。


参考资料
[1]. https://sanyuesha.com/2017/07/22/how-to-understand-go-interface/