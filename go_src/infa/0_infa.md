# infa

本目录是 Golang 跨项目的可复用代码库，未来计划单独写成第三方库。

## 快速使用示例

查找 Go 项目根目录。  
`root, err := infa.FindProjectRoot()`



注册退出清理函数。  
`err := infa.RegisterCleanup("grpc client", cleanFn)`

按注册逆序执行退出清理。  
`err := infa.RunCleanup()`

阻塞等待 Ctrl+C 或 SIGTERM。  
`sig := infa.WaitExitSignal()`

使用统一时间格式常量。  
`layout := itime.TimeLayoutDayMillis.String()`

SlicePreview 返回切片长度和样例值。
`SlicePreview`

泛型的三元表达式函数
`IfElse`

创建字符串键控缓存。
`cache := infa.NewStrMapCache[int]()`
