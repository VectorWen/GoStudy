package main

import (
	"io"
	"log"
	"os"
	"strings"
	"fmt"
)

func main() {
	
	r := strings.NewReader("same io.Reader stream to be read \n")

	_,err := io.Copy(os.Stdout,r)
	//如果成功，err 为nil，而不是io.EOF
	if err != nil{
		log.Fatal(err)
	}

	r = strings.NewReader("same io.Reader stream to be read \n")
	buf := make([]byte,8)
	_,err = io.CopyBuffer(os.Stdout,r,buf)
	if err != nil{
		log.Fatal(err)
	}

	r = strings.NewReader("same io.Reader stream to be read \n")
	_,err = io.CopyBuffer(os.Stdout,r,nil)
	if err != nil{
		log.Fatal(err)
	}

	//Go 建议使用error 来代替抛出错误，但是为什么你又抛出来呢，真是不懂你哦
	defer func () {
		reError :=  recover()
		if reError != nil{
			fmt.Println("出现错误了 -- ",reError)
		}
	}()
	
	r = strings.NewReader("same io.Reader stream to be read \n")
	buf = make([]byte,0)
	_,err = io.CopyBuffer(os.Stdout,r,buf)
	if err != nil{
		log.Fatal(err)
	}


}