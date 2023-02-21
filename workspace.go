package ls

import (
	"context"
	"log"

	"github.com/creachadair/jrpc2"
	"github.com/josa42/go-ls/lsp"
)

type WorkspaceHandler struct {
	server *Server
}

func (h *WorkspaceHandler) DidChangeConfiguration(func(RequestContext)) {}
func (h *WorkspaceHandler) DidChangeWatchedFiles(func(RequestContext))  {}
func (h *WorkspaceHandler) ExecuteCommand(fn func(RequestContext, lsp.ExecuteCommandParams) error) {
	h.server.register("workspace/executeCommand", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := lsp.ExecuteCommandParams{}
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
func (h *WorkspaceHandler) Symbol(func(RequestContext)) {}
