

# 1. Golang Error 概念

## 1.1 Go error 本质

```golang
// 本质上是一个普通的接口
type error interface{
    Error() string
}
```

- 一般用error.New()来返回error对象
`func New(text string) error`

- 基础库中大量自定义的error：https://golang.org/src/bufio/bufio.go

- error.New()返回内部`errorString`对象的指针：https://golang.org/src/errors/errors.go

## 1.2 Error vs Exception

- c: 单返回值
- c++: 引入了`exception`，但是被调用方抛异常情况难以分析
- java: 引入`check excetion`, 方法所有者必须申明，调用者必须处理；异常变得司空见惯，由函数调用者来区分。
- go 不引入excetion，支持多参数返回，很容易在函数签名上实现error interface,由调用者来判定。

```golang
func handle()(int, error){
    return 1, nil
}

func main(){
    i, err := handle()
    if err != nil{
        return
    }
    fmt.Println(i)
}
```

go error 设计理念：先判断error(==nil)，在决定value能不能用

example
```golang
import(
    "errors"
    "fmt"
)

func Positive(n int)(bool, error){
    if n == 0{
        return false, errors.New("Undefine!!!!")
    }
    return n > -1, nil
}

func Check(n int){
    pos, err := Positive(n)m
    if err != nil{
        fmt.Println(n, err)
        return
    }
    if pos{
        fmt.Println(n, "is positive")
    }else{
        fme.Println(n, "is negative")
    }
}
```

## 1.3 panic

go panic == fatal err(挂了)，真正不可恢复的程序错误才会用panic，如索引越界、不可恢复的环境问题、堆栈溢出等等

- panic: 立刻停止执行函数的其他代码
- recover: 终止panic造成的程序崩溃，只能在defer中发挥作用，其他作用域中没用

设计理念：
- easy
- plan to failure, not success
- 无隐藏的控制流
- 交由操作者来控制error
- Error are Value in golang, not excetion

```golang
// case: panic 程序终止退出(野生gorutine)
func main(){
    fmt.Println("go!")
    go func()  {
        fmt.Println("gg??")
        panic("gg seubnida")
    }

    time.Sleep(5 * time.Second)
}


// good case: panic recover 保护程序不会退出，只打印panic信息
func main(){
    fmt.Println("go!")
    Go(func()  {
        fmt.Println("gg??")
        panic("gg seubnida")
    })

    time.Sleep(5 * time.Second)
}

func Go(x func()){
    go func(){
        defer func(){
            if err := recover(); err != nil{
                fmt.Println(err)
            }
        }()
        x()
    }()
}
```


结论：**Error让开发者知道什么时候出了mm问题，真正的异常情况留给了Panic**


# 2 Error Type

## 2.1 Sentinel Error - 预定义特定错误

```golang
var ggError = errors.New("gg seubnida!!")
```
使用特定的值(包级别的)来处理不可能进行下一步工作的错误。
特点表现为：
- if err == ErrSomething{...}， seems like io.EOF, more underlting syscall.ENOENT
- not require error.Error output
缺点也很明显：
- Sentinel errors 成为 API 公共部分。
- Sentinel errors 在两个包之间创建了依赖。
结论：尽可能避免 sentinel errors。


## 2.2 Error Type - 错误类型

实现了 error 接口的自定义类型。

```golang
// Define
type MyError struct{
    Msg string
    File string
    Line int
}

func (e *MyError) Error() string{
    return fmt.Sprintf("%s:%d:%s", e.File, e.Line, e.Msg)
}

func test()Error{
    return &MyError{"Something happened", "server.go", 42}
}

// Transfer
func main(){
    err := test()
    switch err := err.(type) {
    case nil:
        // success, nothing to do
    case *MyError:
        fmt.Println("error occurred on line:", err.Line)
    default:
        // unknow error
    }
}
```

特点表现为：
- 一个指自定义 MyError{} 是一个 type，调用者可以用断言来转换这个类型获取更多的上下文信息。
- 能包装底层错误提供更多上下文
缺点为：
调用者要使用类型断言和类型 switch，就要让自定义的 error 变为 public。这种模型会导致和调用者产生强耦合，从而导致 API 变得脆弱。
结论：也尽可能避免使用 Error Type， 至少说避免将它们作为公共 API 的一部分

## 2.3 Opaque Errors - 不透明错误处理

`不透明错误处理`， 只需返回错误而不假设其内容。

```golang
import "github.com/quux/bar"

func fn() error{
    x, err := bar.Foo()
    if err != nil{
        return err
    }
    // use x
}
```

特点表现为：
- 作为调用者只需要知道它是否起作用了，看不到错误的内部。
- Assert errors for behaviour, not type

结论：最灵活的错误处理策略，要求代码和调用者之间的耦合最少。



# 3. Handing Error

# 3.1 indented flow is for errors

处理原则：先处理error，再进行下一步

```Golang
f, err := os.Open(path)
if err != nil{
    // handle error
}
// do stuff
```

# 3.2 Eliminate error handling by eliminating errors

处理原则：不破坏程序完整性和可读性情况下，尽可能减少对error的处理

```Golang
// bad case
func AuthRequest() error{
    if err:= auth();err!=nil{
        return err
    }
    return nil
}

// nice case
func AuthRequest(r, *Request) error{
    return auth(r.User)
}
```

# 3.3 Wrap errors

处理原则：
1. `error` 只应该被处理一次，输出日志也算处理
2. 每个 `error` 只应该被wrap一次，就是在error第一次产生的地方
3. `error` 的wrap也只应该发生在业务代码中，不应该在出现在基础库中
