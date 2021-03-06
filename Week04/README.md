谢谢老师的关心。家人恢复得挺好。

问个问题：
```
    package main
    import "fmt"
    // 定义一个结构体
    type MyImplement struct{}
    // 实现fmt.Stringer的String方法
    func (m *MyImplement) String() string {
        return "hi"
    }
    // 在函数中返回fmt.Stringer接口
    func GetStringer() fmt.Stringer {
        var s *MyImplement = nil
        if s == nil {
            return nil
        }
        return s
    }
    func main() {
        // 判断返回值是否为nil
        if GetStringer() == nil {
            fmt.Println("GetStringer() == nil")
        } else {
            fmt.Println("GetStringer() != nil")
        }
    }
```

如果把GetStringer中的
  if s == nil {
        return nil
    }
判断拿掉main层得到的结果就不相同了。
这存在隐式转换吗？
我们知道interface有两个部分 type 和 data
在里面判断的时候就可以和nil做比较。在外层却不行。
这是因为go语言是运行时判断的。在里层进行nil判断的时候他还不是 fmt.Stringer接口。还只是那个普通的类型。

透过这个例子就显而易见了。

```
    package main
    import (
        "io"
    )
    // 声明一个设备结构
    type device struct {
    }
    // 实现io.Writer的Write()方法
    func (d *device) Write(p []byte) (n int, err error) {
        return 0, nil
    }
    // 实现io.Closer的Close()方法
    func (d *device) Close() error {
        return nil
    }
    func main() {
        // 声明写入关闭器, 并赋予device的实例
        var wc io.WriteCloser = new(device)
        // 写入数据
        wc.Write(nil)
        // 关闭设备
        wc.Close()
        // 声明写入器, 并赋予device的新实例
        var writeOnly io.Writer = new(device)
        // 写入数据 这个地方就不能使用.Close() 了虽然他还是那个device
        writeOnly.Write(nil)
    }
```


work4作业让我学习到许多东西。
最开始只知道go get 下包。然后放在src目录下。
这次知道用go mod 
然后也用了wire 但是遇到一个小白问题。一直以为是直接go run main.go 但就是会报错误。
提示：.\main.go:17:7: undefined: InitHttpHandler
因此得补一补go包的使用

后来发现原来 go build 不会。他就能自动帮我把当前目录下的.go文件引进来。 

本例中还是比较简单。随着学习的丰富再引入更多的东西吧。
