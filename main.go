package main

import (
	"app-4/account"
	"app-4/files"
	"app-4/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDB){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func promptData[T any](prompt []T) string {
	for i, v := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", v)
		} else {
			fmt.Println(v)
		}
	}
	var res string
	fmt.Scan(&res)
	return res
}

func createAccount(vault *account.VaultWithDB) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите URL"})
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError(err.Error())
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccountByUrl(vault *account.VaultWithDB) {
	url := promptData([]string{"Введите URL для поиска"})
	accs := vault.FindAccount(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})

	outputResult(&accs)
}

func findAccountByLogin(vault *account.VaultWithDB) {
	login := promptData([]string{"Введите логин для поиска"})
	accs := vault.FindAccount(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResult(&accs)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		output.PrintError("Аккаунт не найден \n")
	}

	for _, v := range *accounts {
		v.Output()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
	url := promptData([]string{"Введите URL для поиска"})
	isDeleted := vault.DeleteAccount(url)

	if isDeleted {
		color.Green("Аккаунт успешно удален!")
	} else {
		output.PrintError("Не удалось удалить аккаунт")
	}
}

func main() {
	fmt.Println("___Менеджер паролей___")
	vault := account.NewVault(files.NewJsonDB("data.json"))
	// vault := account.NewVault(cloud.NewCloudDb("https://a.ru"))
Menu:
	for {
		variant := promptData([]string{"1. Создать аккаунт", "2. Найти аккаунт по URL", "3. Найти аккаунт по логину", "4. Удалить аккаунт", "5. Выход", "Выберите вариант"})
		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			findAccountByUrl(vault)
		case "3":
			findAccountByLogin(vault)
		case "4":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}
