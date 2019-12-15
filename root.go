package ls

import (
	"context"

	"github.com/creachadair/jrpc2"
	"github.com/sourcegraph/go-lsp"
)

type RootHandler struct {
	server *Server
}

func (h *RootHandler) Initialize(fn func(Server, context.Context, lsp.InitializeParams) (lsp.InitializeResult, error)) {
	h.server.register("initialize", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := lsp.InitializeParams{}
		// FIXME Always throws error
		// if err := r.UnmarshalParams(&p); err != nil {
		// 	log.Printf("%v", err)
		// 	return nil, err
		// }
		return fn(*h.server, ctx, p)
	})
}

// 		"initialized":                    noop,
// 		"shutdown":                       noop,
// 		"exit":                           noop,
// 		"$/cancelRequest":                noop,
