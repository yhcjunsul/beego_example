package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/astaxie/beego"
)

func HandleHttpRequest(method string, url string, models interface{}, body string) error {
	w := httptest.NewRecorder()

	bodyReader := strings.NewReader(body)

	r, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return err
	}

	beego.BeeApp.Handlers.ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		return fmt.Errorf("Invalid status for getting models, code:%d", w.Code)
	}

	if models == nil {
		return nil
	}

	if err := json.Unmarshal(w.Body.Bytes(), models); err != nil {
		return fmt.Errorf("Failed to unmarshal for getting models, json:%s", string(w.Body.Bytes()))
	}

	return nil
}
