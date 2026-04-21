package main

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	pb "test-module/grps/01/test-module/gen/calculator" // Импортируем из папки calculator
)

type server struct {
	pb.UnimplementedCalcServer
}

func (s *server) Add(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Printf("Adding %d + %d", req.A, req.B)
	return &pb.Response{Result: req.A + req.B}, nil
}

var mtx sync.Mutex
var call int
var mt sync.Mutex
var callCache int

func main() {
	t.TestAddPositive()
	t.Test()
	// fmt.Println(runtime.NumCPU())
	// fmt.Println(runtime.GOMAXPROCS(2))
	// fmt.Println(runtime.GOMAXPROCS(2))

}

var ErrNegativeNumber = errors.New("negative number is not allowed")

func AddPositive(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, ErrNegativeNumber
	}
	return a + b, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
func writeToCache() {

	time.Sleep(100 * time.Millisecond)
	mt.Lock()
	callCache++
	mt.Unlock()
}

func getSomithing() {
	go func() {
		withSemaphore(func() {
			writeToCache()
		})
	}()
	mtx.Lock()
	call++
	mtx.Unlock()
}

var maxgrt = 10

var sem = make(chan struct{}, maxgrt)

func withSemaphore(f func()) {

	select {
	case sem <- struct{}{}:
	default:
		return
	}

	go func() {
		f()
		<-sem
	}()
}
