package servers

import "net/http"


func LoadingServerS() {

	mux := http.NewServeMux()

	mux.HandleFunc(homePagePath, MainHandler)

	http.ListenAndServe(portAddressNumber, mux)
}
