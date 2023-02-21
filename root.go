package ls

import (
	"context"
	"encoding/json"
	"log"

	"github.com/creachadair/jrpc2"
	"github.com/josa42/go-ls/lsp"
)

type RootHandler struct {
	server *Server
}

func unmarshalParams(r *jrpc2.Request, v interface{}) error {
	raw := make(map[string]interface{})
	r.UnmarshalParams(&raw)

	b, _ := json.Marshal(raw)

	return json.Unmarshal(b, v)
}

func (h *RootHandler) Initialize(fn func(Server, context.Context, lsp.InitializeParams) (lsp.InitializeResult, error)) {
	h.server.register("initialize", func(ctx context.Context, r *jrpc2.Request) (interface{}, error) {
		p := lsp.InitializeParams{}
		// FIXME Always throws error
		if err := unmarshalParams(r, &p); err != nil {
			log.Printf("%v", err)
			return nil, err
		}
		return fn(*h.server, ctx, p)
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
