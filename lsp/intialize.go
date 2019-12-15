package lsp

import (
	"github.com/sourcegraph/go-lsp"
)

type InitializeParams struct {
	ProcessID             int                    `json:"processId,omitempty"`
	RootURI               lsp.DocumentURI        `json:"rootUri,omitempty"`
	InitializationOptions interface{}            `json:"initializationOptions,omitempty"`
	Capabilities          lsp.ClientCapabilities `json:"capabilities"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities,omitempty"`
}

type ServerCapabilities struct {
	CodeActionProvider               bool                                 `json:"codeActionProvider,omitempty"`
	CodeLensProvider                 *lsp.CodeLensOptions                 `json:"codeLensProvider,omitempty"`
	CompletionProvider               *lsp.CompletionOptions               `json:"completionProvider,omitempty"`
	DefinitionProvider               bool                                 `json:"definitionProvider,omitempty"`
	DocumentFormattingProvider       bool                                 `json:"documentFormattingProvider,omitempty"`
	DocumentHighlightProvider        bool                                 `json:"documentHighlightProvider,omitempty"`
	DocumentOnTypeFormattingProvider *lsp.DocumentOnTypeFormattingOptions `json:"documentOnTypeFormattingProvider,omitempty"`
	DocumentRangeFormattingProvider  bool                                 `json:"documentRangeFormattingProvider,omitempty"`
	DocumentSymbolProvider           bool                                 `json:"documentSymbolProvider,omitempty"`
	ExecuteCommandProvider           *lsp.ExecuteCommandOptions           `json:"executeCommandProvider,omitempty"`
	FoldingRangeProvider             bool                                 `json:"foldingRangeProvider,omitempty"`
	HoverProvider                    bool                                 `json:"hoverProvider,omitempty"`
	ImplementationProvider           bool                                 `json:"implementationProvider,omitempty"`
	ReferencesProvider               bool                                 `json:"referencesProvider,omitempty"`
	RenameProvider                   bool                                 `json:"renameProvider,omitempty"`
	SignatureHelpProvider            *lsp.SignatureHelpOptions            `json:"signatureHelpProvider,omitempty"`
	TextDocumentSync                 *lsp.TextDocumentSyncOptions         `json:"textDocumentSync,omitempty"`
	TypeDefinitionProvider           bool                                 `json:"typeDefinitionProvider,omitempty"`
	WorkspaceSymbolProvider          bool                                 `json:"workspaceSymbolProvider,omitempty"`
	// 	documentLinkProvider?: DocumentLinkOptions;
	// 	colorProvider?: boolean | ColorProviderOptions | (ColorProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions);
	// 	foldingRangeProvider?: boolean | FoldingRangeProviderOptions | (FoldingRangeProviderOptions & TextDocumentRegistrationOptions & StaticRegistrationOptions);
	// 	declarationProvider?: boolean | (TextDocumentRegistrationOptions & StaticRegistrationOptions);
	// 	workspace?: {
	// 		workspaceFolders?: {
	// 			supported?: boolean;
	// 			changeNotifications?: string | boolean;
	// 		}
	// 	}
	// 	experimental?: any;
}
