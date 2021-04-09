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

	if record.ID == record.Parent && record.ID != 0 {
		return Mismatch{}
	}

	if record.Parent > record.ID {
		return Mismatch{}
	}

	for _, child := range mappers[record.Parent] {
		if record.ID == child.ID {
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

	ids := make([]int, len(records))
	for idx, record := range records {
		ids[idx] = record.ID
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
	}

	if !rootSent {
		return nil, Mismatch{}
	}

	// generate the root
	root := &Node{}

	sort.Ints(ids)
	if len(ids) == 1 {
		return root, nil
	}

	for i := 1; i < len(ids); i++ {
		if ids[i]-ids[i-1] > 1 {
			return nil, fmt.Errorf("y")
		}
	}

	parents := make([]int, 0, len(mappers))
	for parent := range mappers {
		parents = append(parents, parent)
		sort.Slice(mappers[parent], func(i, j int) bool {
			return mappers[parent][i].ID < mappers[parent][j].ID
		})
	}

	sort.Ints(parents)

	for parent := range parents {
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
