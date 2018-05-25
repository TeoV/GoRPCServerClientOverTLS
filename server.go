package main

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net"
	"net/rpc"
)

func main() {

	cert, err := tls.LoadX509KeyPair("server1.crt", "server.key")
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}

	if len(cert.Certificate) != 2 {
		log.Fatal("server.crt should have 2 concatenated certificates: server + CA")
	}

	ca, err := x509.ParseCertificate(cert.Certificate[1])
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(ca)
	config := tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	config.Rand = rand.Reader
	service := "127.0.0.1:2012"
	listener, err := tls.Listen("tcp", service, &config)
	if err != nil {
		log.Fatalf("server: listen: %s", err)
	}

	if err := rpc.Register(new(MyServer)); err != nil {
		log.Fatal("Failed to register RPC method")
	}

	log.Print("server: listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept: %s", err)
			break
		}
		defer conn.Close()
		log.Printf("server: accepted from %s", conn.RemoteAddr())
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	rpc.ServeConn(conn)
	log.Println("server: conn: closed")
}

type MyServer struct{}

func (srv *MyServer) Sum(args *ArgsSum, reply *int) error {
	*reply = args.Item1 + args.Item2
	return nil
}

type ArgsSum struct {
	Item1, Item2 int
}
