package handlers


func LoadingServer(){
	server := NewAPIServer(":8080")
	server.RunServer()
}