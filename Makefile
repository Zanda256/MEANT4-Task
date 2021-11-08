.PHONY : cli factorial clean test

grpcfactorial: server client
	./serverbin

server: factorial
	go build -o serverbin calculator/calculator_server/server.go 

client:	cli 
	go build -o grpcfactorial calculator/calculator_client/client.go

factorial: 
	go build factorial/factorial.go

cli: 
	go build cli/inputs.go

test:
	go test ./factorial/ -v

clean:
	rm ./serverbin
	rm ./grpcfactorial 	