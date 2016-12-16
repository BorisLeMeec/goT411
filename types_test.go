package goT411_test

import (
	"strconv"
	"testing"

	"github.com/BorisLeMeec/goT411"
)

func TestGetCat(t *testing.T) {
	client := goT411.GetClient()
	if err := client.Auth(); err != nil {
		t.Error(err)
	}
	cat, err := client.GetCat()
	if err != nil {
		t.Error(err)
	}
	if len(cat) != 9 {
		t.Error("Not the good number of cat, expected 3, get " + strconv.Itoa(len(cat)))
	}
}
