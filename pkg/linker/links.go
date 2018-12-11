package linker

import (
	"fmt"
)

// DocumentLinks stores an array of Links aggregating all the links for the segments inside it
type DocumentLinks []Links

// Links stores an array of links for a document
type Links []Link

// Link stores the target links id and the links distance
type Link struct {
	Document string  `json:"document"`
	Segment  int     `json:"segment"`
	Dist     float32 `json:"dist"`
}

// Segment is the linker internal segment type containing a document link as well
type Segment struct {
	*Node
	Text string `json:"text"`
}

// Node in the link graph
type Node struct {
	Doc string `json:"doc"`
	Seg int    `json:"seg"`
}

func (n *Node) String() string {
	return fmt.Sprintf("%s-%d", n.Doc, n.Seg)
}

// Edge between two nodes
type Edge struct {
	Source *Node   `json:"source"`
	Target *Node   `json:"target"`
	Weight float32 `json:"weight"`
}

func (e *Edge) String() string {
	return fmt.Sprintf("%s-%s#%v", e.Source, e.Target, e.Weight)
}
