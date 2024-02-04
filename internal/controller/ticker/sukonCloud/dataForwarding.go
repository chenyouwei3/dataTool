package sukonCloud

//func Chan() {
//	fmt.Println("初始化管道成功")
//	global.TestChan = make(chan model.RealtimeData, 1)
//	for {
//		data, ok := <-global.TestChan // 从通道接收数据
//		if !ok {
//			break // 通道已关闭，退出循环
//		}
//		fmt.Println("Received:", data)
//	}
//	conn, err := net.Dial("tcp", "127.0.0.1:8010")
//	if err != nil {
//		fmt.Println("Dial err:", err)
//		return
//	}
//	defer conn.Close()
//	// 主动向服务器发送数据
//	conn.Write([]byte("hello"))
//	// 接收服务器返回数据
//	buf := make([]byte, 4096)
//	n, err := conn.Read(buf)
//	if err != nil {
//		fmt.Println("client read err:", err)
//		return
//	}
//	fmt.Println("client receive：", string(buf[:n]))
//}
