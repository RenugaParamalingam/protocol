package main

import (
	"fmt"
	"log"

	"github.com/RenugaParamalingam/protocol/playftp"
	"github.com/RenugaParamalingam/protocol/playsftp"
	"github.com/pkg/sftp"
)

func main() {
	// playWithSFTP()
	playWithFTP()
}

func playWithSFTP() {

	// Create a url
	sftpUser := "renugadevi"
	sftpPass := "rendev"
	sftpHost := "192.168.1.5"

	rawurl := fmt.Sprintf("sftp://%v:%v@%v", sftpUser, sftpPass, sftpHost)

	conn := playsftp.NewConnection(rawurl)
	defer conn.Close()

	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatal("unable to create new client: ", err)
	}
	defer client.Close()

	playsftp.ListDirectory(client)
}

func playWithFTP() {
	conn := playftp.NewConnection("192.168.1.8:21", "renugaftp", "password")
	defer conn.Quit()

	// playftp.ReadFile(conn)
	playftp.StoreFile(conn)

	// playftp.DeleteFile(conn)
}
