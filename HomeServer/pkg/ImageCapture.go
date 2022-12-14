package server

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
)

const BUF_SIZE int = 4096

// takes a webcam image and writes it to the conn
func writeImage(conn net.Conn) {
	fmt.Printf("writing image to conn...\n")
	// create image name
	filename := "webcam_photo.jpg"

	// take photo
	app := "imagesnap"
	arg0 := "-w"
	arg1 := "3.00"
	arg2 := filename
	cmd := exec.Command(app, arg0, arg1, arg2)
	stdout, _ := cmd.Output()
	fmt.Println(string(stdout))

	// open photo and convert to []byte
	fd, err := os.Open(filename)
	checkErr(err)
	buf := make([]byte, BUF_SIZE)
	for {
		n, err := fd.Read(buf)
		fmt.Printf("writing %d to conn\n", n)
		conn.Write(buf[:n])
		if err == io.EOF {
			break
		} else {
			checkErr(err)
		}
	}

}

// takes a photograph in the .jpg form then packages it into []byte for transport
func saveImage() []byte {
	// create image name
	filename := "webcam_photo.jpg"

	// take photo
	app := "imagesnap"
	arg0 := "-w"
	arg1 := "3.00"
	arg2 := filename
	cmd := exec.Command(app, arg0, arg1, arg2)
	stdout, _ := cmd.Output()
	fmt.Println(string(stdout))

	// open photo and convert to []byte
	fd, err := os.Open(filename)
	checkErr(err)
	photoBytes := make([]byte, 0)
	buf := make([]byte, BUF_SIZE)
	for {
		n, err := fd.Read(buf)
		photoBytes = append(photoBytes, buf[:n]...)
		if err == io.EOF {
			break
		} else {
			checkErr(err)
		}
	}

	// return
	return photoBytes
}
