package ls

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/channel"
	"github.com/creachadair/jrpc2/handler"
	"github.com/josa42/go-ls/files"
)

func New() *Server {
	var handlers = handler.Map{}

	s := &Server{
		handlers: handlers,
		State:    files.NewState(),
	}

	s.Root = RootHandler{server: s}
	s.TextDocument = TextDocumentHandler{server: s}
	s.Workspace = WorkspaceHandler{server: s}

	return s
}

type Server struct {
	jrpc2Server *jrpc2.Server
	handlers    handler.Map

	VerboseLogging bool

	State *files.State

	Root         RootHandler
	TextDocument TextDocumentHandler
	Workspace    WorkspaceHandler
}

func (s *Server) register(method string, handlerFn handler.Func) {
	if s.VerboseLogging {
		log.Printf("register: %s", method)
	}
	s.handlers[method] = handlerFn
}

func (s *Server) push(ctx context.Context, method string, p interface{}) error {
	ps, _ := json.MarshalIndent(p, "", "  ")
	if s.VerboseLogging {
		log.Printf("=> (push) %s", ps)
	}

	return s.jrpc2Server.Notify(ctx, method, p)
}

func (s *Server) Start() {
	s.jrpc2Server = jrpc2.NewServer(s.handlers, &jrpc2.ServerOptions{
		AllowPush: true,
		RPCLog:    s,
	})
	s.jrpc2Server.Start(channel.Header("")(os.Stdin, os.Stdout))
}

func (s *Server) Wait() error {
	return s.jrpc2Server.Wait()
}

func (s *Server) StartAndWait() error {
	s.Start()
	return s.Wait()
}

func (s *Server) LogRequest(ctx context.Context, req *jrpc2.Request) {
	d := map[string]interface{}{}
	req.UnmarshalParams(&d)

	ps, _ := json.MarshalIndent(d, "", "  ")
	if s.VerboseLogging {
		log.Printf("<= (request) %s | %s", req.Method(), ps)
	} else {
		log.Printf("<= (request) %s", req.Method())
	}
}

func (s *Server) LogResponse(ctx context.Context, resp *jrpc2.Response) {
	t, _ := json.MarshalIndent(resp, "", "  ")
	if s.VerboseLogging {
		log.Printf("=> (response) %s", t)
	}
}
