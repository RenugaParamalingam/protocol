package main

import (
	"fmt"
	"log"

	"github.com/RenugaParamalingam/protocol/playsftp"
	"github.com/pkg/sftp"
)

func main() {
	playWithSFTP()
}

func playWithSFTP() {

	// Create a url
	sftpUser := "renugadevi"
	sftpPass := "rendev"
	sftpHost := "192.168.1.5"

	rawurl := fmt.Sprintf("sftp://%v:%v@%v", sftpUser, sftpPass, sftpHost)

	conn := playsftp.NewConnection(rawurl)

	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatal("unable to create new client: ", err)
	}
	defer client.Close()

	playsftp.ListDirectory(client)
}
