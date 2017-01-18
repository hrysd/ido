package idobata

import (
  "bytes"
	"fmt"
	"encoding/json"
	"net/http"
	"net/url"
  "os"
	"os/user"
	"path/filepath"

	"github.com/dickeyxxx/netrc"
)

const ENDPOINT = "https://idobata.io"

type GetTokenParams struct {
	GrantType string `json:"grant_type"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
}

type CreateMessageParams struct {
  RoomId int `json:"room_id"`
  Source string `json:"source"`
}

type Token struct {
	TokenType   string `json:"token_type"`
	CreatedAt   int    `json:"created_at"`
	AccessToken string `json:"access_token"`
}

type Hoge struct {
  Rooms []Room `json:"rooms"`
}

type Room struct {
  Id int `json:"id"`
  Name string `json:"name"`
}

func GetToken(email string, password string) (*Token, error) {
	// TODO: If token is
	params, _ := json.Marshal(GetTokenParams{GrantType: "password", UserName: email, Password: password})

  request, _ := newRequest("POST", "/oauth/token", params)

	client := new(http.Client)
	response, err := client.Do(request)

	var token Token

	err = decodeResponse(response, &token)

	if err != nil {
    // TODO: return error
		fmt.Println(err.Error())
		os.Exit(1)
	}

  return &token , nil
}

func CreateMessage(source string, roomId int) {
	params, _ := json.Marshal(CreateMessageParams{RoomId: roomId, Source: source})

  token, error := getTokenFromNetrc()

  request, _ := newRequest("POST", "/api/messages", params)

	client := new(http.Client)
	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetRooms(organizationSlug string, roomName string) Hoge {
	values := url.Values{}
	values.Add("organization_slug", organizationSlug)
	values.Add("room_name", roomName)

	request, _ := newRequest("GET", "/api/rooms?" + values.Encode(), nil)

	client := new(http.Client)
	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

  var data Hoge

  decodeResponse(response, &data)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return data
}

func newRequest(method string, spath string, params []byte) (*http.Request, error) {
  url := ENDPOINT + spath

  token, err := getTokenFromNetrc()

  if err != nil {
    // TODO: return error
		fmt.Println(err.Error())
		os.Exit(1)
  }

  request, err := http.NewRequest(method, url, bytes.NewBuffer(params))

  if err != nil {
    // TODO: return error
		fmt.Println(err.Error())
		os.Exit(1)
  }

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+ token)

  return request, nil
}

func decodeResponse(response *http.Response, i interface{}) error {
  defer response.Body.Close()

  decorder := json.NewDecoder(response.Body)

  return decorder.Decode(i)
}


func getTokenFromNetrc() (string, error) {
  var token string

  n := getNetrc()
  machine := n.Machine("idobata.io")

  if machine == nil {
    return nil, error("hoge")
  }

  token = machine.Get("password")

  return token, nil
}

// TODO: まとめる
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

func netrcPath() string {
	user, err := user.Current()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return filepath.Join(user.HomeDir, ".netrc")
}
