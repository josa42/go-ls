package ls

import (
	"context"

	"github.com/creachadair/jrpc2"
	"github.com/josa42/go-ls/lsp"
)

type RootHandler struct {
	server *Server
}

func (h *RootHandler) Initialize(fn func(Server, context.Context, lsp.InitializeParams) (lsp.InitializeResult, error)) {
	h.server.register("initialize", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {

		p := lsp.InitializeParams{}

		err := r.UnmarshalParams(&p)
		if err != nil {
			return nil, err
		}
		res, err := fn(*h.server, ctx, p)

		if len(h.server.Workspace.commands) > 0 {
			if res.Capabilities.ExecuteCommandProvider == nil {
				res.Capabilities.ExecuteCommandProvider = &lsp.ExecuteCommandOptions{}
			}

			if res.Capabilities.ExecuteCommandProvider.Commands == nil {
				res.Capabilities.ExecuteCommandProvider.Commands = []string{}
			}

			for cmd := range h.server.Workspace.commands {
				found := false

				for _, registered := range res.Capabilities.ExecuteCommandProvider.Commands {
					if registered == cmd {
						found = true
						break
					}
				}

				if !found {
					res.Capabilities.ExecuteCommandProvider.Commands = append(res.Capabilities.ExecuteCommandProvider.Commands, cmd)
				}

			}
		}

		return res, err
	})
}

func (h *RootHandler) Shutdown(fn func(RequestContext) error) {
	h.server.register("shutdown", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		return nil, fn(RequestContext{
			Server:  *h.server,
			Context: ctx,
		})
	})
}

// 		"initialized":                    noop,
// 		"shutdown":                       noop,
// 		"exit":                           noop,
// 		"$/cancelRequest":                noop,
