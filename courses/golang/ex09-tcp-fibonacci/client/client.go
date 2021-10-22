package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"tcp-fibonacci/fibonacci_con_types"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Println(err.Error())
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString('\n') // read user input
		if err != nil {
			log.Println("User input reading err:", err.Error())
			continue
		}
		text = strings.Replace(text, "\n", "", -1) // remove end of line symbol from user input

		num, err := strconv.ParseInt(text, 10, 64) // parse int64
		if err != nil {
			log.Println("User input parsing err", err.Error())
			continue
		}

		req := fibonacci_con_types.FibRequest{Num: num} // make FibRequest

		err = req.Send(conn)
		if err != nil {
			log.Println("Sending request err:", err.Error())
			continue
		}

		resp, err := fibonacci_con_types.GetResponse(conn)
		if err != nil {
			log.Println("Getting response err:", err.Error())
		}

		fmt.Println(*resp)
	}
}
