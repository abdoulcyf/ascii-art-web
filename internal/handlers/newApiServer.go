package handlers


func NewAPIServer(listener string) (*Server) {
	return &Server{
		listener: listener,
	}
}
