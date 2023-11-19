# go协程概念
- go语言中每个并发的执行单元叫做goroutine
- 当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine，新的goroutine会用go语言来创建
- 在语法上，go语句是一个普通的函数或方法，调用前加上关键字go，go语句会使其语句中的函数在一个新创建的goruntine中运行，而go语句本身会迅速地完成
- 主函数退出或者直接终止程序，会打断所有的goroutine