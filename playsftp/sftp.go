package playsftp

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func NewConnection(sftpURL string) *ssh.Client {
	// Parse the URL
	parsedUrl, err := url.Parse(sftpURL)
	if err != nil {
		log.Fatalf("Failed to parse SFTP To Go URL: %s", err)
	}

	// Get user name and pass
	user := parsedUrl.User.Username()
	pass, _ := parsedUrl.User.Password()

	// Parse Host and Port
	host := parsedUrl.Host
	hostKey := getHostKey(host)
	port := "22"

	cfg := ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		// HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		HostKeyCallback: ssh.FixedHostKey(hostKey),
		Timeout:         30 * time.Second,
	}
	addr := fmt.Sprintf("%s:%s", host, port)

	conn, err := ssh.Dial("tcp", addr, &cfg)
	if err != nil {
		log.Fatalf("Failed to connec to host [%s]: %v", addr, err)
	}

	return conn
}

func getHostKey(host string) ssh.PublicKey {
	// parse OpenSSH known_hosts file
	// ssh or use ssh-keyscan to get initial key
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("filePath:", filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				log.Fatalf("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}

	if hostKey == nil {
		log.Fatalf("no hostkey found for %s", host)
	}

	return hostKey
}

func ListDirectory(client *sftp.Client) {
	w := client.Walk("/Users/renugadevi/Documents/Movs")
	for w.Step() {
		if w.Err() != nil {
			continue
		}
		log.Println(w.Path())
	}
}

func Create(client *sftp.Client) {
	f, err := client.Create("/Users/renugadevi/Documents/Movs/hello.txt")
	if err != nil {
		log.Fatal("unable to create: ", err)
	}
	if _, err := f.Write([]byte("Hello world!")); err != nil {
		log.Fatal("unable to write: ", err)
	}
	f.Close()
}

func PrintFileInfo(client *sftp.Client) {
	fi, err := client.Lstat("/Users/renugadevi/Documents/Movs/hello.txt")
	if err != nil {
		log.Fatal("unable to check file: ", err)
	}
	log.Println(fi)
}
