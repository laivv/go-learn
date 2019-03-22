package log

import (
	"strconv"
)



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


func Log(args ...interface{}){
	str := ""
	for _,arg := range args{
		str += toString(arg)
	}
	fmt.Println(str)
}
