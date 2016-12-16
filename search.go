package goT411

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

const (
	offset = iota
	limit
	cat
	subcat
	lang
	quality
	ep
	season
)

//Torrent represent a torrent
type Torrent struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Seeders   string `json:"seeders"`
	Added     string `json:"added"`
	Size      string `json:"size"`
	Completed string `json:"times_completed"`
}

//Torrents represent a list of Torrent
type Torrents []Torrent

//ResultSearch represent a list of Torrent
type ResultSearch struct {
	Query    string   `json:"query"`
	Total    string   `json:"total"`
	Torrents Torrents `json:"torrents"`
}

//SearchTorrents return a list of torrent that corresponds to query and params
func (c Client) SearchTorrents(query string, params map[int]string) (*ResultSearch, error) {
	var out ResultSearch
	//TODO use params
	resp, err := c.requestURI("http://" + apiRoot + "/torrents/search/" + query + "")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(resp), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

//DownloadFile return the url of the on-server Downloaded file
func (c Client) DownloadFile(id string) (string, error) {
	resp, err := c.requestURI("http://" + apiRoot + "/torrents/download/" + id)
	if err != nil {
		return "", err
	}
	filename := "/tmp/torrent_" + id + ".torrent"
	output, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer output.Close()

	_, err = io.Copy(output, bytes.NewReader([]byte(resp)))
	if err != nil {
		return "", err
	}
	return resp, nil
}
