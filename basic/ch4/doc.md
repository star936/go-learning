### internal机制
> 这是包用来封装Go程序最重要的机制

以这种方式定义的标识符可以被一个小的可信任的包集合访问，但不是所有的包都可以访问.
`go build`工具会特殊对待导入路径中包含路径片段`internal`的情况，这些包叫做`内部包`.改`内部包`只能被以`internal`目录的父目录为根目录树下的其它目录导入。例如:
* net/http
* net/http/internal/chunked
* net/http/httputil
* net/url
  
其中`chunked`包可以被`net/http`、`net/http/httputil`包导入，而不能被`net/url`导入.
