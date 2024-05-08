package main

type Server struct {
}

/*
Build a new api server instance.
*/
func New() *Server {
	return &Server{}
}

/*
Initialize a new server instance.
*/
func (s *Server) Init() error {
	return nil
}
