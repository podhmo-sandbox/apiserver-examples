package app

func (s *Server) routes() {
	s.router.HandleFunc("/cellar/accounts", s.handleAccountList())
}
