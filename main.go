package main

import (
	"./algorithm"
	"./fileIO"
	"fmt"
)

func main() {
	//a := "1A2C3D4B56"
	//b := "B1D23CA45B6A"
	//findCommon(a, b)
	//med(a,b)
	//time.Sleep(5 * time.Second)
	fileIO.MakeFile()
	stringArr := fileIO.GetInput()
	fmt.Print("字符串a：")
	fmt.Println(stringArr[0])
	fmt.Print("字符串b：")
	fmt.Println(stringArr[1])
	algorithm.FindCommon(stringArr[0],stringArr[1])
	algorithm.Med(stringArr[0],stringArr[1])
}
