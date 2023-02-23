package ls

import (
	"context"
	"log"

	"github.com/creachadair/jrpc2"
	"github.com/josa42/go-ls/lsp"
)

type TextDocumentHandler struct {
	server          *Server
	listenToChanges bool
}

func (h *TextDocumentHandler) RegisterChangesListener() {
	h.listenToChanges = true

	if !h.server.has("textDocument/didChange") {
		h.DidChange(func(RequestContext, lsp.DidChangeTextDocumentParams) error {
			return nil
		})
	}

	if !h.server.has("textDocument/didClose") {
		h.DidClose(func(RequestContext, lsp.DidCloseTextDocumentParams) error {
			return nil
		})
	}

	if !h.server.has("textDocument/didOpen") {
		h.DidOpen(func(RequestContext, lsp.DidOpenTextDocumentParams) error {
			return nil
		})
	}
}

func (h *TextDocumentHandler) CodeAction(fn func(RequestContext, lsp.CodeActionParams) ([]lsp.CodeAction, error)) {
	h.server.register("textDocument/codeAction", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := lsp.CodeActionParams{}
		if err := r.UnmarshalParams(&p); err != nil {
			log.Printf("%v", err)
			return nil, err
		}

		return fn(RequestContext{
			Server:  *h.server,
			Context: ctx,
		}, p)
	})
}

func (h *TextDocumentHandler) CodeLens(func(RequestContext)) {

}

func (h *TextDocumentHandler) ColorPresentation(func(RequestContext)) {

}

func (h *TextDocumentHandler) Completion(fn func(RequestContext, lsp.CompletionParams) (lsp.CompletionList, error)) {
	// TODO enable completion in capabilities

	h.server.register("textDocument/completion", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := lsp.CompletionParams{}
		if err := r.UnmarshalParams(&p); err != nil {
			log.Printf("%v", err)
			return nil, err
		}
		return fn(RequestContext{
			Server:  *h.server,
			Context: ctx,
		}, p)
	})
}

func (h *TextDocumentHandler) Declaration(func(RequestContext)) {

}

func (h *TextDocumentHandler) Definition(func(RequestContext)) {

}

func (h *TextDocumentHandler) DocumentColor(func(RequestContext)) {

}

func (h *TextDocumentHandler) DocumentHighlight(func(RequestContext)) {

}

func (h *TextDocumentHandler) DocumentLink(func(RequestContext)) {

}

func (h *TextDocumentHandler) DocumentSymbol(fn func(RequestContext, lsp.DocumentSymbolParams) ([]lsp.DocumentSymbol, error)) {
	h.server.register("textDocument/documentSymbol", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := lsp.DocumentSymbolParams{}
		if err := r.UnmarshalParams(&p); err != nil {
			log.Printf("%v", err)
			return nil, err
		}

		return fn(RequestContext{
			Server:  *h.server,
			Context: ctx,
		}, p)
	})

}

func (h *TextDocumentHandler) FoldingRange(fn func(RequestContext, lsp.FoldingRangeParams) ([]lsp.FoldingRange, error)) {
	h.server.register("textDocument/foldingRange", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := lsp.FoldingRangeParams{}
		if err := r.UnmarshalParams(&p); err != nil {
			log.Printf("%v", err)
			return nil, err
		}

		return fn(RequestContext{
			Server:  *h.server,
			Context: ctx,
		}, p)
	})

}

func (h *TextDocumentHandler) Formatting(fn func(RequestContext, lsp.DocumentFormattingParams) ([]lsp.TextEdit, error)) {
	h.server.register("textDocument/formatting", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := lsp.DocumentFormattingParams{}
		if err := r.UnmarshalParams(&p); err != nil {
			log.Printf("%v", err)
			return nil, err
		}
		return fn(RequestContext{
			Server:  *h.server,
			Context: ctx,
		}, p)
	})
}

func (h *TextDocumentHandler) Hover(fn func(RequestContext, lsp.TextDocumentPositionParams) (*lsp.Hover, error)) {
	h.server.register("textDocument/hover", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := lsp.TextDocumentPositionParams{}
		if err := r.UnmarshalParams(&p); err != nil {
			log.Printf("%v", err)
			return nil, err
		}
		return fn(RequestContext{
			Server:  *h.server,
			Context: ctx,
		}, p)
	})
}

func (h *TextDocumentHandler) Implementation(func(RequestContext)) {

}

func (h *TextDocumentHandler) OnTypeFormatting(func(RequestContext)) {

}

func (h *TextDocumentHandler) PrepareRename(func(RequestContext)) {

}

func (h *TextDocumentHandler) RangeFormatting(func(RequestContext)) {

}

func (h *TextDocumentHandler) References(func(RequestContext)) {

}

func (h *TextDocumentHandler) Rename(func(RequestContext)) {

}

func (h *TextDocumentHandler) SignatureHelp(func(RequestContext)) {

}

func (h *TextDocumentHandler) TypeDefinition(func(RequestContext)) {

}

func (h *TextDocumentHandler) WillSaveWaitUntil(func(RequestContext)) {

}

func (h *TextDocumentHandler) DidChange(fn func(RequestContext, lsp.DidChangeTextDocumentParams) error) {
	h.server.register("textDocument/didChange", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := lsp.DidChangeTextDocumentParams{}
		if err := r.UnmarshalParams(&p); err != nil {
			log.Printf("%v", err)
			return nil, err
		}

		if h.listenToChanges {
			h.server.State.ApplyCanges(p.TextDocument.URI, p.ContentChanges)
		}

		return nil, fn(RequestContext{
			Server:  *h.server,
			Context: ctx,
		}, p)
	})
}

func (h *TextDocumentHandler) DidClose(fn func(RequestContext, lsp.DidCloseTextDocumentParams) error) {
	h.server.register("textDocument/didClose", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := lsp.DidCloseTextDocumentParams{}
		if err := r.UnmarshalParams(&p); err != nil {
			log.Printf("%v", err)
			return nil, err
		}

		if h.listenToChanges {
			h.server.State.Remove(p.TextDocument.URI)
		}

		return nil, fn(RequestContext{
			Server:  *h.server,
			Context: ctx,
		}, p)
	})
}

func (h *TextDocumentHandler) DidOpen(fn func(RequestContext, lsp.DidOpenTextDocumentParams) error) {
	h.server.register("textDocument/didOpen", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := lsp.DidOpenTextDocumentParams{}
		if err := r.UnmarshalParams(&p); err != nil {
			log.Printf("%v", err)
			return nil, err
		}

		if h.listenToChanges {
			h.server.State.SetDocument(p.TextDocument)
		}

		return nil, fn(RequestContext{
			Server:  *h.server,
			Context: ctx,
		}, p)
	})
}

func (h *TextDocumentHandler) PublishDiagnostics(ctx context.Context, vs lsp.PublishDiagnosticsParams) error {
	return h.server.push(ctx, "textDocument/publishDiagnostics", vs)
}
