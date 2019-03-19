package main

import (
	"fmt"
	"strconv"
	"time"
)


func main(){
	log("hello go")
	doSwap(1,2)
	forTest()
	testChan()
	testVar()
	testIota()
	testIf()
	selectTest()
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
	fmt.Println(a)
}


//指针相关
func swap(a *int, b *int){
	var temp int = *a
	*a = *b
	*b = temp
}

func doSwap(x int, y int){
	log2("交换前",x,y)
	swap(&x,&y)
	log2("交换后",x,y)
}

func forTest(){
	for i := 0;i<10;i++{
		log3(i)
	}
}

// channel相关

func worker(c chan int){
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
	go worker(c)
	//从通道接收一条消息
	msg := <- c 
	log2("来自goroutine的通道消息",msg)
	//遍历通道的所有消息
	for data := range c {
		log2("来自goroutine的通道消息",data)
	}
	log("testChan执行完毕")
}



//定义变量相关
func testVar(){
	//定义一个常量
	const a = 10
	var (
		b int = 123
		c = 2
	)
	const (
		d = "567"
		e = 6
	)
	var f int = 1
	var g = 2
	var h,i,j int= 10,11,12
	k,l,m := 13,14,"15"
	log3(a)
	log3(b)
	log3(c)
	// len计算字符串常量的长度，不可用于变量
	log2("d的长度是：",len(d))
	log3(e)

	log3(f)
	log3(g)
	log3(h)
	log3(i)
	log3(j)
	log3(k)
	log3(l)
	log(m)
}

//iota相关
func testIota(){
	const (
		a = iota //iota = 0
		b //1 iota + 1
		c = "str" //iota + 1
		d //str  iota + 1
		e = iota //4
	)
	log3(a)
	log3(b)
	log(c)
	log(d)
	log3(e)
}

//条件相关
func testIf(){
	if(true){
		log("if（ture）就打印这行")
	}

	if(false){

	}else if(true){
		log("else if（ture）就打印这行")
	}

	switch "abc"{ //或switch ("abc")
		case "abc":
			log("case \"abc\"就打印这行")
			//自带break效果
		case "def":
			log("case\"def\"就打印这行")
		default:
			log("default就打印这行")
	}

	flag := "abc"

	switch{
		case flag == "abc":
			log("case flag == \"abc\"就打印这行")
		default:
			log("没有匹配时的默认项")
	}

	var s interface{};
	switch s.(type){
		case int :
			log("s是int类型就打印这行")
		case string:
			log("s是string类型就打印这行")
		default:
			log("s未知类型")
	}
	var str string = "123"
	//获取字符串长度
	log3(len(str))

	var m uint = 1
	switch m{
		case 2:
			log("这条不会打印")
		case 1:
			log("这条一定会打印，在下一行用fallthrough关键字，下一个case无论是否成立都会执行")
			fallthrough
		case 3:
			log("这条case不匹配也会打印，因为上一个case用了fallthrough关键字")
		default:
			log("这条不会打印")
	}

}


//select相关

func select_worker(c chan int){
	for i := 0; i < 10 ; i++{
		c <- i
		//需要导入time包
		time.Sleep(time.Duration(1) * time.Second)
	}
	close(c) //关闭通道
}

func selectTest(){
	ch1,ch2 := make(chan int,10),make(chan int,10)

	go select_worker(ch1)
	go select_worker(ch2)

	 finish1, finish2 := false ,false

	//利用死循环和select的阻塞特性来接收消息
	for (true){
		select {
			case r1, ok1 := (<- ch1): //(<- ch1) 同 <- ch1
				 log2("收到ch1通道发来的消息",r1)
				 if(!ok1){
					finish1 = true
					if(finish1 && finish2){
						goto PRINT
					}
				 }
			case r2, ok2 := <- ch2:
				log2("收到ch2通道发来的消息",r2)
				if(!ok2){
					finish2 = true
					if(finish1 && finish2){
						goto PRINT
					}
				 }
			// default: //如果开启default则select语句不阻塞，否则阻塞,不开启default语句则不要写死循环这种危险操作
			// 	log("暂无消息")
		}
	}

	PRINT:log("消息接收完毕！")


}