package tree

import (
	"fmt"
	"testing"
)

func TestBuildBinaryTreeFromArray(t *testing.T) {
	arr := []*int{
		PointInt(1), PointInt(2), PointInt(2), nil, PointInt(4),
	}

	node := BuildBinaryTreeFromArray(arr)

	fmt.Printf("%+v\n", node)
}
