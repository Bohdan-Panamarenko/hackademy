package goroutines

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Queue struct {
	arr []int
}

func NewQueue() *Queue {
	q := &Queue{
		arr: make([]int, 0),
	}

	return q
}

func (q *Queue) Push(item int) {
	q.arr = append(q.arr, item)
}

func (q *Queue) Pop() int {
	item := q.arr[0]
	q.arr = q.arr[1:]
	return item
}

type WorkerPool struct {
	sync.RWMutex
	idQueue *Queue
	size    int
	// workerList map[int]chan float64
}

func NewWorkerPool(poolSize int) *WorkerPool {
	wp := &WorkerPool{
		idQueue: NewQueue(),
		size:    poolSize,
	}

	for i := 1; i <= poolSize; i++ {
		wp.idQueue.Push(i)
	}

	return wp
}

func (wp *WorkerPool) GetId() int {
	wp.Lock()
	defer wp.Unlock()
	return wp.idQueue.Pop()
}

func (wp *WorkerPool) PutId(id int) {
	wp.Lock()
	defer wp.Unlock()
	wp.idQueue.Push(id)
}

func (wp *WorkerPool) CanStart() bool {
	wp.RLock()
	defer wp.RUnlock()
	if len(wp.idQueue.arr) == 0 {
		return false
	}
	return true
}

func (wp *WorkerPool) newWorker(jobs <-chan float64, wg *sync.WaitGroup) {
	id := wp.GetId()
	defer func() {
		wp.PutId(id)
		wg.Done()
	}()

	fmt.Printf("worker:%d spawning\n", id)

	for {
		select {
		case j := <-jobs:
			fmt.Printf("worker:%d sleep:%.1f\n", id, j)
			time.Sleep(time.Duration(int(j * float64(int(time.Second)))))
		default:
			fmt.Printf("worker:%d stopping\n", id)
			return
		}
	}
}

func (wp *WorkerPool) Stats() int {
	return wp.size - len(wp.idQueue.arr)
}

func statsHandler(wp *WorkerPool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprint(wp.Stats())))
	}
}

func (wp *WorkerPool) Serve() *http.Server {
	r := mux.NewRouter()
	r.HandleFunc("/stats", statsHandler(wp)).Methods(http.MethodGet)
	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Println("Server exited with error:", err)
		}
	}()

	return &srv
}

func Run(poolSize int) {
	wp := NewWorkerPool(poolSize)

	errChan := make(chan int, 1)

	requests := make(chan float64, poolSize)
	jobs := make(chan float64, 2*poolSize)

	var wg sync.WaitGroup

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	srv := wp.Serve()
	go readInput(requests, errChan)

	for {
		select {
		case req := <-requests:
			jobs <- req
			if wp.CanStart() {
				wg.Add(1)
				go wp.newWorker(jobs, &wg)
			}
		case <-interrupt:
		case <-errChan:
			wg.Wait()
			srv.Shutdown(context.Background())
			// fmt.Println("Good bye")
			return
			// default:
			// 	go func
		}
	}

}

func readInput(results chan<- float64, errChan chan<- int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString('\n') // read user input
		if err != nil {
			if err == io.EOF {
				errChan <- 1
				return
			}
			log.Println("User input reading err:", err.Error())
			continue
		}

		text = strings.Replace(text, "\n", "", -1) // remove end of line symbol from user input

		float, err := strconv.ParseFloat(text, 32)
		if err != nil {
			log.Println("Float parsing err:", err.Error())
		}

		results <- float
	}
}
