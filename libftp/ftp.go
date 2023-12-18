package libftp

type EntityType string

const (
	EntityTypeKDM EntityType = "kdm"
	EntityTypeDCP EntityType = "dcp"
)

type IngestParams struct {
	EntityType EntityType
	// URL will be in the format ftp://username:password@hostname/path
	URL      string
	Data     []byte
	FileName string
}

type FTP interface {
	CloseConnection() error
	Ingest(params IngestParams) error
}
