package main

import (
	"flag"
	"fmt"
	"net/http"

	"rsc.io/getopt"
)

func main() {
	serverDir := "bin/"
	port := 9090
	flag.StringVar(&serverDir, "serverDir", "", "Path to the server directory")
	flag.IntVar(&port, "port", 9090, "Port to listen on")

	getopt.Aliases(
		"s", "serverDir",
		"p", "port",
	)

	getopt.Parse()

	fmt.Printf("Server started at http://localhost:%d and serving files from %s\n", port, serverDir)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(serverDir)))
	if err != nil {
		fmt.Println(err)
		return
	}
}
