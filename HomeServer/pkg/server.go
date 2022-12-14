package server

import (
	"fmt"
	"log"
	"net"
	"reflect"
)

var CODE_REQUEST_IMAGE []byte = []byte{'p'}
var CODE_CLOSE_CONNECTION []byte = []byte{'c'}

func RunServer(host string) {
	listener, err := net.Listen("tcp", host)
	checkErr(err)
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		go HandleConn(conn, err)
	}
}

func HandleConn(conn net.Conn, e error) {
	defer conn.Close()
	fmt.Printf("recieved connection from %s\n", conn.LocalAddr())
	checkErr(e)
	code := make([]byte, 1)
	for {
		conn.Read(code)
		fmt.Printf("got code %d, not %d\n", code, []byte{'p'})

		if reflect.DeepEqual(code, CODE_REQUEST_IMAGE) {
			fmt.Printf("writing image\n")
			writeImage(conn)
			conn.Close()
		} else if reflect.DeepEqual(code, CODE_CLOSE_CONNECTION) {
			conn.Close()
			return
		}
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}
