package tree

import (
	"fmt"
)

// Record is the specific instance of a file to be added into the document tree
type Record struct {
	ID, Parent int
}

// Node is the primary component of our document tree
type Node struct {
	ID       int
	Children []*Node
}

// Build is the primary function to create the document tree from the provided set of records.
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	nodes := make(map[int]Record)
	parents := make(map[int][]Record)

	for id, record := range records {
		fmt.Printf("%d record: %+v\n", id, record)

		// parent must always be lower (or equal to) the ID
		if record.Parent > record.ID {
			return nil, fmt.Errorf("parent higher than id: %+v", record)
		}

		if record.Parent == record.ID && record.ID != 0 {
			return nil, fmt.Errorf("only root can have the same parent and ID: %+v", record)
		}

		if _, ok := nodes[record.ID]; ok {
			return nil, fmt.Errorf("duplicate node found: %+v", record)
		}

		nodes[record.ID] = record
		parents[record.Parent] = append(parents[record.Parent], record)
	}
	//	{
	//		name: "binary tree",
	//		input: []Record{
	//			{ID: 5, Parent: 1},
	//			{ID: 3, Parent: 2},
	//			{ID: 2, Parent: 0},
	//			{ID: 4, Parent: 1},
	//			{ID: 1, Parent: 0},
	//			{ID: 0},
	//			{ID: 6, Parent: 2},
	//		},
	//		expected: &Node{
	//			ID: 0,
	//			Children: []*Node{
	//				{
	//					ID: 1,
	//					Children: []*Node{
	//						{ID: 4},
	//						{ID: 5},
	//					},
	//				},
	//				{
	//					ID: 2,
	//					Children: []*Node{
	//						{ID: 3},
	//						{ID: 6},
	//					},
	//				},
	//			},
	//		},
	//	},

	root := &Node{ID: 0}
	//	out := make([][]*Node, len(parents))
	//	for i := len(parents) - 1; i >= 0; i-- {
	//		if _, ok := parents[i]; !ok {
	//			return nil, fmt.Errorf("parents out of order")
	//		}
	//
	//		if node := children(parents[i]); node != nil {
	//			out[i] = node
	//		}
	//	}
	//
	//	if len(out) > 0 && out[0] != nil {
	//		root.Children = out
	//	}

	fmt.Printf("root: %+v\n", root)

	return root, nil

}

//
//func children(records []Record) []*Node {
//	node := make([]*Node, len(records))
//	for _, child := range records {
//		fmt.Printf("child: %+v, node: %+v\n", child, node)
//		if child.ID == 0 && child.Parent == 0 {
//			continue
//		}
//
//		node = append(node, &Node{ID: child.ID})
//	}
//
//	fmt.Printf("returning node: %+v\n", node)
//	return node
//}

//	root := &Node{}
//	todo := []*Node{root}
//	n := 1
//	for {
//		if len(todo) == 0 {
//			break
//		}
//		newTodo := []*Node(nil)
//		for _, c := range todo {
//			for _, r := range records {
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
//	if n != len(records) {
//		return nil, Mismatch{}
//	}
//	if err := chk(root, len(records)); err != nil {
//		return nil, err
//	}
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
