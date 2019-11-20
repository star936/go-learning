### internal机制
> 这是包用来封装Go程序最重要的机制

以这种方式定义的标识符可以被一个小的可信任的包集合访问，但不是所有的包都可以访问.
`go build`工具会特殊对待导入路径中包含路径片段`internal`的情况，这些包叫做`内部包`.改`内部包`只能被以`internal`目录的父目录为根目录树下的其它目录导入。例如:
* net/http
* net/http/internal/chunked
* net/http/httputil
* net/url
  
其中`chunked`包可以被`net/http`、`net/http/httputil`包导入，而不能被`net/url`导入.


### defer

#### 触发时机

1. 包裹着defer语句的函数返回时
2. 包裹着defer语句的函数执行到最后时
3. 当前goroutine发生Panic时

#### return，defer，返回值的执行顺序

1. 先给返回值赋值
2. 执行defer语句
3. 包裹函数return返回

**例子分析:**

```go
func f() int { //匿名返回值
	var r int = 6
	defer func() {
		fmt.Printf("r=%d\n", r)
		r *= 7
		fmt.Printf("r=%d\n", r)
	}()
	return r
}
```

**输出结果: 6**

**分析: `f函数`是匿名返回值，匿名返回值是在return执行时被声明；因此defer声明时，还不能访问匿名返回值，所以defer的修改不会影响到返回值.**