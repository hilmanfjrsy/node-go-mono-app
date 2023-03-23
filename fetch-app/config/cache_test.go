package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestCacheFound(t *testing.T) {
	defer os.RemoveAll("cache")
	name := "test-cache"
	data := `{"id":1}`
	err := SetCache(name, []byte(data), time.Hour)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	b, err := GetCache(name)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	assert.JSONEq(t, data, string(b))
}

func TestCacheNotFound(t *testing.T) {
	name := "test-cache"
	_, err := GetCache(name)
	assert.ErrorContains(t, err, "no such file or directory")
}
