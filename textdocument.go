package ls

import (
	"context"
	"log"

	"github.com/creachadair/jrpc2"
	"github.com/josa42/go-ls/lsp"
	golsp "github.com/sourcegraph/go-lsp"
)

type TextDocumentHandler struct {
	server *Server
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

func (h *TextDocumentHandler) Completion(fn func(RequestContext, golsp.CompletionParams) (golsp.CompletionList, error)) {
	// TODO enable completion in capabilities

	h.server.register("textDocument/completion", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := golsp.CompletionParams{}
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

func (h *TextDocumentHandler) DocumentSymbol(fn func(RequestContext, golsp.DocumentSymbolParams) ([]golsp.SymbolInformation, error)) {
	h.server.register("textDocument/documentSymbol", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := golsp.DocumentSymbolParams{}
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

func (h *TextDocumentHandler) Formatting(fn func(RequestContext, golsp.DocumentFormattingParams) ([]golsp.TextEdit, error)) {
	h.server.register("textDocument/formatting", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := golsp.DocumentFormattingParams{}
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

func (h *TextDocumentHandler) Hover(fn func(RequestContext, golsp.TextDocumentPositionParams) (*golsp.Hover, error)) {
	h.server.register("textDocument/hover", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := golsp.TextDocumentPositionParams{}
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

func (h *TextDocumentHandler) DidChange(fn func(RequestContext, golsp.DidChangeTextDocumentParams) error) {
	h.server.register("textDocument/didChange", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := golsp.DidChangeTextDocumentParams{}
		if err := r.UnmarshalParams(&p); err != nil {
			log.Printf("%v", err)
			return nil, err
		}
		return nil, fn(RequestContext{
			Server:  *h.server,
			Context: ctx,
		}, p)
	})
}

func (h *TextDocumentHandler) DidClose(fn func(RequestContext, golsp.DidCloseTextDocumentParams) error) {
	h.server.register("textDocument/didClose", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := golsp.DidCloseTextDocumentParams{}
		if err := r.UnmarshalParams(&p); err != nil {
			log.Printf("%v", err)
			return nil, err
		}
		return nil, fn(RequestContext{
			Server:  *h.server,
			Context: ctx,
		}, p)
	})
}

func (h *TextDocumentHandler) DidOpen(fn func(RequestContext, golsp.DidOpenTextDocumentParams) error) {
	h.server.register("textDocument/didOpen", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := golsp.DidOpenTextDocumentParams{}
		if err := r.UnmarshalParams(&p); err != nil {
			log.Printf("%v", err)
			return nil, err
		}
		return nil, fn(RequestContext{
			Server:  *h.server,
			Context: ctx,
		}, p)
	})
}

func (h *TextDocumentHandler) PublishDiagnostics(ctx context.Context, vs golsp.PublishDiagnosticsParams) error {
	return h.server.push(ctx, "textDocument/publishDiagnostics", vs)
}
