package _4_Query_in_Segment_Tree

import (
	"fmt"
	"testing"
)

func TestSegmentTree(t *testing.T) {
	nums := []interface{}{-2, 0, 3, -5, 2, -1}
	segTree := NewSegmentTree(nums, func(a, b interface{}) interface{} {
		return a.(int) + b.(int)
	})
	fmt.Println(segTree)

	fmt.Println(segTree.Query(0, 2)) // 1
	fmt.Println(segTree.Query(2, 5)) // -1
	fmt.Println(segTree.Query(0, 5)) // -3
}

/**
[-3, 1, -4, -2, 3, -3, -1, -2, 0, nil, nil, -5, 2, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil]
-3是所有节点之和
1=-2+3
-4=-5+2-1

-2+0+3=1
*/
