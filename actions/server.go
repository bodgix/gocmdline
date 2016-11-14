package actions

import "fmt"

// Server an action for server
type Server struct {
	Name string
}

// Do Action interface implementation
func (s *Server) Do() {
	fmt.Println("Server", s.Name, "doing an action")
}
