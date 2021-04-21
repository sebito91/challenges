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

// Build stitches together the node tree
func Build(records []Record) (*Node, error) {
	mappers = make(map[int]*Node)

	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	for idx, record := range records {
		if record.ID != idx || record.Parent > record.ID || record.ID > 0 && record.Parent == record.ID {
			return nil, fmt.Errorf("not in sequence or has bad parent: %v", record)
		}

		if _, ok := mappers[record.ID]; !ok {
			mappers[record.ID] = &Node{ID: record.ID}
		}

		if record.ID == 0 {
			continue
		}

		mappers[record.Parent].Children = append(mappers[record.Parent].Children, mappers[record.ID])
	}

	return mappers[0], nil
}
