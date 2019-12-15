package lsp

import "github.com/sourcegraph/go-lsp"

type WorkspaceEdit struct {
	/**
	 * Holds changes to existing resources.
	 */
	Changes map[lsp.DocumentURI][]lsp.TextEdit `json:"changes"`
}
