package rest

import "net/http"

type Serverless struct {
	server *Server
}

func NewServerless(server *Server) (*Serverless, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Serverless) Serve(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
