package playftp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
)

func NewConnection(add, user, password string) *ftp.ServerConn {
	c, err := ftp.Dial(add, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(user, password)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

func ReadFile(conn *ftp.ServerConn) {
	fmt.Println("reading file")
	r, err := conn.Retr("test.txt")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal("Failed to read: ", err)
	}
	println(string(buf))
}

func StoreFile(conn *ftp.ServerConn) {
	conn.MakeDir("/fromMac")
	data := bytes.NewBufferString("Hello FTP")
	err := conn.Stor("/fromMac/mac.txt", data)
	if err != nil {
		panic(err)
	}
	fmt.Println("file stored")
}

func DeleteFile(conn *ftp.ServerConn) {
	conn.Delete("test.txt")
}
