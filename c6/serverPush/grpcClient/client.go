package main

import (
	"log"

	pb "github.com/Miraddo/go-restful/c6/grpcExample/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

// ReceiveStream listens to the stream contents and use them
func ReceiveStream(client pb.MoneyTransactionClient, request *pb.TransactionRequest) {
	log.Println("Started listening to the server stream!")
	_, err := client.MakeTransaction(context.Background(), request)
	if err != nil {
		log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
	}

	// Listen to the stream of messages
	//for {
		//log.Fatalf(stream.)
		//response, err :=
		//if err == io.EOF {
		//	// If there are no more messages, get out of loop
		//	break
		//}
		//if err != nil {
		//	log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
		//}
		//log.Printf("Status: %v, Operation: %v", response.Status, response.Description)
	//}
}

func main()  {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	client := pb.NewMoneyTransactionClient(conn)

	// Prepare data. Get this from clients like Front-end or Android App
	from := "1234"
	to := "5678"
	amount := float32(1250.75)

	// Contact the server and print out its response.
	ReceiveStream(client, &pb.TransactionRequest{From: from,
		To: to, Amount: amount})
}
