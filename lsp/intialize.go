package lsp

type InitializeParams struct {
	ProcessID             int                `json:"processId,omitempty"`
	RootURI               DocumentURI        `json:"rootUri,omitempty"`
	InitializationOptions interface{}        `json:"initializationOptions,omitempty"`
	Capabilities          ClientCapabilities `json:"capabilities"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities,omitempty"`
}

type ServerCapabilities struct {
	CodeActionProvider               bool                             `json:"codeActionProvider,omitempty"`
	CodeLensProvider                 *CodeLensOptions                 `json:"codeLensProvider,omitempty"`
	CompletionProvider               *CompletionOptions               `json:"completionProvider,omitempty"`
	DefinitionProvider               bool                             `json:"definitionProvider,omitempty"`
	DocumentFormattingProvider       bool                             `json:"documentFormattingProvider,omitempty"`
	DocumentHighlightProvider        bool                             `json:"documentHighlightProvider,omitempty"`
	DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingOptions `json:"documentOnTypeFormattingProvider,omitempty"`
	DocumentRangeFormattingProvider  bool                             `json:"documentRangeFormattingProvider,omitempty"`
	DocumentSymbolProvider           bool                             `json:"documentSymbolProvider,omitempty"`
	ExecuteCommandProvider           *ExecuteCommandOptions           `json:"executeCommandProvider,omitempty"`
	FoldingRangeProvider             bool                             `json:"foldingRangeProvider,omitempty"`
	HoverProvider                    bool                             `json:"hoverProvider,omitempty"`
	ImplementationProvider           bool                             `json:"implementationProvider,omitempty"`
	ReferencesProvider               bool                             `json:"referencesProvider,omitempty"`
	RenameProvider                   bool                             `json:"renameProvider,omitempty"`
	SignatureHelpProvider            *SignatureHelpOptions            `json:"signatureHelpProvider,omitempty"`
	TextDocumentSync                 *TextDocumentSyncOptions         `json:"textDocumentSync,omitempty"`
	TypeDefinitionProvider           bool                             `json:"typeDefinitionProvider,omitempty"`
	WorkspaceSymbolProvider          bool                             `json:"workspaceSymbolProvider,omitempty"`
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
