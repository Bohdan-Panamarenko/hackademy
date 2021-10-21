package main

import (
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"tcp-fibonacci/calc_fibonacci.go"
	"tcp-fibonacci/fibonacci_con_types"
	"time"
)

func failResponse(conn net.Conn) {
	resp := &fibonacci_con_types.FibResponse{}

	resp.Num = big.NewInt(-1)
	resp.Send(conn)
}

func main() {

	fmt.Println("Launching server...")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept() // wait for new connection
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		go func() {
			log.Println("New connection")
			for {
				req, err := fibonacci_con_types.GetRequest(conn) // get request
				if err == io.EOF {
					log.Println("Connection broken")
					break
				} else if err != nil {
					log.Println("Getting request err:", err.Error())
					failResponse(conn)
					continue
				}

				resp := &fibonacci_con_types.FibResponse{}

				start := time.Now()                          // time
				resp.Num, err = calc_fibonacci.Calc(req.Num) // calculate fibonacci
				if err != nil {
					log.Println("Calculating err:", err.Error())
					failResponse(conn)
					continue
				}
				resp.Time = time.Since(start) // time

				err = resp.Send(conn)
				if err != nil {
					log.Println("Sending response err:", err.Error())
					failResponse(conn)
				}
			}
		}()
	}
}
