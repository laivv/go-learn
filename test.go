package main

import (
	"fmt"
	"strconv"
	"time"
	"reflect"
)


func main(){
	log("hello go")
	doSwap(1,2)
	testFor()
	testChan()
	testVar()
	testIota()
	testIf()
	testSlect()
	testClass()
	testScope()
	testMultipleReturn()
	testArray()
}

func typeOf(arg interface{}) string{
	return reflect.TypeOf(arg).String()
}

func toString(arg interface{}) string{
	ret := ""
	if typeOf(arg) == "int" {
		// 类型断言
		ret = strconv.Itoa(arg.(int))
	}
	if typeOf(arg) == "string" {
		// 类型断言 
		ret = arg.(string)
	}
	return ret
}


func log(args ...interface{}){
	str := ""
	for _,arg := range args{
		str += toString(arg)
	}
	fmt.Println(str)
}


func log1(args ...string){
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
	*a ,*b = *b , *a //交换
}

func doSwap(x int, y int){
	log("交换前",x,y)
	swap(&x,&y)
	log("交换后",x,y)
}

func testFor(){
	for i := 0;i<10;i++{
		log(i)
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
	log("来自goroutine的通道消息",msg)
	//遍历通道的所有消息
	for data := range c {
		log("来自goroutine的通道消息",data)
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
	log(a)
	log(b)
	log(c)
	// len计算字符串常量的长度，不可用于变量
	log("d的长度是：",len(d))
	log(e)

	log(f)
	log(g)
	log(h)
	log(i)
	log(j)
	log(k)
	log(l)
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
	log(a)
	log(b)
	log(c)
	log(d)
	log(e)
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
	//利用 switch进行类型判断
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
	log(len(str))

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

func testSlect(){
	ch1,ch2 := make(chan int,10),make(chan int,10)

	go select_worker(ch1)
	go select_worker(ch2)

	 finish1, finish2 := false ,false

	//利用死循环和select的阻塞特性来接收消息
	for true{
		select {
			case r1, ok1 := (<- ch1): //(<- ch1) 同 <- ch1
				 log("收到ch1通道发来的消息",r1)
				 if(!ok1){
					finish1 = true
					if(finish1 && finish2){
						goto PRINT
					}
				 }
			case r2, ok2 := <- ch2:
				log("收到ch2通道发来的消息",r2)
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


//结构 类相关
type Person struct {
	name string
	age int
}

func (p *Person) getName() string{
	return p.name
}

func (p *Person) getAge() int{
	return p.age
}

func testClass(){
	p := Person{
		name:"lingluo",
		age:18 } // }不能放在下一行，否则报错

	log("p.getName() 的值是:",p.getName())
}


// 作用域 、提升相关
func testScope(){
		//测试块级作用域
		log("测试块级作用域和变量提升")
		var a = 1
		{
			log(a) // 1
			var a = 2
			log(a) // 2
		}
		log(a) // 1
		// go语言具有块级作用域 ，并且不会提升变量
	
		log("测试函数提升")
	
		// (1)
		var myfunc = func (){ 
			log("我是在函数体内定义的函数")
		}
	
		// (2)
		myfunc()
	
		// (1) (2) 不可交换位置，否则报错
		// go语言在函数体内定义函数时，只能用变量赋值的形式，不可采用 func funcName(){}的形式

		//go语言 全局函数、全局变量 有 ‘提升’的效果，可以写在main函数后面
}


// 多个返回值 go语言支持返回多个值
func multipleReturn()(int,string){
	return 1,"abc"
}

func testMultipleReturn(){
		log("go函数支持返回多个值")
		a,b:= multipleReturn()
		log("返回的第一个值是",a)
		log("返回的第二个值是",b)
}



//数组相关
func testArray(){
	//定义数组方式,不指定成员个数
	var array = [...]int{1,2,3}
	//指定成员个数
	array2 := [4]string{"a","b","c","d"}

	array3 := [2]Person{
		Person{"lingluo",18},
		Person{"lingluo2",19} } // 最后一个}不能换行，要报错

	array4 := [...]interface{}{ //利用空接口 interface{} 放入任意类型
		"str",
		123,
		Person{"lingluo",18}}

	// 访问数组
	log(array3[0].name)
	
	// 数组遍历，
	for i ,len := 0,len(array2);i < len; i ++ {
		log(array2[i])
	}
	// 数组遍历， range方式
	for _,item := range array2 {
		log(item)
	}
	_ = array
	_ = array2
	_ = array3
	_ = array4
}