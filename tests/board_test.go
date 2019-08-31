package test

import (
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/yhcjunsul/beego_example/models"
	_ "github.com/yhcjunsul/beego_example/routers"
	"github.com/yhcjunsul/beego_example/utils"

	"github.com/astaxie/beego"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	utils.InitTestSql()
	beego.TestBeegoInit(apppath)
}

func TestBoard(t *testing.T) {
	const BoardName = "test"

	var categories []models.BoardCategory
	err := HandleHttpRequest(http.MethodGet, "/v1/board_categories", &categories, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting all categories, err:%s", err.Error())
	}

	var board models.Board
	body := `{"name":"` + BoardName + `"}`
	boardCreateUrl := fmt.Sprintf("/v1/board_category/%d/board", categories[0].Id)
	err = HandleHttpRequest(http.MethodPost, boardCreateUrl, &board, body)
	if err != nil {
		t.Fatalf("Failed to handle request for creating board, err:%s", err.Error())
	}

	var boards []models.Board
	boardsByCategoryUrl := fmt.Sprintf("/v1/board_category/%d/boards", categories[0].Id)
	err = HandleHttpRequest(http.MethodGet, boardsByCategoryUrl, &boards, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting boards by category, err:%s", err.Error())
	}

	foundIndex := -1
	for i, b := range boards {
		if b.Id == board.Id {
			foundIndex = i
		}
	}

	if foundIndex < 0 || boards[foundIndex].Name != BoardName {
		t.Fatalf("Invalid found index, found index:%d, board id:%d", foundIndex, board.Id)
	}

	boards = nil
	err = HandleHttpRequest(http.MethodGet, "/v1/boards", &boards, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting all boards, err:%s", err.Error())
	}

	foundIndex = -1
	for i, b := range boards {
		if b.Id == board.Id {
			foundIndex = i
		}
	}

	if foundIndex < 0 || boards[foundIndex].Name != BoardName {
		t.Fatalf("invalid found index, found index:%d, board id:%d", foundIndex, board.Id)
	}

	deleteUrl := fmt.Sprintf("/v1/board/%d", board.Id)
	err = HandleHttpRequest(http.MethodDelete, deleteUrl, nil, "")
	if err != nil {
		t.Fatalf("Failed to handle request for deleting board, err:%s", err.Error())
	}

	boards = nil
	err = HandleHttpRequest(http.MethodGet, boardsByCategoryUrl, &boards, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting boards by category, err:%s", err.Error())
	}

	foundIndex = -1
	for i, b := range boards {
		if b.Id == board.Id {
			foundIndex = i
		}
	}

	if foundIndex < 0 || boards[foundIndex].IsDeleted != true {
		t.Fatalf("Invalid found index, found index:%d, board id:%d", foundIndex, board.Id)
	}

	boards = nil
	err = HandleHttpRequest(http.MethodGet, "/v1/boards", &boards, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting all boards, err:%s", err.Error())
	}

	foundIndex = -1
	for i, b := range boards {
		if b.Id == board.Id {
			foundIndex = i
		}
	}

	if foundIndex < 0 || boards[foundIndex].IsDeleted != true {
		t.Fatalf("Invalid found index, found index:%d, board id:%d", foundIndex, board.Id)
	}
}
