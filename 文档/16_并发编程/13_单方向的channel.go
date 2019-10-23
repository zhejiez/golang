package main

func main() {
	ch := make(chan int)

	//双向channel能隐式转换为单向channel
	var writeCh chan<- int = ch //只能写,不能读
	var readCh <-chan int = ch  //只能读,不能写
	writeCh <- 666
	//<-writeCh  //err, invalid operation: <-writeCh (receive from send-only type chan<- int

	<-readCh
	//readCh <- 666 //err, invalid operation: readCh <- 666 (send to receive-only type <-chan

	//单向无法转换为双向
	var ch2 chan int = writeCh //err, cannot use writeCh (type chan<- int) as type chan int in assignment
}
