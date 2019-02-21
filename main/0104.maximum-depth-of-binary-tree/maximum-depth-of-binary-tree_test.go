package _104_maximum_depth_of_binary_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.intra.xiaojukeji.com/chishui/awesomeProject/kit"
	"testing"
)

func Test_Problem0104(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		pre, in []int
		ans     int
	}{

		{
			[]int{},
			[]int{},
			0,
		},

		{
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
			3,
		},

		{
			[]int{3, 9, 20, 15, 10, 7},
			[]int{9, 3, 10, 15, 20, 7},
			4,
		},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		fmt.Printf("~~%v~~\n", levelOrder1(kit.PreIn2Tree(tc.pre, tc.in)))
		ast.Equal(tc.ans, maxDepth(kit.PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)

		fmt.Printf("~~%v~~\n", levelOrder1(kit.PreIn2Tree(tc.pre, tc.in)))
	}
}
