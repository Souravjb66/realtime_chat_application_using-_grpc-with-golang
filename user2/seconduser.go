package main

import (
	proto "chatapp/myproto"
	"context"
	"fmt"
	"log"
	"net/http"
    "io"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
func main(){
	router:=http.NewServeMux()

    router.HandleFunc("/secuser",callutwo)
	http.ListenAndServe(":8082",router)
	
}
func callutwo(w http.ResponseWriter,r *http.Request){
	result:=Clienttwo()
	if result==400{
		fmt.Fprintf(w,"second user disconnect")
	}
	w.WriteHeader(result)

}
func Clienttwo()int{
	listen,err:=grpc.Dial("localhost:8080",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
		log.Panic("user 2 server problem")
	}
	defer listen.Close()
	client:=proto.NewChatClient(listen)
	var chat string

	strem,rr:=client.SendToUserTwo(context.TODO())
	if rr!=nil{
		log.Println(rr)
		return http.StatusBadRequest

	}
	for{
		fmt.Println("bool true or false")
		var yesno bool
		fmt.Scan(&yesno)
		if !yesno{
			fmt.Println("user left")
			strem.CloseSend()
			break
		}
		fmt.Println("-----")
		fmt.Scan(&chat)
		se:=strem.Send(&proto.UserTwo{Chat: chat,Id: 2})
		if se!=nil{
			log.Println(se)
			break
		}

	}
	for{
		rd,re:=strem.Recv()
		if re==io.EOF{
			break
		}
		fmt.Println(rd.Reply)
	}
	
	return 200

}