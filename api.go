package goT411

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var apiRoot = "api.t411.li"

//Client is used to deal with server t411
type Client struct {
	user   *user
	client *http.Client
	auth   string
}

//GetClient return a new client to deal with t411, password and id are set by default from env T411_ID and T411_PWD
func GetClient() (out *Client) {
	out = new(Client)
	out.user = new(user)
	out.user.ID = os.Getenv("T411_ID")
	out.user.Pwd = os.Getenv("T411_PWD")
	out.client = new(http.Client)
	return
}

func (c Client) requestURI(uri string) (string, error) {
	r, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", err
	}
	r.Header.Add("Authorization", c.auth)
	// str, err := httputil.DumpRequestOut(r, true)
	// if err != nil {
	// 	return "", err
	// }
	// fmt.Printf("%s\n", str)
	resp, err := c.client.Do(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// str, err = httputil.DumpResponse(resp, true)
	// if err != nil {
	// 	return "", err
	// }
	// fmt.Printf("%s\n", str)
	if resp.StatusCode != 200 { //OK
		return "", errors.New("Request sent error code " + strconv.Itoa(resp.StatusCode))
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyString := string(bodyBytes)
	return bodyString, nil
}
