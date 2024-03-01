#realtime chat app using grpc
##in this project grpc use with golang
##we use bidirectional streaming for this
**in main server we use net to create the tcp connection
**then use grpc to create grpc webserver
**and i create a proto file in myproto package
** then use this use this command in terminal in myproto path "go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28"
**    "go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2"
** "protoc --go-grpc_out=. --go_out=. *.proto"
## in client
**connect grpc client with grpc server
**send data using function created by proto



































@https://github.com/Souravjb66
