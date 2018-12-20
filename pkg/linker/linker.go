package linker

import (
	"context"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/document"
)

// Linker stores all documents and compares them for building a link matrix
type Linker interface {
	Run(context.Context) error
	InsertDocument(ctx context.Context, doc *document.Document) error
	ListDocuments(ctx context.Context) []*document.Document
	ListLinks(ctx context.Context) []*Edge
}

// Metadata for storing linker settings
type Metadata struct {
	ID        string  `json:"id"`
	Rater     string  `json:"rater"`
	Threshold float32 `json:"threshold"`
}

// Storage defines the minimal interface for a linker persistence layer
type Storage interface {
	SetMetadata(ctx context.Context, meta *Metadata)
	Metadata(ctx context.Context) (*Metadata, error)

	InsertDocument(ctx context.Context, doc *document.Document)
	InsertSegment(ctx context.Context, doc *Segment)
	InsertLink(ctx context.Context, doc *Edge)

	Documents(ctx context.Context) (map[string]*document.Document, error)
	Segments(ctx context.Context) (map[string]*Segment, error)
	Links(ctx context.Context) (map[string]*Edge, error)
}
