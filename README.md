# Using the program
1. Open the terminal and navigate to the MEANT4-Task directory

2. Run `make` in the terminal. You should see the following out put
```
go build factorial/factorial.go
go build -o serverbin calculator/calculator_server/server.go 
go build cli/inputs.go
go build -o grpcfactorial calculator/calculator_client/client.go
(./serverbin)
factorial server is up.
2021/11/06 20:14:15 server listening at 127.0.0.1:50051
```
3. Optionally, you can build the server and client binaries seperately by running the following commands.

`make server`

`make client`

You can then run the server (`./serverbin`) followed by the client (`./grpcfactorial` In a different terminal). The order matters.
`./serverbin`
`./grpcfactorial`

4. This will create two binaries named 'serverbin' and 'grpcfactorial' in the project-level directory and 
run the server on address 127.0.0.1:50051.

5. Open another terminal and run the `grpcfactorial` binary with the `--inputs` flag set to `integers` followed by a space-seperated
list of the integers whose factorials you want to calculate as shown below.
```
./grpcfactorial --inputs integers 37 58 10000000000 5 6 7 8 1
```
#Tests
You can navigate into the directory  `./factorial/` at the terminal and type `go test -v` to run tests.
Optionally you can run `make test` at the project root directory to run tests.

