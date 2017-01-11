package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bgentry/speakeasy"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	initCommand,
	listCommand,
}

var initCommand = cli.Command{
	Name:        "init",
	Usage:       "Initialize with your credential",
	Description: "Not yet",
	Action:      initialize,
}

func initialize(c *cli.Context) {
	email := prompt("Email: ")
	password := passwordPrompt("Password (hidden): ")

	getToken(email, password)
}

var listCommand = cli.Command{
	Name:        "list",
	Usage:       "Show all stored hook",
	Description: "Show all stored hook in `~/.ido`",
	Action:      list,
}

func list(c *cli.Context) {
	hooks := AllHooks()

	var lines bytes.Buffer

	for i, h := range hooks {
		lines.Write([]byte(fmt.Sprintf("%v: %v", h.Name, h.Token)))

		if i < (len(hooks) - 1) {
			lines.Write([]byte("\n"))
		}
	}

	fmt.Println(lines.String())
}

func prompt(label string) string {
	var input string

	fmt.Println(label)
	fmt.Scanln(&input)

	return input
}

func passwordPrompt(label string) string {
	password, err := speakeasy.Ask(label)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(len(password))

	return password
}

type Params struct {
	GrantType string `json:"grant_type"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
}

type Token struct {
	TokenType   string `json:"token_type"`
	CreatedAt   int    `json:"created_at"`
	AccessToken string `json:"access_token"`
}

func getToken(email string, password string) {
	// TODO:
	params, _ := json.Marshal(Params{GrantType: "password", UserName: email, Password: password})

	fmt.Println(string(params))

	response, err := http.Post("https://idobata.io/oauth/token", "application/json", bytes.NewBuffer(params))

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	token := Token{}
	err = json.Unmarshal(body, &token)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("%#v", token)
}
