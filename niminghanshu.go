package main

import (
	"fmt"
)
func add(a,b int) int   {   //函数也可以当做一个变量  这个函数为int类型
	return b+a
}
func testfunc1()  {
	f1 :=add    //f1用来接收add这个函数   传值十时只需要用f1（int，int）
	fmt.Printf("%T\n",f1)
	sum:=f1(2,5)
	fmt.Printf("%d",sum)
}
func testfunc2(){
	f1:=func (a,b int )int{    //这个函数不同于上面这个，上面是有add这个函数的，
								// 而func2函数没有实际的函数，是匿名的韩旭尧自己构建函数体，这就是匿名函数  感觉和func1没啥区别啊
		return a+b
	}
	fmt.Printf("%T\n",f1)
	sum:=f1(2,5)
	fmt.Printf("%d",sum)
}
func testfunc3()  {
	var i int=0
	defer fmt.Printf("%d" ,i)   //defer函数最后才打印为0 
	i=100
	fmt.Printf("%d\n" ,i)   //最后的结果是  100\n0
	return
}
func testfunc4()  {
	var i int=0
	defer func(){
		fmt.Printf("defer i=%d\n",i)  //这个是defer匿名函数的用法
	}()  //这个括号还没搞懂  只是知道没有接受参数而已
	i=100
	fmt.Printf("i=%d\n" ,i)    //最后的结果为i=100\n  defer i=100
	return
}
func sub(a,b int )int   {
	return a-b
}
func calc(a,b int,op func(int,int)int )int   {  //op func(int,int)int==int op(int,int)   int函数类型的op
												//func是个匿名函数
	return op(a,b)
}
func testfunc5()  {
	sum:=calc(100,300,add)   //op是个函数类型的变量所以传值是add  不是add（）  这个要理解
	sub:=calc(100,300,sub)
	fmt.Printf("sum=%d sub=%d",sum,sub)
}
func main()  {
	//testfunc1()
	//testfunc2()
	//testfunc3()
	//testfunc4()
	testfunc5()
}


