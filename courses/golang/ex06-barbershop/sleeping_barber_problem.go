package main

import (
	"fmt"
	"time"
)

func Barber(c, closeBarber chan int, clientServing *int) {
	var client int
	sleep := 1
	for {
		select {
		case client = <-c:
			*clientServing = client
			if sleep == 1 {
				sleep = 0
				fmt.Println("--Barber woke up")
			}
			fmt.Printf("Barber serving client %v\n", client)
			time.Sleep(60 * time.Millisecond)
			fmt.Printf("Barber served client %v\n", client)
			*clientServing = 0
		case <-closeBarber:
			return
		default:
			if sleep != 1 {
				fmt.Println("--Barber is asleep")
				sleep = 1
			}
		}
	}
}

func Queue(clients, closeQueue chan int) {
	const numOfSeats = 4
	var Seats []int
	var ClientServing int = 0

	Push := func(client int) bool {
		message := fmt.Sprintf("<->Client %v", client)

		if (len(Seats) < numOfSeats) {
			Seats = append(Seats, client)
			fmt.Println(message + " is in queue")
			return true
		}
		fmt.Println(message + " didn't find free seat and went out")
		return false
	}

	Pop := func() int {
		if (len(Seats) > 0) {
			client := Seats[0]
			Seats = Seats[1:]
			return client
		}
		return 0
	}

	barberQueue := make(chan int)
	barberClose := make(chan int)
	go Barber(barberQueue, barberClose, &ClientServing)

	for {
		select {
		case c := <-clients:

			if (ClientServing == 0) {
				if len(Seats) == 0 {
					barberQueue <- c 
				} else {
					Push(c)
					barberQueue <- Pop()
				}
			} else {
				Push(c)
			}
		case <-closeQueue:
			barberClose <- 1
			return
		default:
			if (ClientServing == 0 && len(Seats) > 0) {
				barberQueue <- Pop()
			}
		}
	}
}

func main() {
	clients := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	queue := make(chan int, 2)
	close := make(chan int)
	
	go Queue(queue, close)
	time.Sleep(50 * time.Millisecond)

	length := len(clients)
	half := length / 2

	for i := 0; i < half; i++ {
		queue <- clients[i]
	}

	time.Sleep(1000 * time.Millisecond)

	for i := half; i < length; i++ {
		queue <- clients[i]
		time.Sleep(50 * time.Millisecond)
	}

	// for _, value := range clients {
	// 	goToQueue(value, queue)
	// 	time.Sleep(10 * time.Millisecond)
	// 	// queue <- value
	// }
	close <- 1
}