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
