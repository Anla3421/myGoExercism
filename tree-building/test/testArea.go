// func BuildHelper(records []Record, i int) []*Node {
// 	if records != nil {
// 		node := BuildHelper(records, i+1)
// 		// node = append(node...)
// 		for _, v := range records {
// 			if v.ID == 0 && v.ID == v.Parent {
// 				node.ID = v.ID
// 				// node.Children
// 			}
// 		}
// 		return nil
// 	}
// 	return nil
// }

package tree

import "sort"

type Record struct {
	ID     int
	Parent int
}

// type Node struct {
// 	ID       int     `json:"id"`
// 	Children []*Node `json:"Children"`
// }

type Node struct {
	ID       int
	Children []*Node
}

var mainNode *Node

func Build(records []Record) (*Node, error) {
	// 先建立一個上下層關係的map[Parent][]*Node
	// key 為Parent value是slice 存有其所歸屬的node
	relationMap := make(map[int][]*Node)
	if len(records) == 0 {
		return nil, nil
	}
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	// 迴圈
	for _, record := range records {
		// 這邊先找到主要的節點 我先當作Parent 跟Id相同時 代表是主要節點
		if record.Parent == record.ID {
			mainNode = &Node{
				ID:       record.ID,
				Children: []*Node{},
			}
			// 主要節點代表沒有在更上層的node了 所以跳開
			continue
		}

		//	當map中已有Parent 則 直接加到slice
		_, isExist := relationMap[record.Parent]
		if isExist {
			relationMap[record.Parent] = append(relationMap[record.Parent], &Node{
				ID:       record.ID,
				Children: []*Node{},
			})
		} else { // 不存在 則開新的slice
			relationMap[record.Parent] = []*Node{
				{
					ID:       record.ID,
					Children: []*Node{},
				},
			}
		}
	}
	// 從主節點開始
	GetChildren(mainNode, relationMap)
	return mainNode, nil
}

// 跑遞迴 從最上面的node開始往下找
func GetChildren(node *Node, relationMap map[int][]*Node) (nodes []*Node) {
	// 開始尋找下方的node
	// 從主節點 開始找他下面的child
	childrens, isExist := relationMap[node.ID]
	if isExist {
		// 如果有child 則一個一個加入
		for _, eachChildNode := range childrens {
			node.Children = append(node.Children, eachChildNode)
			// fmt.Println("找", node.ID, "的小孩", relationMap[node.ID])

			// call自己 再繼續往下找 直到找不到為止
			GetChildren(eachChildNode, relationMap)
		}
	}
	return
}
