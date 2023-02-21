package main

import (
	"context"
	"log"

	"github.com/josa42/go-ls"
	"github.com/josa42/go-ls/lsp"
	golsp "github.com/sourcegraph/go-lsp"
)

func main() {
	s := ls.New()

	s.Root.Initialize(func(ls.Server, context.Context, lsp.InitializeParams) (lsp.InitializeResult, error) {
		return lsp.InitializeResult{
			Capabilities: lsp.ServerCapabilities{
				TextDocumentSync: &golsp.TextDocumentSyncOptions{
					OpenClose: true,
					Change:    golsp.TDSKFull,
				},

				CompletionProvider: &golsp.CompletionOptions{
					ResolveProvider:   false,
					TriggerCharacters: []string{" "},
				},

				DocumentFormattingProvider: true,
				DocumentSymbolProvider:     true,
			},
		}, nil

	})

	s.TextDocument.DidOpen(func(ctx ls.RequestContext, p golsp.DidOpenTextDocumentParams) error {
		ctx.Server.State.SetDocument(p.TextDocument)
		return nil
	})

	s.TextDocument.DidChange(func(ctx ls.RequestContext, p golsp.DidChangeTextDocumentParams) error {
		ctx.Server.State.ApplyCanges(p.TextDocument.URI, p.ContentChanges)
		return nil
	})

	s.TextDocument.DidClose(func(ctx ls.RequestContext, p golsp.DidCloseTextDocumentParams) error {
		ctx.Server.State.Remove(p.TextDocument.URI)
		return nil
	})

	if err := s.StartAndWait(); err != nil {
		log.Fatalf("%s", err)
	}
}
