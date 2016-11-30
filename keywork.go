package main

import "fmt"
import "time"

const a int = 9

func main() {
	//格式化输出，format
	fmt.Println("Hello Test!!")

	//1. 可以有初始化语句。2. if 后面不需要()
	if i := 9;i < 10{
		fmt.Println("if -- i < 10")
	}else{
		fmt.Println("if -- i >= 10")
	}

	//上面初始化语句的i 到了这里面是未定义的
	//fmt.Println(i)

	//1. switch 后面不需要()。2. 默认case 是不穿透的，使用fallthrough 才穿透
	switch 3{
		case 1:
			fmt.Println("switch -- 1")
		case 2:
			fmt.Println("switch -- 2")
		case 3:
			fallthrough
		case 4:
			fmt.Println("switch -- 3,4")
		default:
			fmt.Println("switch -- default")
	}

	//可以使用string 类型
	switch "a"{
		case "a":
			fmt.Println("switch -- string -- a")
	}

	//1. for 后面不需要()。
	for i:=0;i<10;i++{
		fmt.Println("for -- ",i)

		if i == 4{
			continue
		}

		if i == 6{
			break
		}
	}

	//实现一个接口只需要含有该接口的方法
	var p Person
	testInterface(p)

	//select 用来多通道通信的
	c := make(chan bool)
	select{
	case v := <-c:
		fmt.Println("select -- ",v)
	case <-time.After(1 * time.Second):
		fmt.Println("select -- timeout")
	}

	//defer
	//把参数改为0 就可以测试了
	testDefer(1)

	//go 多线程
	go print099()
	// go print099()
	// go print099()

	//map 的创建
	m := make(map[string]int)
	m["s"] = 3
	fmt.Println("map -- ",m) //map[s:3]
	m["b"] = 8
	fmt.Println("map -- ",m) //map[s:3 b:8]
	delete(m,"b")
	fmt.Println("map -- ",m) //map[s:3]

	if 4 == 9{
		fmt.Println("9 == 4")
	}else{
		goto L
	}

	fmt.Println("goto -- no goto")
	L:
	fmt.Println("goto -- if no print no goto is goto")

	//const
	fmt.Println("const -- ",a)

	var xs = []int{1,2,3,4,5,6,7,8,9,10}
	var sum = 0;

	//遍历
	for v := range xs{
		sum += v
	}
	fmt.Println("range -- ",sum)

	//pause main thread 3 second
	select{
	case <-time.After(3 * time.Second):
		fmt.Println("Test Success!!")
	}
}

// interface test need
func testInterface(eat Eater) {
	eat.Eat()
	
}

// 定义一个接口
type Eater interface{
	Eat() string
}

type Person struct{
}

//实现接口只要有这个方法就够了
func (p Person) Eat() string {
	fmt.Println("interface -- Person eat")
	return "ok"
}

//defer
func testDefer(a int) {
	defer func () {
		fmt.Println("defer -- run defer func")
		if err:=recover(); err!= nil{
			fmt.Println(err)
		}
	}()
	v := 1/a
	fmt.Println("defer -- ",v)
}


// print 0 - 99
func print099() {
	for i := 0; i < 99; i++ {
		fmt.Println("go -- ",i)
	}
}














