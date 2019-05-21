// 最长公共子序列
package algorithm

import (
	"fmt"
	"strings"
	"../fileIO"
)

//func main() {
//	//a := "1A2C3D4B56"
//	//b := "B1D23CA45B6A"
//	//findCommon(a, b)
//	//med(a,b)
//	//time.Sleep(5 * time.Second)
//	fileIO.GetInput()
//}

func FindCommon(a, b string){
	var dp [][] int
	output := new(fileIO.Output)
	dp = make([][]int, len(a))
	for i:=0;i<len(a);i++{
		dp[i] = make([]int, len(b))
	}
	stringA := []rune(a)
	stringB := []rune(b)

	if stringB[0] == stringA[0]{
		dp[0][0] = 1
	} else {
		dp[0][0] = 0
	}

	for tmp:=1;tmp<len(a);tmp++{
		if stringB[0] == stringA[tmp]{
			dp[tmp][0] = max(1, dp[tmp-1][0])
		} else {
			dp[tmp][0] = max(0, dp[tmp-1][0])
		}
	}

	for tmp:=1;tmp<len(b);tmp++{
		if stringB[tmp] == stringA[0]{
			dp[0][tmp] = max(1, dp[0][tmp-1])
		} else {
			dp[0][tmp] = max(0, dp[0][tmp-1])
		}
	}

	for tmp:=1; tmp<len(a); tmp++{
		for tmp2:=1; tmp2<len(b); tmp2++{
			dp[tmp][tmp2] = max(dp[tmp-1][tmp2],dp[tmp][tmp2-1])
			if stringA[tmp] == stringB[tmp2]{
				dp[tmp][tmp2] = max(dp[tmp][tmp2],dp[tmp-1][tmp2-1]+ 1)
			}
		}
	}
	num := dp[len(a)-1][len(b)-1]


	// 打印dp矩阵
	//for tmp:=0; tmp<len(a); tmp++{
	//	for tmp2:=0; tmp2<len(b); tmp2++{
	//		fmt.Print(dp[tmp][tmp2])
	//		fmt.Print(" ")
	//	}
	//	fmt.Println()
	//}
	fmt.Print("最长子序列长度为：")
	fmt.Println(num)
	output.SetlongestDis(num)
	endCol := len(a) -1
	endRow := len(b) -1
	var sb strings.Builder
	for num > 0{
		if endRow>0 && dp[endCol][endRow] == dp[endCol][endRow-1]{
			endRow --
		} else if endCol>0 && dp[endCol][endRow] == dp[endCol-1][endRow]{
			endCol --
		} else {
			sb.WriteRune(stringA[endCol])
			endCol--
			endRow--
			num--
		}
	}

	fmt.Print("最长子序列为：")
	fmt.Println(reverse(sb.String()))
	output.SetLcs(reverse(sb.String()))
	fileIO.WriteOut(output)
}

func max(a,b int) int{
	if a>b {
		return a
	} else {
		return b
	}
}

func reverse(a string) string{
	runes := []rune(a)
	for from, to := 0, len(runes)-1; from < to; from, to = from + 1, to - 1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}