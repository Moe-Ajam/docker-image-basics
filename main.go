package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    port,
		Handler: mux,
	}
	fmt.Printf("Listening on port: %s\n", port)

	mux.HandleFunc("GET /", handlePage)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	const page = `<html>
<head></head>
<body>
	<p> Hello from Docker! I'm a Go server. </p>
</body>
</html>
`
	w.Write([]byte(page))
}
