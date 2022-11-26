[如何编译调试 Go runtime 源码](https://www.luozhiyun.com/archives/506)
dlv调试器: 
go install github.com/go-delve/delve/cmd/dlv@latest
cd $GOPATH/bin && cp dlv /usr/local/bin
dlv version
查看项目代码行数: cloc src(yum -y install cloc)
设置为schedtrace=X 和 scheddetail=1 会使调度器每 X 毫秒打印详细的多行信息，调度器的描述状态，处理器、线程和goroutines: 
GODEBUG=schedtrace=1000,scheddetail=1000 mygo run helloworld.go
GODEBUG=asyncpreemptoff=1 