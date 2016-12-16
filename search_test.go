package goT411_test

import (
	"testing"

	"github.com/BorisLeMeec/goT411"
)

func TestSearchTorrents(t *testing.T) {
	client := goT411.GetClient()
	var params map[int]string

	if err := client.Auth(); err != nil {
		t.Error(err)
		return
	}

	res, err := client.SearchTorrents("Expendables", params)
	if err != nil {
		t.Error(err)
		return
	}
	if len(res.Torrents) != 10 {
		t.Error("Not enough torrent, semms weird\n")
	}
}

func TestDownloadFile(t *testing.T) {
	client := goT411.GetClient()

	if err := client.Auth(); err != nil {
		t.Error(err)
		return
	}
	_, err := client.DownloadFile("50096")
	if err != nil {
		t.Error(err)
	}
}
