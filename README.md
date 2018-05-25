### GoRPCServerClientOverTLS ###
Go RPC Server Client comunicate over TLS + creation script for certificates

### HowToInstall ###

Simply run : 
`go get -v github.com/TeoV/GoRPCServerClientOverTLS`

### HowToUse ###

+ Run the script : `./script.sh`
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

+ Start in a terminat the server with `go run server.go`

+ Start in another terminal the client with :  `go run client.go`