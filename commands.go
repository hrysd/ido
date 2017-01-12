package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	"github.com/bgentry/speakeasy"
	"github.com/codegangsta/cli"
	"github.com/dickeyxxx/netrc"
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

	defer response.Body.Close()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	decorder := json.NewDecoder(response.Body)

	var token Token

	err = decorder.Decode(&token)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	n := getNetrc()
	n.RemoveMachine("idobata.io")
	n.AddMachine("idobata.io", email, token)
	n.Machine("idobata.io").Set("token", token.AccessToken)
	n.Save()
}

func netrcPath() string {
	user, err := user.Current()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return filepath.Join(user.HomeDir, ".netrc")
}

func getNetrc() *netrc.Netrc {
	n, err := netrc.Parse(netrcPath())

	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			return &netrc.Netrc{Path: netrcPath()}
		}

		fmt.Println(err.Error())
		os.Exit(1)
	}

	return n
}
