package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

const BUF_SIZE int = 4096

var CODE_REQUEST_IMAGE []byte = []byte{'p'}
var CODE_CLOSE_CONNECTION []byte = []byte{'q'}

func main() {
	fmt.Printf("hello world\n")
	host := os.Args[1]
	ConnectToServer(host)
}

func ConnectToServer(host string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("connected to %s\n", conn.RemoteAddr())
	getInput(conn)

}

func getInput(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if text == "q" {
		conn.Close()
		fmt.Printf("closing connection and exiting...\n")
		return
	}
	if text == "p" {
		conn.Write([]byte("p"))
		filename := recieveImage(conn)
		displayPhoto(filename)
	}
}

func displayPhoto(filename string) {
	app := "open"
	arg0 := filename
	cmd := exec.Command(app, arg0)
	stdout, _ := cmd.Output()
	fmt.Println(string(stdout))
}

func recieveImage(conn net.Conn) string {
	t := time.Now()
	filename := t.Format("images/20060102150405")
	filename += ".jpg"

	// dowload file
	photoBytes := make([]byte, 0)
	buf := make([]byte, BUF_SIZE)
	for {
		n, err := conn.Read(buf)
		fmt.Printf("reading %d from conn\n", n)
		photoBytes = append(photoBytes, buf[:n]...)
		if err == io.EOF {
			break
		} else {
			checkErr(err)
		}
	}
	fmt.Printf("finished reading from conn\n")

	// write
	fmt.Printf("downloading image [%s]...\n", filename)
	fd, err := os.Create(filename)
	checkErr(err)
	n, err := fd.Write(photoBytes)
	checkErr(err)
	fmt.Printf("downloaded %d bytes total \n", n)
	return filename
}

func checkErr(e error) {
	if e != nil {
		log.Fatalf(e.Error())
	}
}
