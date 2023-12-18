package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/RenugaParamalingam/protocol/libftp"
	"github.com/RenugaParamalingam/protocol/libftp/tms"
	"github.com/RenugaParamalingam/protocol/playftp"
	"github.com/RenugaParamalingam/protocol/playsftp"
	"github.com/pkg/sftp"
)

func main() {
	// playWithSFTP()
	// playWithFTP()

	t, err := tms.New("ftp://renugaftp:password@192.168.1.8/ingestTest")
	if err != nil {
		log.Fatal("unable to connect: ", err)
	}

	b, err := json.Marshal("hello renuga")
	if err != nil {
		log.Fatal("unable to marshal: ", err)
	}

	err = t.Ingest(libftp.IngestParams{EntityType: libftp.EntityTypeKDM, Data: b, FileName: "k1.xml"})
	if err != nil {
		log.Fatal("unable to ingest: ", err)
	}

	if err := t.CloseConnection(); err != nil {
		log.Fatal("unable to close: ", err)
	}
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
