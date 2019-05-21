// 随机产生20以上的字符，放入输入文件input.txt，如：X={A,B,C,B,D,A,B}和Y={B,D,C,A,B,A}

package fileIO

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	s1 := rand.NewSource(time.Now().UnixNano()) //用当前时间创建一个随机数种子
	r1 := rand.New(s1)
	for i := range b {
		b[i] = letters[r1.Intn(len(letters))]
	}
	return string(b)
}

func MakeFile(){
	fmt.Println("正在产生随机字符串...")
	outputFile , err :=os.OpenFile("datas/input.txt",os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil{
		fmt.Println("产生随机字符串出现错误！")
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	for i:=0; i<2; i++ {
		s1 := rand.NewSource(time.Now().UnixNano()) //用当前时间创建一个随机数种子
		r1 := rand.New(s1)
		_, errs := outputWriter.WriteString(randSeq(r1.Intn(80)+20) + "\n")
		if errs != nil{
			fmt.Println("产生随机字符串出现错误！")
			return
		}
		time.Sleep(time.Millisecond*50)
	}
	outputWriter.Flush()
	fmt.Println("随机字符串生成完成")
}

