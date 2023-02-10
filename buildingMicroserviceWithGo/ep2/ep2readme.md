go mod init github.com/ocean5tech/repository4go/buildingMicroserviceWithGo/ep2

l *log.Logger 应用于依赖注入，测试等场景


遇到 go.mod中的 go module不生效两个解决方法
   1.  go env -w GO111MODULE=on
       用 go env 确认GO111MODULE状态，没有打开就用上面命令打开
   2.  可能时vs code插件报错 gopls？， 需要在go.mod的当前文件夹打开工程，不能用上级文件夹，因为可能包含多个mod


http.ListenAndServe(":9090", sm) 可以用定制化的server来代替

优雅shutdown


	go func() {      // 用 go 启动 goroute
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)   // make一个channel接受os消息，机制需要再研究一下
	signal.Notify(sigChan, os.Interrupt)  // 有很多种signal,不要持续监听
	signal.Notify(sigChan, os.Kill)
    // 是的，信号量没有捕获到就会阻塞在这里
	sig := <-sigChan // 何时执行有点搞不懂，难道没有signal就block住了，不会进到这条语句？
    
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
