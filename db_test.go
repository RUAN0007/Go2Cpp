package db_test

import (
	"db"
	"testing"
)

func TestGet(t *testing.T) {
	c := db.NewDb()
	v := c.Get("gokey")
	if v != "gval" {
		t.Error("unexpected value: ", v)
	}
}

func TestPut(t *testing.T) {
  c := db.NewDb()
  v := c.Put("gokey", "goval")
  if !v {
    t.Error("unexpected put: ", v)
  }
}
