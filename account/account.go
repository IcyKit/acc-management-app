package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) Output() {
	fmt.Println("Аккаунт найден:", acc.Login, acc.Password)
}

func (acc *Account) generatePassword(l int) {
	res := make([]rune, l)
	for i := range res {
		res[i] = letterRunes[rand.IntN((len(letterRunes)))]
	}
	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Неверный URL")
	}

	if len(login) == 0 {
		return nil, errors.New("Неверный формат логина")
	}

	newAcc := &Account{
		Url:       urlString,
		Password:  password,
		Login:     login,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if len(password) == 0 {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
}

var letterRunes = []rune("abcdefghijklmonpqrsuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")
