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

	return s
}

type Server struct {
	jrpc2Server *jrpc2.Server
	handlers    handler.Map

	VerboseLogging bool

	State *files.State

	Root         RootHandler
	TextDocument TextDocumentHandler
}

func (s *Server) register(method string, handlerFn handler.Func) {
	if s.VerboseLogging {
		log.Printf("register: %s", method)
	}
	s.handlers[method] = logMiddleware(s.VerboseLogging, handlerFn)
}

func logMiddleware(verbose bool, h handler.Func) handler.Func {
	return func(c context.Context, r *jrpc2.Request) (interface{}, error) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered: '%v'", r)
			}
		}()

		resp, err := h(c, r)

		if verbose {
			s, _ := json.MarshalIndent(resp, "", "  ")
			log.Printf("=> (response) %s", s)

			if err != nil {
				log.Printf("=> (error)  %s", err.Error())
			}
		}

		return resp, err
	}
}

func (s *Server) push(ctx context.Context, method string, p interface{}) error {
	ps, _ := json.MarshalIndent(p, "", "  ")
	log.Printf("=> (push) %s", ps)

	return s.jrpc2Server.Push(ctx, method, p)
}

func (s *Server) Start() {
	s.jrpc2Server = jrpc2.NewServer(s.handlers, &jrpc2.ServerOptions{
		AllowPush: true,
		CheckRequest: func(ctx context.Context, req *jrpc2.Request) error {
			d := map[string]interface{}{}
			req.UnmarshalParams(&d)

			if s.VerboseLogging {
				ps, _ := json.MarshalIndent(d, "", "  ")
				log.Printf("<= (request) %s | %s", req.Method(), ps)
			} else {
				log.Printf("<= (request) %s", req.Method())
			}

			return nil
		},
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
