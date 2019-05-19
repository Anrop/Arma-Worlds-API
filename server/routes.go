package server

func (s *Server) SetupRoutes() {
	s.router.Handle("/worlds", s.Worlds())
}
