package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	certFile = flag.String("cert", "server.crt", "A PEM eoncoded certificate file.")
)

func main() {
	flag.Parse()

	certPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(*certFile)

	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(pem)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	// Do GET something
	resp, err := client.Get("https://localhost:4443/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Dump response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
}
