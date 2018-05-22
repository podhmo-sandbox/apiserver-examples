package app

import (
	"net/http"
	"github.com/podhmo/apiserver-examples/vanilla/app/middleware"
	"github.com/podhmo/apiserver-examples/vanilla/store"
)

// Server :
type Server struct {
	router *http.ServeMux
	db     *store.DB
}

// NewServer :
func NewServer(db *store.DB) *Server {
	s := &Server{
		router: &http.ServeMux{},
		db:     db,
	}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h := middleware.Recover(s.router.ServeHTTP)
	h(w, r)
}
