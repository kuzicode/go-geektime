
# Qes

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

# Ask

应该。

由于需要将堆栈信息给到调用者，在dao层warp(只处理一次，且不需要每个地方都打印日志)抛给上层。

上层处理这个error，通过写入日志等等的业务逻辑返回给用户。


example: [main.go](main.go)
