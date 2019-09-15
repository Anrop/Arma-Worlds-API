package server

func (s *Server) setupRoutes() {
	s.router.Handle("/worlds", s.worlds())
}
