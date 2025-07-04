package cloud

type CloudDB struct {
	url string
}

func NewCloudDb(url string) *CloudDB {
	return &CloudDB{
		url: url,
	}
}

func (db *CloudDB) Read() ([]byte, error) {
	return []byte{}, nil
}

func (db *CloudDB) Write(content []byte) {}
