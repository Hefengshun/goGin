package main

var str string = "123"

//相当于 var str string = "123"
//var data *string = &str
//*data = 20
//func name(data *string) { //按引用传递
//	*data += "12"
//}
//
//func main() {
//	name(&str)  想改变结构体调用的时候才穿指针
//	println(str)
//}

func name(data string) { //按值传递
	data += "12"
	println(&data, "&data 地址已经变化了", data, "data world")
}

// func main() {
// 	name(str)
// 	println(str, "hello world")
// }

//总结
//按值传递：传递的是变量的副本，函数内部的修改不会影响原始变量。
//按引用传递：传递的是变量的地址（指针），函数内部的修改会直接影响原始变量。
//在 Go 语言中，默认情况下是按值传递的。如果你需要在函数内部修改传递的变量并在函数外部看到这些修改，就需要使用指针（传递变量的地址）。
