package main

import "fmt"
import "strconv"

func main(){
	log("hello go")
	doChange(1,2)
	forTest()
	testChan()
}

func log(args ...string){
	str := ""
	for _, arg := range args{
		str += arg
	}
	fmt.Println(str)
}

func log2(mark string, others ...int){
	str := mark
	for _,other := range others{
		//int转string
		str += strconv.Itoa(other)
	}
	fmt.Println(str)
}

func log3(a int){
	fmt.Println(a);
}


//指针
func change(a *int, b *int){
	var temp int = *a
	*a = *b
	*b = temp
}

func doChange(x int, y int){
	log2("交换前",x,y)
	change(&x,&y)
	log2("交换后",x,y)
}

func forTest(){
	for i:=0;i<10;i++{
		log3(i)
	}
}


func channel(c chan int){
	var i int = 0;
	for ; i< 10 ; i++{
		c <- i
	}
	close(c) //关闭通道
}

func testChan(){
	// 创建一个通道
	var c chan int = make(chan int,100) //第二个参数为缓冲区大小，带缓冲的通道,发送到缓冲区，不用等待接收方，除非缓存满了
	//c := make(chan int) // 不带缓冲的通道,发送方会阻塞，直到接收方收到
	//开启协程
	go channel(c)
	//从通道接收一条消息
	msg := <- c 
	log2("来自goroutine的通道消息",msg)
	//遍历通道的所有消息
	for data := range c {
		log2("来自goroutine的通道消息",data)
	}
	log("testChan执行完毕")
}
