package main

import (
	"context"
	"log"
	"time"

	pb "test-module/grps/01/test-module/gen/calculator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure" // Добавь этот импорт
)

func main() {
	// Подключаемся к серверу
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Создаем клиент
	client := pb.NewCalcClient(conn)

	// Подготавливаем запрос
	req := &pb.Request{
		A: 10,
		B: 20,
	}

	// Устанавливаем таймаут (на случай если сервер не отвечает)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Вызываем метод Add
	resp, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}

	log.Printf("Result: %d + %d = %d", req.A, req.B, resp.Result)
}
