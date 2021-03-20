# 项目内错误处理

### goroutine中panic

goroutine 中的 panic 无法被外围的 recover 捕获。
因此启用goroutine时需做好recover


#### 使用errors.Wrap 处理错误

在多个函数调用时，不推荐通过 fmt.Errorf 处理错误,这样会导致除了错误文本信息之外的堆栈信息丢失。

```go
fmt.Errorf("ScanFileFailed : %v", err) //导致堆栈信息丢失
```

通过 pkg/errors 包中的 errors.Wrap, errors.Unwrap, errors.Cause, 等方法处理错误，


- `errors.Wrap()`, 可以对收到的错误信息，进一步包装，添加额外的描述；但是不会导致的堆栈信息丢失
- `errors.Unwrap()`,剥开包装，拿到带有额外描述的错误描述信息
- `errors.Cause()`, 得到导致错误的描述信息，即 root error

### errors.Wrap()的使用

如果和其他库进行协作，考虑使用 errors.Wrap 或者 errors.Wrapf 保存堆栈信息。同样适用于和标准库协作的时候。 

直接返回错误，而不是每个错误产生的地方到处打日志。

建议在程序的顶部使用 *%+v* 把堆栈详情记录。比如在 web 框架中，会有中间件专门对 panic ， recover 进行处理的。可以在这统一处理。

errors.Wrap() 通常在应用代码中使用，在一些基础库中，直接返回根错误，不做额外的处理，Wrap error 的操作留给业务调用方法处理。

如果在应用代码中，不打算处理错误，就直接用 errors.Wrap 包装错误，直接返回。



**业务代码少用 sentinel error**

sentinel error：包级别 定义的错误类型。使用 sentinel 值并不灵活，因为调用方必须使用 `==` 将结果与预先声明的值进行比较。

- Sentinel errors  不好添加上下文，额外信息时，不够灵活, 
- Sentinel errors 在两个包之间创建了依赖。



