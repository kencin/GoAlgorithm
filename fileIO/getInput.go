// 从文件中得到input字符串

package fileIO

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func GetInput() []string{
	inputFile, err := os.Open("datas/input.txt")
	if err!=nil{
		fmt.Println("打开input.txt出错！")
		return nil
	}
	defer inputFile.Close()

	var stringArr []string
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, _, readerErr := inputReader.ReadLine()
		if readerErr == io.EOF{
			break
		}
		stringArr = append(stringArr, string(inputString))
	}
	//for _, line := range stringArr {
	//	fmt.Println(line)
	//}
	return stringArr
}
