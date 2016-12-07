package goT411

import "os"

type user struct {
	id, pwd string
}

//Client is used to deal with server t411
type Client struct {
	user *user
	auth string
}

//Client return a new client to deal with t411, password and id are set by default from env T411_ID and T411_PWD
func getClient() (out *Client) {
	out = new(Client)
	out.user = new(user)
	out.user.id = os.Getenv("T411_ID")
	out.user.pwd = os.Getenv("T411_PWD")
	return
}

//Auth return nil if connection with t411 is ok, or error if not
func (c Client) Auth() error {

	return nil
}

//SetPasswd set a new password for the Client
func (c Client) SetPasswd(pwd string) error {
	//TODO little verifications for pwd
	c.user.pwd = pwd
	return nil
}

//SetLogin set a new login for the Client
func (c Client) SetLogin(login string) error {
	//TODO little verifications for login
	c.user.id = login
	return nil
}
