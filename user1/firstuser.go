package main

import (
	proto "chatapp/myproto"
	
	"context"
	"log"
	"io"
	"net/http"
	"fmt"
    "sync"
	"bufio"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
type thread struct{
	wg sync.WaitGroup
}
func main(){
	http.HandleFunc("/fun",callall)
	http.ListenAndServe(":8081",nil)


}
func callall(w http.ResponseWriter,r *http.Request){
	ClientOne()
	w.WriteHeader(200)
}
var th thread
func ClientOne(){
	con,err:=grpc.Dial("localhost:8080",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Println(err)
		os.Exit(1)
	}
	client:=proto.NewChatClient(con)
	var chat string
	
	strem,srr:=client.SendToUserOne(context.TODO())
	if srr!=nil{
		log.Println(srr)
	}
	th.wg.Add(1)
	go func(){
		
		for{
			next,nrr:=strem.Recv()
			if nrr==io.EOF{
				break
			}
			if nrr!=nil{
				break
			}
			log.Println(next)
		}
		defer th.wg.Done()

	}()
	for{
		
		fmt.Println("---")
		scanner:=bufio.NewScanner(os.Stdin)
		if scanner.Scan(){
			chat=scanner.Text()
		}else{
			fmt.Println("___")
		}

		ft:=strem.Send(&proto.UserOne{Chat:chat})
		if ft!=nil{
			log.Println(ft)
			strem.CloseSend()
			break
		}
	}
	
	th.wg.Wait()
	defer con.Close()
}