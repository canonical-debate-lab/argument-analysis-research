package linker

// DocumentLinks stores an array of Links aggregating all the links for the segments inside it
type DocumentLinks []Links

// Links stores an array of links for a document
type Links []Link

// Link stores the target links id and the links distance
type Link struct {
	Document int
	Segment  int
	Dist     float32
}
