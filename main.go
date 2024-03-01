package main
import(
	"log"
	"net"
	"fmt"
	"io"
	"os"
	"bufio"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	proto "chatapp/myproto"
	"sync"
)
type server struct{
	proto.UnimplementedChatServer
}
type thread struct{
	wg sync.WaitGroup
}
var th thread
func main(){
	listen,err:=net.Listen("tcp",":8080")
	if err!=nil{
		log.Fatal(err)
	}
    th.wg=sync.WaitGroup{}
	srv:=grpc.NewServer()
	proto.RegisterChatServer(srv,&server{})
	reflection.Register(srv)
	if ru:=srv.Serve(listen);ru!=nil{
		log.Panic("server fault",ru)
	}

}
func(s *server)SendToUserOne(strem proto.Chat_SendToUserOneServer)error{
	th.wg.Add(1)
	go func(){
		for{
			rv,mn:=strem.Recv()
			if mn==io.EOF{
				
				break
			}
			if mn!=nil{
				log.Println("waiting--")
				break
			}
			log.Println(rv)
			
			
		}
		defer th.wg.Done()

	}()
	
	for{
		
        var Chatone string
		fmt.Println("---")
		scanner:=bufio.NewScanner(os.Stdin)
		if scanner.Scan(){
			Chatone=scanner.Text()
		}else{
			fmt.Println("____")
		}
		se:=strem.Send(&proto.ServerResOne{Reply: Chatone})
		if se!=nil{
			log.Println(se)
			break
		
		}


	}
	
	th.wg.Wait()
	return nil
	
}







// func(s *server)SendToUserTwo(strem proto.Chat_SendToUserTwoServer)error{
// 	for{
// 		se:=strem.Send(&proto.ServerResTwo{Reply: chatdata.Chatone})
// 		if se!=nil{
// 			log.Println(se)
// 			break
// 		}

// 	}
// 	for{
// 		rc,re:=strem.Recv()
// 		if re!=nil{
// 			log.Println(re)
// 			break

// 		}
// 		log.Println(rc)
// 		chatdata.Chattwo=rc.Chat
		
// 	}	


	
	
// 	return nil
	

// }