package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/rpc"
)

func main() {

	cert, err := tls.LoadX509KeyPair("client1.crt", "client.key")
	if err != nil {
		log.Fatalf("client: loadkeys: %s", err)
	}

	if len(cert.Certificate) != 2 {
		log.Fatal("client.crt should have 2 concatenated certificates: client + CA")
	}

	ca, err := x509.ParseCertificate(cert.Certificate[1])
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(ca)

	config := tls.Config{
		Certificates: []tls.Certificate{cert},
		//InsecureSkipVerify: true,
		RootCAs: certPool,
	}

	conn, err := tls.Dial("tcp", "localhost:2012", &config)
	if err != nil {
		log.Fatalf("client: dial: %s", err)
	}
	defer conn.Close()
	log.Println("client: connected to: ", conn.RemoteAddr())
	rpcClient := rpc.NewClient(conn)
	var reply int
	if err := rpcClient.Call("MyServer.Sum", &ArgsSum{Item1: 13, Item2: 5}, &reply); err != nil {
		log.Fatal("Failed to call RPC", err)
	}
	log.Printf("Returned result is %d", reply)
}

type ArgsSum struct {
	Item1, Item2 int
}
