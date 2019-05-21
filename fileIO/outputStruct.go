package fileIO

type Output struct {
	lcs string
	longestDis int
}

func (output *Output)SetLcs(lcs string)  {
	output.lcs = lcs
}

func (output *Output)GetLcs() string {
	return output.lcs
}

func (output *Output)SetlongestDis(longestDis int)  {
	output.longestDis = longestDis
}

func (output *Output)GetlongestDis() int {
	return output.longestDis
}
