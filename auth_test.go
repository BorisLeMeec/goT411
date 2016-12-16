package goT411_test

import (
	"testing"

	"github.com/BorisLeMeec/goT411"
)

func TestAuth(t *testing.T) {
	client := goT411.GetClient()

	if err := client.Auth(); err != nil {
		t.Error(err)
	}
	client.SetPasswd("nil")
	if err := client.Auth(); err == nil {
		t.Error(err)
	}
}
