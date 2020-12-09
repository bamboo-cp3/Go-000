### Golang ERROR

`error`是一个预定义标识符，它代表了一个Go语言内建的接口类型。这个接口类型的声明如下：

~~~go
type error interface {
    Error() string
}
~~~

Go语言的标准库代码包`errors`为我们提供了一个用于创建`error`类型值的函数`New`。该方法的声明如下：

~~~go
func New(text string) error {
    return &errorString{text}
}
~~~

Go panic 意味着 fatal error（就是挂了）。不能假设调用者来解决panic，意味着代码不能继续运行了

### ERRORS VS EXCEPTION
Go 的处理异常不是引入Exception，支持多参数返回，在函数签名中带上实现error 接口的对象，交由调用者处理
* exception 存在的问题：
~~~

~~~

### Sentinel Error

预定义的特定错误，叫sentinel error

1. sentinel 值是最不灵活的错误处理策略，因为调用方必须使用 == 讲结果与预先声明的值进行比较
2. sentinel errors在两个包之间创建了依赖
3. 尽可能避免使用sentinel errors

### Error types

Error type 是实现了 error 接口的自定义类型

1. error type 可以捕获出错的更多上下文但是会导致和调用者产生强耦合

避免使用错误类型



### Opaque errors

这是最灵活的错误处理策略，因为它要求代码和调用者之间的耦合最少，作为调用者，关于操作的结果，只需要知道它起作用了，或者没有起作用(成功还是失败)

1. 断言错误实现了特定的行为，而不是断言错误是特定的类型或值（Assert errors for behaviour, not type）



### Error Handling

1. 无错误的正常流程代码，将成为一条直线，而不是缩进的代码
2. 避免多次处理错误
3. 错误要被日志记录
4. 应用程序处理错误，保证100%完整性
5. 记录日志代表错误已经被处理，之后不再报告当前错误，程序降级



### Wrap errors

you should only handle errors once. Handling an error means inspecting the error value, and making a single decision.