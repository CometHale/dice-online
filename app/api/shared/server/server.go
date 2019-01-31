package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Server stores relevant information about the server
type Server struct {
	Hostname  string `json:"Hostname"`  // Server name
	UseHTTP   bool   `json:"UseHTTP"`   // Listen on HTTP
	UseHTTPS  bool   `json:"UseHTTPS"`  // Listen on HTTPS
	HTTPPort  int    `json:"HTTPPort"`  // HTTP port
	HTTPSPort int    `json:"HTTPSPort"` // HTTPS port
	CertFile  string `json:"CertFile"`  // HTTPS certificate
	KeyFile   string `json:"KeyFile"`   // HTTPS private key
}

// Run starts the HTTP and/or HTTPS listener
func Run(httpHandlers http.Handler, httpsHandlers http.Handler, server Server) {

	if server.UseHTTP && server.UseHTTPS {
		go func() {
			startHTTPS(httpsHandlers, server)
		}()

		startHTTP(httpHandlers, server)
	} else if server.UseHTTP {
		startHTTP(httpHandlers, server)
	} else if server.UseHTTPS {
		startHTTPS(httpsHandlers, server)
	} else {
		log.Println("Configuration file does not specify a listener to start.")
	}
}

// startHTTP starts the HTTP listener
func startHTTP(handlers http.Handler, server Server) {
	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), "Running HTTP "+httpAddress(server))
	// Start the HTTP listener
	log.Fatal(http.ListenAndServe(httpAddress(server), handlers))
}

// startHTTPS starts the HTTPS listener
func startHTTPS(handlers http.Handler, server Server) {
	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), "Running HTTPS "+httpsAddress(server))

	// Start the HTTPS listener
	log.Fatal(http.ListenAndServeTLS(httpsAddress(server), server.CertFile, server.KeyFile, handlers))
}

// httpAddress returns the HTTP address
func httpAddress(server Server) string {
	return server.Hostname + ":" + fmt.Sprintf("%d", server.HTTPPort)
}

// httpsAddress returns the HTTPS address
func httpsAddress(server Server) string {
	return server.Hostname + ":" + fmt.Sprintf("%d", server.HTTPSPort)
}
