package ls

import (
	"context"
	"fmt"
	"log"

	"github.com/creachadair/jrpc2"
	"github.com/josa42/go-ls/lsp"
)

type CommandFn func(ctx RequestContext, args []interface{}) error

type WorkspaceHandler struct {
	server   *Server
	commands map[string]CommandFn
}

func NewWorkspaceHandler(server *Server) WorkspaceHandler {
	return WorkspaceHandler{server: server, commands: map[string]CommandFn{}}
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

		reqCtx := RequestContext{
			Server:  *h.server,
			Context: ctx,
		}

		if fn, ok := h.commands[p.Command]; ok {
			return nil, fn(reqCtx, p.Arguments)
		}

		return nil, fn(reqCtx, p)
	})
}

func (h *WorkspaceHandler) RegisterCommand(command string, fn CommandFn) {
	if !h.server.has("workspace/executeCommand") {
		h.ExecuteCommand(func(ctx RequestContext, p lsp.ExecuteCommandParams) error {
			return fmt.Errorf("Command \"%s\" not found", p.Command)
		})
	}

	h.commands[command] = fn
}

func (h *WorkspaceHandler) Symbol(func(RequestContext)) {}
