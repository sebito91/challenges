package tree

import (
	"fmt"
	"sort"
)

// Record is the default struct for each entry
type Record struct {
	ID, Parent int
}

// Node is the default struct for parent nodes
type Node struct {
	ID       int
	Children []*Node
}

// Mismatch is a placeholder variable
type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

// declare mappers to be usable across functions
var mappers map[int][]*Node

func checkRecord(record Record) error {
	if record.ID == 0 && record.Parent != 0 {
		return Mismatch{}
	}

	for _, child := range mappers[record.Parent] {
		if record.ID == child.ID {
			fmt.Printf("found dupe: %+v\n%+v\n", record, mappers)
			return Mismatch{}
		}
	}

	return nil
}

// Build stitches together the node tree
func Build(records []Record) (*Node, error) {
	var rootSent bool
	mappers = make(map[int][]*Node)

	if len(records) == 0 {
		return nil, nil
	}

	for idx, record := range records {
		if record.Parent == 0 && record.ID == 0 {
			if rootSent {
				return nil, Mismatch{}
			}

			rootSent = true
			continue
		}

		if err := checkRecord(record); err != nil {
			return nil, err
		}

		mappers[record.Parent] = append(mappers[record.Parent], &Node{ID: record.ID})
		fmt.Printf("%d: %+v\n", idx, record)
		fmt.Printf("DEBUG -- mappers: %+v\n", mappers)
	}

	if !rootSent {
		return nil, Mismatch{}
	}

	// generate the root
	root := &Node{}

	parents := make([]int, 0, len(mappers))
	for parent := range mappers {
		parents = append(parents, parent)
		sort.Slice(mappers[parent], func(i, j int) bool {
			return mappers[parent][i].ID < mappers[parent][j].ID
		})
	}

	sort.Ints(parents)

	for parent := range parents {
		fmt.Printf("handling parent: %d\n", parent)

		if parent == 0 {
			root.Children = append(root.Children, mappers[parent]...)
			continue
		}

		var temproot *Node
		for _, child := range root.Children {
			if child.ID == parent {
				temproot = child
				break
			}
		}

		if temproot != nil {
			temproot.Children = append(temproot.Children, mappers[parent]...)
		}
	}

	return root, nil
}

// Build stitches together the node tree
//func Build(records []Record) (*Node, error) {
//	if len(records) == 0 {
//		return nil, nil
//	}
//	root := &Node{}
//	todo := []*Node{root}
//	n := 1
//	for {
//		if len(todo) == 0 {
//			break
//		}
//
//		newTodo := []*Node(nil)
//
//		for _, c := range todo {
//			for _, r := range records {
//
//				if r.Parent == c.ID {
//					if r.ID < c.ID {
//						return nil, errors.New("a")
//					} else if r.ID == c.ID {
//						if r.ID != 0 {
//							return nil, fmt.Errorf("b")
//						}
//					} else {
//						n++
//						switch len(c.Children) {
//						case 0:
//							nn := &Node{ID: r.ID}
//							c.Children = []*Node{nn}
//							newTodo = append(newTodo, nn)
//						case 1:
//							nn := &Node{ID: r.ID}
//							if c.Children[0].ID < r.ID {
//								c.Children = []*Node{c.Children[0], nn}
//								newTodo = append(newTodo, nn)
//							} else {
//								c.Children = []*Node{nn, c.Children[0]}
//								newTodo = append(newTodo, nn)
//							}
//						default:
//							nn := &Node{ID: r.ID}
//							newTodo = append(newTodo, nn)
//						breakpoint:
//							for range []bool{false} {
//								for i, cc := range c.Children {
//									if cc.ID > r.ID {
//										a := make([]*Node, len(c.Children)+1)
//										copy(a, c.Children[:i])
//										copy(a[i+1:], c.Children[i:])
//										copy(a[i:i+1], []*Node{nn})
//										c.Children = a
//										break breakpoint
//									}
//								}
//								c.Children = append(c.Children, nn)
//							}
//						}
//					}
//				}
//			}
//		}
//		todo = newTodo
//	}
//
//	if n != len(records) {
//		return nil, Mismatch{}
//	}
//
//	if err := chk(root, len(records)); err != nil {
//		return nil, err
//	}
//
//	return root, nil
//}

func chk(n *Node, m int) (err error) {
	if n.ID > m {
		return fmt.Errorf("z")
	} else if n.ID == m {
		return fmt.Errorf("y")
	} else {
		for i := 0; i < len(n.Children); i++ {
			err = chk(n.Children[i], m)
			if err != nil {
				return
			}
		}
		return
	}
}
