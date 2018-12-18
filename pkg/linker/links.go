package linker

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Segment is the linker internal segment type containing a document link as well
type Segment struct {
	*Node
	Text string `json:"text"`
}

// Node in the link graph
type Node struct {
	hash string
	Doc  string `json:"doc"`
	Seg  int    `json:"seg"`
}

func (n *Node) String() string {
	return fmt.Sprintf("%s-%d", n.Doc, n.Seg)
}

// Hash returns a unique reproducible hash for each node
func (n *Node) Hash() string {
	if n.hash == "" {
		h := sha256.New()
		h.Write([]byte(fmt.Sprintf("%s-%d", n.Doc, n.Seg)))
		n.hash = hex.EncodeToString(h.Sum(nil))
	}

	return n.hash
}

// Edge between two nodes
type Edge struct {
	hash   string
	Source *Node   `json:"source"`
	Target *Node   `json:"target"`
	Weight float32 `json:"weight"`
}

func (e *Edge) String() string {
	return fmt.Sprintf("%s-%s#%v", e.Source, e.Target, e.Weight)
}

// Hash returns a unique reproducible hash for each edge
func (e *Edge) Hash() string {
	if e.hash == "" {
		h := sha256.New()
		h.Write([]byte(fmt.Sprintf("%s-%s", e.Source, e.Target)))
		e.hash = hex.EncodeToString(h.Sum(nil))
	}

	return e.hash
}
