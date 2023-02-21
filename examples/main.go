package main

import (
	"context"
	"log"

	"github.com/josa42/go-ls"
	"github.com/josa42/go-ls/lsp"
)

func main() {
	s := ls.New()

	s.Root.Initialize(func(ls.Server, context.Context, lsp.InitializeParams) (lsp.InitializeResult, error) {
		return lsp.InitializeResult{
			Capabilities: lsp.ServerCapabilities{
				TextDocumentSync: &lsp.TextDocumentSyncOptions{
					OpenClose: true,
					Change:    lsp.TDSKFull,
				},

				CompletionProvider: &lsp.CompletionOptions{
					ResolveProvider:   false,
					TriggerCharacters: []string{" "},
				},

				DocumentFormattingProvider: true,
				DocumentSymbolProvider:     true,
			},
		}, nil

	})

	s.TextDocument.DidOpen(func(ctx ls.RequestContext, p lsp.DidOpenTextDocumentParams) error {
		ctx.Server.State.SetDocument(p.TextDocument)
		return nil
	})

	s.TextDocument.DidChange(func(ctx ls.RequestContext, p lsp.DidChangeTextDocumentParams) error {
		ctx.Server.State.ApplyCanges(p.TextDocument.URI, p.ContentChanges)
		return nil
	})

	s.TextDocument.DidClose(func(ctx ls.RequestContext, p lsp.DidCloseTextDocumentParams) error {
		ctx.Server.State.Remove(p.TextDocument.URI)
		return nil
	})

	if err := s.StartAndWait(); err != nil {
		log.Fatalf("%s", err)
	}
}
