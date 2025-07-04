package account

import (
	"app-4/output"
	"encoding/json"
	"strings"

	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDB struct {
	Vault
	db Db
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewVault(db Db) *VaultWithDB {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	return &VaultWithDB{
		Vault: vault,
		db:    db,
	}
}

func (vault *VaultWithDB) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *VaultWithDB) FindAccount(url string, checker func(Account, string) bool) []Account {
	var accs []Account
	for _, v := range vault.Accounts {
		isMatched := checker(v, url)
		if isMatched {
			accs = append(accs, v)
		}
	}
	return accs
}

func (vault *VaultWithDB) DeleteAccount(url string) bool {
	var accs []Account
	isDeleted := false
	for _, v := range vault.Accounts {
		isMatched := strings.Contains(v.Url, url)
		if !isMatched {
			accs = append(accs, v)
			continue
		}
		isDeleted = true
	}

	vault.Accounts = accs
	vault.save()
	return isDeleted
}

func (vault *VaultWithDB) save() {
	vault.UpdatedAt = time.Now()

	data, err := vault.Vault.ToBytes()
	if err != nil {
		output.PrintError(err.Error())
	}
	vault.db.Write(data)
}
