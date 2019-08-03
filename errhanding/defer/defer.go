package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer(){
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occureed")
	fmt.Println(4)
}

func writeFile(filename string){
	fmt.Printf("flag %d\n",os.O_EXCL|os.O_CREATE)
	file,err := os.OpenFile(filename,os.O_EXCL|os.O_CREATE,0666)
	defer file.Close()
	if err != nil{

		if pathError,ok := err.(*os.PathError); !ok{
			panic(err)
		}else{
			fmt.Println(pathError.Op,
				pathError.Err,
				pathError.Path)
		}

		fmt.Println("error: ",err.Error())
	}

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for i :=5; i < 10; i++{
		fmt.Fprintln(writer,i)
		return
	}
}


func main(){
	writeFile("file")
}
