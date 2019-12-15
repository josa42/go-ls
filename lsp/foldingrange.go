package lsp

import "github.com/sourcegraph/go-lsp"

type FoldingRangeParams struct {
	// The text document.
	TextDocument lsp.TextDocumentIdentifier
}

type FoldingRangeKind string

// Enum of known range kinds
const (
	FRKComment FoldingRangeKind = "comment" // Folding range for a comment
	FRKImports FoldingRangeKind = "imports" // Folding range for a imports or includes
	FRKRegion  FoldingRangeKind = "region"  // Folding range for a region (e.g. `#region`)
)

// Represents a folding range.
type FoldingRange struct {
	StartLine      int              `json:"startLine"`                // The zero-based line number from where the folded range starts.
	StartCharacter *int             `json:"startCharacter,omitempty"` // The zero-based character offset from where the folded range starts. If not defined, defaults to the length of the start line.
	EndLine        int              `json:"endLine"`                  // The zero-based line number where the folded range ends.
	EndCharacter   *int             `json:"endCharacter,omitempty"`   // The zero-based character offset before the folded range ends. If not defined, defaults to the length of the end line.
	Kind           FoldingRangeKind `json:"kind"`
}
