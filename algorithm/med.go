// 最小编辑距离
package algorithm

import "fmt"

//func main(){
//	med("ACD","ABCD")
//}

func Med(a,b string){
	var dp[][] int
	dp = make([][]int, len(a)+1)
	for i:=0;i<len(a)+1;i++{
		dp[i] = make([]int, len(b)+1)
	}
	stringA := []rune(a)
	stringB := []rune(b)

	dp[0][0] = 0

	for i:=1;i<len(a)+1;i++{
		dp[i][0] = i
	}

	for i:=1;i<len(b)+1;i++{
		dp[0][i] = i
	}

	for i:=1;i<len(a)+1;i++{
		for j:=1;j<len(b)+1;j++{
			if stringA[i-1] == stringB[j-1]{
				dp[i][j] = minInThree(dp[i][j-1]+1, dp[i-1][j]+1, dp[i-1][j-1])
			} else {
				dp[i][j] = minInThree(dp[i][j-1]+1, dp[i-1][j]+1, dp[i-1][j-1]+1)
			}
		}
	}
	// 打印dp矩阵
	//for tmp:=0; tmp<len(a)+1; tmp++{
	//	for tmp2:=0; tmp2<len(b)+1; tmp2++{
	//		fmt.Print(dp[tmp][tmp2])
	//		fmt.Print(" ")
	//	}
	//	fmt.Println()
	//}

	fmt.Print("最小编辑距离为：")
	fmt.Println(dp[len(a)][len(b)])
}

func minInThree(a,b,c int)int{
	if a<b{
		if a<c{
			return a
		} else {
			return c
		}
	} else {
		if b<c{
			return b
		} else {
			return c
		}
	}
}