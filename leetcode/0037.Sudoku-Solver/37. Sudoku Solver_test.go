package leetcode

import (
	"fmt"
	"testing"
)

type question37 struct {
	para37
	ans37
}

// para 是参数
// one 代表第一个参数
type para37 struct {
	s [][]byte
}

// ans 是答案
// one 代表第一个答案
type ans37 struct {
	s [][]byte
}

func Test_Problem37(t *testing.T) {

	qs := []question37{

		question37{
			para37{[][]byte{
				[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}},
			ans37{[][]byte{
				[]byte{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				[]byte{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				[]byte{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				[]byte{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				[]byte{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				[]byte{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				[]byte{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				[]byte{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				[]byte{'3', '4', '5', '2', '8', '6', '1', '7', '9'}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 37------------------------\n")

	for _, q := range qs {
		_, p := q.ans37, q.para37
		fmt.Printf("【input】:%v \n\n", p)
		solveSudoku(p.s)
		fmt.Printf("【output】:%v \n\n", p)
	}
	fmt.Printf("\n\n\n")
}