package tms

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/RenugaParamalingam/protocol/libftp"
	"github.com/jlaffaye/ftp"
)

type tms struct {
	conn *ftp.ServerConn
	path string
}

type ftpURL struct {
	user     string
	password string
	host     string
	port     string
	path     string
}

const (
	ftpConnTimeout = 30 * time.Second
	defaultFtpPort = ":21"
	anonymous      = "anonymous"
)

// FTP URL is expected in format ftp://user:password@host:port/path
func New(ftpURL string) (tms, error) {
	parsedURL, err := parseFTPURL(ftpURL)
	if err != nil {
		return tms{}, fmt.Errorf("failed to parse URL: %w", err)
	}

	if parsedURL.host == "" {
		return tms{}, fmt.Errorf("invalid host")
	}

	c, err := ftp.Dial(parsedURL.host, ftp.DialWithTimeout(ftpConnTimeout))
	if err != nil {
		return tms{}, fmt.Errorf("failed to connect: %w", err)
	}

	err = c.Login(parsedURL.user, parsedURL.password)
	if err != nil {
		log.Fatal("failed to login: ", err)
	}

	return tms{
		conn: c,
		path: parsedURL.path,
	}, nil
}

func parseFTPURL(rawURL string) (ftpURL, error) {
	url, err := url.Parse(rawURL)
	if err != nil {
		return ftpURL{}, err
	}

	userName, password := getUserPass(url.User)

	host := url.Host
	if tokens := strings.Split(host, ":"); len(tokens) == 1 {
		host += defaultFtpPort
	}

	return ftpURL{
		user:     userName,
		password: password,
		host:     host,
		port:     url.Port(),
		path:     url.Path,
	}, nil
}

func getUserPass(user *url.Userinfo) (username string, password string) {
	if username = user.Username(); username != "" {
		password, _ = user.Password()
	} else {
		username = anonymous
		password = anonymous
	}
	return
}

func (t *tms) CloseConnection() error {
	return t.conn.Quit()
}

func (t *tms) Ingest(params libftp.IngestParams) error {
	switch params.EntityType {
	case libftp.EntityTypeKDM:
		return t.ingestKDM(params)
	}

	return nil
}

func (t *tms) ingestKDM(params libftp.IngestParams) error {
	return t.conn.Stor(t.path+"/"+params.FileName, bytes.NewBuffer(params.Data))
}

func (t *tms) DeleteFile(fileName string) {
	t.conn.Delete("k1.xml")
}
