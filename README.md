## GoRPCServerClientOverTLS ##
Go RPC Server Client comunicate over TLS + creation script for certificates

## Installation ##

`go get -v github.com/TeoV/GoRPCServerClientOverTLS`

## HowToUse ##

+ Run the script : `./script.sh`.
It will create the following files : 
	+ Certificate authority files : 
		+ ca.key
		+ ca.crt
	+ Server certificate :
		+ server.key
		+ server.csr
		+ server.crt
		+ server1.crt which contains server.crt and ca.crt 
	+ Client cerificate :
		+ client.key
		+ client.csr
		+ client.crt
		+ client1.crt which contains client.crt and ca.crt

+ Run `go run server.go` to start the server. 

+ Run `go run client.go` to connect the client to server. 