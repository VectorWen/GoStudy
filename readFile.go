package main

import "fmt"
import "os"
import "bufio"
import "io"

func main() {
	file,err:=os.Open("/Users/vector.huang/Study/Go/src/github.com/vector/GoStudy/file.go")
	if err != nil{
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for{
		str,err := reader.ReadString(20)
		fmt.Println(str)
		if(err != nil){
			if err == io.EOF{
				fmt.Println("读取完毕！！")
			}else{
				fmt.Println(err)
			}
			break
		}
	}
}
