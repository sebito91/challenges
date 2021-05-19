// Package tree provides a tree structure for the records presented to it.
package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID int
	Parent int
}

type Node struct {
	ID int
	Children []*Node
}

type Records []Record

//Build receives the records with ID and Parent ID and creates a tree after ensuring the records adhere to the rules:
// 1. All records have a parent ID lower than their own ID, except for the root record,
// which has a parent ID that's equal to its own ID.
// 2. IDs are contiguous
// 3. The records don't create a cyclic tree
// 4. Parents will always ID higher than children (when IDs are equal it will be for root node only)
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.Sort(Records(records))
	var n *Node
	for i, record := range records {
		switch {
		case i > 0 && record.Parent == record.ID :
			return nil,errors.New("duplicate root")
		case i == 0 && ( record.Parent != record.ID  || record.Parent !=0) :
			return nil,errors.New("no valid root")
		case record.Parent > record.ID :
			return nil,errors.New("higher id parent of lower id")
		case i > 0 && record.ID == records[i-1].ID :
			return nil,errors.New("duplicate node")
		case i > 0 && record.ID != records[i-1].ID +1 :
			return nil,errors.New("nodes not contiguous")
		}

     n = insert(n, record)
	}
return n,nil
}

// Sort functions

func (r Records) Len() int {
	return len(r)
}
func (r Records) Less(i, j int) bool {
		return r[i].ID < r[j].ID
}

func (r Records) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}




func insert(n *Node, r Record) *Node {
	if n == nil {
		return &Node{r.ID, nil}
	}
	if r.Parent == n.ID {
		n.Children = append(n.Children, &Node{r.ID, nil})
		return n
	}

	for i := 0; i < len(n.Children); i++ {
		n.Children[i] = insert(n.Children[i], r)

	}
	return n
}