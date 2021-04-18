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
var mappers map[int]*Node

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
	mappers = make(map[int]*Node)

	if len(records) == 0 {
		return nil, nil
	}

	ids := make([]bool, len(records))

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	for _, record := range records {
		if err := checkRecord(record, ids); err != nil {
			return nil, err
		}

		ids[record.ID] = true
		if _, ok := mappers[record.ID]; !ok {
			mappers[record.ID] = &Node{ID: record.ID}
		}

		if record.ID == 0 {
			continue
		}

		if _, ok := mappers[record.Parent]; !ok {
			mappers[record.Parent] = &Node{ID: record.Parent, Children: []*Node{mappers[record.ID]}}
		} else {
			mappers[record.Parent].Children = append(mappers[record.Parent].Children, mappers[record.ID])
		}
	}

	for idx, id := range ids {
		if !id {
			return nil, fmt.Errorf("did not receive expected id: %d", idx)
		}
	}

	return mappers[0], nil
}
