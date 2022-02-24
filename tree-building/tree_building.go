package tree

import (
	"errors"
	"sort"
)

// Define the Record type
type Record struct {
	ID     int
	Parent int
}

// Define the Node type
type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if records != nil {
		sort.Slice(records, func(i, j int) bool {
			return records[i].ID < records[j].ID
		})
		temp := map[int]*Node{}
		for k, v := range records {
			if v.ID != k+1 {
				return nil, errors.New("non-continuous")
			}
			if v.ID < v.Parent {
				return nil, errors.New("higher id parent of lower id")
			}
			if v.ID == 0 && v.Parent != 0 {
				return nil, errors.New("root node has parent")
			}
			result := &Node{}
			result.ID = v.ID
			temp[v.ID] = result
		}
		if len(records) != len(temp) {
			return nil, errors.New("node have duplicate")
		}
		if _, ok := temp[0]; !ok && len(temp) != 0 {
			return nil, errors.New("no root node")
		}
		for _, v := range records {
			for i := 0; i < len(temp)-1; i++ {
				if v.Parent == temp[i].ID && v.ID != 0 {
					temp[i].Children = append(temp[i].Children, temp[v.ID])
				}
			}
		}
		return temp[0], nil
	}
	return nil, errors.New("input can not be nil")
}
