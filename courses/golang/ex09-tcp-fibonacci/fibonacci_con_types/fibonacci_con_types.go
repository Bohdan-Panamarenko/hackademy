package fibonacci_con_types

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"strings"
	"time"
)

type FibRequest struct {
	Num int64 `json:"number"`
}

type FibResponse struct {
	Time time.Duration `json:"time"`
	Num  *big.Int      `json:"number"`
}

func (r FibResponse) String() string {
	return fmt.Sprint(r.Time.String() + " " + r.Num.String())
}

type Sender interface {
	Send()
}

func GetRequest(conn net.Conn) (*FibRequest, error) {
	reqMarsh, err := bufio.NewReader(conn).ReadString('\n') // read request
	if err != nil {
		return nil, err
	}

	reqMarsh = strings.TrimSuffix(reqMarsh, "\n") // remove end of line

	req := &FibRequest{}
	err = json.Unmarshal([]byte(reqMarsh), req) // unmarshal request
	if err != nil {
		return nil, err
	}

	return req, nil
}

func GetResponse(conn net.Conn) (*FibResponse, error) {
	respMarsh, err := bufio.NewReader(conn).ReadString('\n') // read response
	if err != nil {
		return nil, err
	}

	respMarsh = strings.TrimSuffix(respMarsh, "\n") // remove end of line

	resp := &FibResponse{}
	err = json.Unmarshal([]byte(respMarsh), resp) // unmarshal response
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r FibRequest) Send(conn net.Conn) error {
	reqMarsh, err := json.Marshal(r) // marshal request
	if err != nil {
		return err
	}

	reqMarsh = append(reqMarsh, byte('\n'))

	_, err = conn.Write(reqMarsh) // send response
	if err != nil {
		return err
	}

	return nil
}

func (r FibResponse) Send(conn net.Conn) error {
	respMarsh, err := json.Marshal(r) // marshal request
	if err != nil {
		return err
	}

	respMarsh = append(respMarsh, byte('\n'))

	_, err = conn.Write(respMarsh) // send response
	if err != nil {
		return err
	}

	return nil
}
