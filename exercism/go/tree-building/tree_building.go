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

// declare mappers to be usable across functions
var mappers map[int][]*Node

func checkRecord(record Record, ids []bool) error {
	if record.ID >= len(ids) || ids[record.ID] {
		return fmt.Errorf("node already seen: %d", record.ID)
	}

	if record.ID <= record.Parent && record.ID != 0 {
		return fmt.Errorf("parent cannot be greater than ID, parent: %d, ID: %d", record.Parent, record.ID)
	}

	if record.ID == 0 && record.Parent != 0 {
		return fmt.Errorf("root not cannot have a parent")
	}

	return nil

}

// Build stitches together the node tree
func Build(records []Record) (*Node, error) {
	mappers = make(map[int][]*Node)

	if len(records) == 0 {
		return nil, nil
	}

	ids := make([]bool, len(records))

	for _, record := range records {
		if err := checkRecord(record, ids); err != nil {
			return nil, err
		}

		ids[record.ID] = true

		if record.ID != 0 {
			mappers[record.Parent] = append(mappers[record.Parent], &Node{ID: record.ID})
		}
	}

	for idx, id := range ids {
		if !id {
			return nil, fmt.Errorf("did not receive expected id: %d", idx)
		}
	}

	// generate the root
	root := &Node{}

	parents := make([]int, len(mappers))
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
