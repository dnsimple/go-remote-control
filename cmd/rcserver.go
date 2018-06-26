package main

import (
	"fmt"
	"github.com/aetrion/go-remote-control"
	"github.com/codegangsta/martini"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	version         = "1.2.0"
	httpBindAddress = os.Getenv("HTTP_BIND_ADDRESS")
	httpBindPort    = os.Getenv("HTTP_BIND_PORT")
	httpCertFile    = os.Getenv("HTTP_CERT_FILE")
	httpKeyFile     = os.Getenv("HTTP_KEY_FILE")
	authToken       = os.Getenv("HTTP_AUTH_TOKEN")
)

func authorize(log *log.Logger, req *http.Request, res http.ResponseWriter) {
	auth := req.Header.Get("Authorization")
	if auth == "" || auth != authToken {
		log.Printf("Authorization failed")
		res.WriteHeader(401)
	}

}

func httpRoot(req *http.Request, res http.ResponseWriter) {
	fmt.Fprintf(res, "Version %s\r\n", version)
}

func httpStatus(req *http.Request, res http.ResponseWriter) {
	err := grc.RunStatus(res)
	if err != nil {
		fmt.Fprintf(res, "Error: %q\r\n", err)
		return
	}
}

func httpStart(req *http.Request, res http.ResponseWriter) {
	err := grc.RunStart(res)
	if err != nil {
		fmt.Fprintf(res, "Error: %q\r\n", err)
		return
	}
}

func httpStop(req *http.Request, res http.ResponseWriter) {
	err := grc.RunStop(res)
	if err != nil {
		fmt.Fprintf(res, "Error: %q\r\n", err)
		return
	}
}

func httpUpdate(req *http.Request, res http.ResponseWriter) {
	err := grc.RunUpdate(res)
	if err != nil {
		fmt.Fprintf(res, "Error: %q\r\n", err)
		return
	}
}

// make a martini, shaken not stirred
func makeMartini() *martini.ClassicMartini {
	m := martini.Classic()

	m.Get("/", authorize, httpRoot)
	m.Get("/status", authorize, httpStatus)
	m.Get("/stop", authorize, httpStop)
	m.Get("/start", authorize, httpStart)
	m.Get("/update", authorize, httpUpdate)

	return m
}

func main() {
	// add services to every request
	m := makeMartini()
	// m.Map(object)

	// listen for requests
	httpAddressAndPort := net.JoinHostPort(httpBindAddress, httpBindPort)
	log.Println("listening on ", httpAddressAndPort)
	err := http.ListenAndServeTLS(httpAddressAndPort, httpCertFile, httpKeyFile, m)
	if err != nil {
		log.Printf("Failed to start HTTP server: %s", err)
	}
}
