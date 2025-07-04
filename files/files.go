package files

import (
	"fmt"
	"os"
)

type JsonDB struct {
	filename string
}

func NewJsonDB(name string) *JsonDB {
	return &JsonDB{
		filename: name,
	}
}

func (db *JsonDB) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func (db *JsonDB) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Запись файла успешна!")
}
