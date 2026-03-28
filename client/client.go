package main

import (
	"context"
	"log"
	"time"

	pb "grpc-student-6609650079/studentpb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetStudent(ctx, &pb.StudentRequest{
		Id: 101,
	})

	if err != nil {
		log.Fatalf("Error calling GetStudent: %v", err)
	}

	log.Printf("Student Info:")
	log.Printf("ID: %d", res.Id)
	log.Printf("Name: %s", res.Name)
	log.Printf("Major: %s", res.Major)
	log.Printf("Email: %s", res.Email)
	log.Printf("Phone: %s", res.Phone)

	log.Printf("\n--- All Students List ---")
	listRes, err := client.ListStudents(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error calling ListStudents: %v", err)
	}

	for _, s := range listRes.Student {
		log.Printf("ID: %d", s.Id)
		log.Printf("Name: %s", s.Name)
		log.Printf("Major: %s", s.Major)
		log.Printf("Email: %s", s.Email)
		log.Printf("Phone: %s", s.Phone)
		log.Printf("-------------------------")
	}
}
