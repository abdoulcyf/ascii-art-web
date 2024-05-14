package servers

func NewAPIServer(listener string) *server {
	return &server{
		listener: listener,
	}
}
