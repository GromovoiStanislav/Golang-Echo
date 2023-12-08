package main

import (
	"net/http"
	"strconv"
	"testing"
	"io/ioutil"

	"github.com/stretchr/testify/assert"
)

func BenchmarkRequestEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, err := http.Get("http://localhost:" + strconv.Itoa(portEcho))
		if err != nil {
			b.FailNow()
		}

		if resp.StatusCode != http.StatusOK {
			print("error")
		}
		err = resp.Body.Close()
		if err != nil {
			b.FailNow()
		}
	}
}


func TestRootRoute(t *testing.T) {
	resp, err := http.Get("http://localhost:" + strconv.Itoa(portEcho))
	if err != nil {
		t.FailNow()
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.FailNow()
	}

	assert.Equal(t, 200, resp.StatusCode)
	assert.JSONEq(t, `{"name":"echo"}`, string(body))
}