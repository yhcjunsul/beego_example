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

func TestPost(t *testing.T) {
	const title = "test"
	const contents = "test"

	var boards []models.Board
	err := HandleHttpRequest(http.MethodGet, "/v1/boards", &boards, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting all boards, err:%s", err.Error())
	}

	var post models.Post
	body := `{"title":"` + title + `","contents":"` + contents + `"}`
	postCreateUrl := fmt.Sprintf("/v1/board/%d/post", boards[0].Id)
	err = HandleHttpRequest(http.MethodPost, postCreateUrl, &post, body)
	if err != nil {
		t.Fatalf("Failed to handle request for creating post, err:%s", err.Error())
	}

	var posts []models.Post
	postsByBoardUrl := fmt.Sprintf("/v1/board/%d/posts", boards[0].Id)
	err = HandleHttpRequest(http.MethodGet, postsByBoardUrl, &posts, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting posts by board, err:%s", err.Error())
	}

	foundIndex := -1
	for i, p := range posts {
		if p.Id == post.Id {
			foundIndex = i
		}
	}

	if foundIndex < 0 || posts[foundIndex].Title != title {
		t.Fatalf("Invalid found post index, found index:%d, post id:%d", foundIndex, post.Id)
	}

	var postFromServer models.Post
	postByIdUrl := fmt.Sprintf("/v1/post/%d", post.Id)
	err = HandleHttpRequest(http.MethodGet, postByIdUrl, &postFromServer, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting post by id, err:%s", err.Error())
	}

	if postFromServer.Title != title {
		t.Fatalf("Invaild post from server, post from server id:%d, post id :%d",
			postFromServer.Id, post.Id)
	}

	deleteUrl := fmt.Sprintf("/v1/post/%d", post.Id)
	err = HandleHttpRequest(http.MethodDelete, deleteUrl, nil, "")
	if err != nil {
		t.Fatalf("Failed to handle request for deleting post, err:%s", err.Error())
	}

	posts = nil
	err = HandleHttpRequest(http.MethodGet, postsByBoardUrl, &posts, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting posts by board, err:%s", err.Error())
	}

	foundIndex = -1
	for i, p := range posts {
		if p.Id == post.Id {
			foundIndex = i
		}
	}

	if foundIndex < 0 || posts[foundIndex].IsDeleted != true {
		t.Fatalf("Invalid found index, found index:%d, post id:%d", foundIndex, post.Id)
	}

	err = HandleHttpRequest(http.MethodGet, postByIdUrl, &postFromServer, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting post by id, err:%s", err.Error())
	}

	if postFromServer.IsDeleted != true {
		t.Fatalf("Invalid post from server, post from server id:%d,delete flag of post from server:%v, post id:%d",
			postFromServer.Id, postFromServer.IsDeleted, post.Id)
	}
}
