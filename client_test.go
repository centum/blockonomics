package blockonomics

import (
	"crypto/md5" //nolint:gosec
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type BlockonomicsHandlers struct {
	token    string
	contents map[int]string
	server   *httptest.Server
}

func (s *BlockonomicsHandlers) SetServer(server *httptest.Server) {
	s.server = server
}

func (s *BlockonomicsHandlers) NewAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	authHeader := r.Header.Get("Authorization")
	b := strings.Split(authHeader, " ")
	if len(b) != 2 || b[0] != "Bearer" || b[1] != s.token {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	data := NewAddress{
		Address: fmt.Sprintf("%x", md5.Sum(genPass(16))), //nolint:gosec
	}
	if r.URL.Query().Get("reset") != "" {
		data.Reset = 1
	}
	data.Account = r.URL.Query().Get("match_account")

	body, err := json.Marshal(&data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

func (s *BlockonomicsHandlers) Invoice(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	var data struct {
		Content string `json:"content"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	key := rand.Intn(100000)
	s.contents[key] = data.Content

	body, err = json.Marshal(&struct {
		Number int `json:"number"`
	}{
		Number: key,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

func NewTestBlockonomicsHandler() *BlockonomicsHandlers {
	b := &BlockonomicsHandlers{token: "token123", contents: make(map[int]string)}
	m := http.NewServeMux()
	m.HandleFunc("/api/new_address", b.NewAddress)
	m.HandleFunc("/api/invoice", b.Invoice)
	server := httptest.NewServer(m)
	b.SetServer(server)
	return b
}

func TestClient(t *testing.T) {
	b := NewTestBlockonomicsHandler()

	t.Run("make unauthorized request", func(t *testing.T) {
		c := NewClient("token", WithTimeout(time.Duration(30)*time.Second))
		c.APIBase = b.server.URL

		_, err := c.NewAddress("", false)
		if !errors.Is(err, ErrUnauthorised) {
			t.Error("must be error: ErrUnauthorised")
		}
	})

}
