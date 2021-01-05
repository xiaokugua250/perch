  # golang

  ## golang高级
  ### golang内存管理
  - [x] golang 内存模型
  - [x] golang 内存管理
  - happens before
  - [x] golang 内存分配
   - > Go语言内置运行时（就是runtime），抛弃了传统的内存分配方式，改为自主管理。这样可以自主地实现更好的内存使用模式，比如内存池、预分配等等。这样，不会每次内存分配都需要进行系统调用。
Golang运行时的内存分配算法主要源自 Google 为 C 语言开发的[TCMalloc算法](https://wallenwang.com/2018/11/tcmalloc/)，全称Thread-Caching Malloc。核心思想就是把内存分为多级管理，从而降低锁的粒度。它将可用的堆内存采用二级分配的方式进行管理：每个线程都会自行维护一个独立的内存池，进行内存分配时优先从该内存池中分配，当内存池不足时才会向全局内存池申请，以避免不同线程对全局内存池的频繁竞争。[图解golang内存分配]([图解golang内存分配](https://juejin.im/post/5c888a79e51d456ed11955a8))

  - [x] golang gc垃圾回收
  - [x] golang语言设计与实现[golang语言设计与实现](https://draveness.me/golang/)
  - [x] 深入解析go [深入解析go](https://tiancaiamao.gitbooks.io/go-internals/content/zh/) 


参考资料
1. https://www.linuxzen.com/go-memory-allocator-visual-guide.html [golang 内存分配,mcache,mcentral,mspan等]
2. https://golang.org/ref/mem [golang 内存模型 init,goroutine,chan,mux...]
3. https://zhuanlan.zhihu.com/p/93793349 [golang 内存分配]
4. golang https://colobu.com/2019/08/28/go-memory-leak-i-dont-think-so/ [内存泄漏]
5. https://segmentfault.com/a/1190000020338427 [golang 内存分配]
6. http://legendtkl.com/2017/04/28/golang-gc/ [垃圾收集]
7. https://juejin.im/post/5d56b47a5188250541792ede [垃圾回收]
8. https://juejin.im/post/5d78b3276fb9a06b1829e691 [垃圾回收]
9. https://tiancaiamao.gitbooks.io/go-internals/content/zh/ [深入解析golang]
10. https://jin-yang.github.io/post/theme-language-golang.html [golang高级]