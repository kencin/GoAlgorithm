package fileIO

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func WriteOut(output *Output){
	outputFile , err :=os.OpenFile("datas/output.txt",os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil{
		fmt.Println("写入LCS结果出现错误！")
		fmt.Println(err)
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	_, errs := outputWriter.WriteString("最长子序列长度为：" + strconv.Itoa(output.GetlongestDis()) + "\r\n")
	if errs != nil{
		fmt.Println("写入LCS结果出现错误！")
		return
	}
	_, errs2 := outputWriter.WriteString("最长子序列为：" + output.GetLcs() + "\n")
	if errs2 != nil{
		fmt.Println("写入LCS结果出现错误！")
		return
	}
	outputWriter.Flush()
	fmt.Println("写入LCS结果成功")
}
