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

func TestBoardCategory(t *testing.T) {
	const CategoryName = "test"

	var category models.BoardCategory
	body := `{"name":"` + CategoryName + `"}`
	err := HandleHttpRequest(http.MethodPost, "/v1/board_category", &category, body)
	if err != nil {
		t.Fatalf("Failed to handle request for creating category, err:%s", err.Error())
	}

	var categories []models.BoardCategory
	err = HandleHttpRequest(http.MethodGet, "/v1/board_categories", &categories, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting all categories, err:%s", err.Error())
	}

	foundIndex := -1
	for i, c := range categories {
		if c.Id == category.Id {
			foundIndex = i
		}
	}

	if categories[foundIndex].Name != CategoryName {
		t.Fatalf("Invalid index, found index:%d, found id:%d, category id:%d",
			foundIndex, categories[foundIndex].Id, category.Id)
	}

	deleteUrl := fmt.Sprintf("/v1/board_category/%d", category.Id)
	err = HandleHttpRequest(http.MethodDelete, deleteUrl, nil, "")
	if err != nil {
		t.Fatalf("Failed to handle request for deleting category, err:%s", err.Error())
	}

	categories = nil
	err = HandleHttpRequest(http.MethodGet, "/v1/board_categories", &categories, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting all categories, err:%s", err.Error())
	}

	foundIndex = -1
	for i, c := range categories {
		if c.Id == category.Id {
			foundIndex = i
		}
	}

	if foundIndex >= 0 {
		t.Fatalf("invalid found index, found index:%d, found index id:%d, category id:%d",
			foundIndex, categories[foundIndex].Id, category.Id)
	}
}
