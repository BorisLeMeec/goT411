package goT411

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type user struct {
	ID  string `json:"username"`
	Pwd string `json:"password"`
}

type response struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
	UID   string `json:"uid"`
	Token string `json:"token"`
}

//Auth return nil if connection with t411 is ok, or error if not
func (c *Client) Auth() error {
	var resp response

	r, err := http.Post("http://"+apiRoot+"/auth", "application/x-www-form-urlencoded", bytes.NewBuffer([]byte("username="+c.user.ID+"&password="+c.user.Pwd)))
	if err != nil {
		return err
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&resp)
	if resp.Code != 0 {
		return errors.New("T411 auth response is " + strconv.Itoa(resp.Code))
	}
	c.auth = resp.Token
	return nil
}

//SetPasswd set a new password for the Client
func (c Client) SetPasswd(pwd string) error {
	//TODO little verifications for pwd
	c.user.Pwd = pwd
	return nil
}

//SetLogin set a new login for the Client
func (c Client) SetLogin(login string) error {
	//TODO little verifications for login
	c.user.ID = login
	return nil
}
