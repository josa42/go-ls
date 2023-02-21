package lsp

type WorkspaceEdit struct {
	/**
	 * Holds changes to existing resources.
	 */
	Changes map[DocumentURI][]TextEdit `json:"changes"`
}
