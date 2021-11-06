.PHONY : cli factorial clean

grpcfactorial: server client
	(./serverbin)

server: factorial
	go build -o serverbin calculator/calculator_server/server.go 

client:	cli 
	go build -o grpcfactorial calculator/calculator_client/client.go

factorial: 
	go build factorial/factorial.go

cli: 
	go build cli/inputs.go

clean:
	rm ./serverbin
	rm ./grpcfactorial 	