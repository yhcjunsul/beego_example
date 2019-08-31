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

func TestPostComment(t *testing.T) {
	const Contents = "test"

	var boards []models.Board
	err := HandleHttpRequest(http.MethodGet, "/v1/boards", &boards, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting all boards, err:%s", err.Error())
	}

	var posts []models.Post
	postsByBoardUrl := fmt.Sprintf("/v1/board/%d/posts", boards[0].Id)
	err = HandleHttpRequest(http.MethodGet, postsByBoardUrl, &posts, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting posts by board, err:%s", err.Error())
	}

	var comment models.PostComment
	body := `{"contents":"` + Contents + `"}`
	commentCreateUrl := fmt.Sprintf("/v1/post/%d/post_comment", posts[0].Id)
	err = HandleHttpRequest(http.MethodPost, commentCreateUrl, &comment, body)
	if err != nil {
		t.Fatalf("Failed to handle request for creating comment, err:%s", err.Error())
	}

	var comments []models.PostComment
	commentsByPostUrl := fmt.Sprintf("/v1/post/%d/post_comments", posts[0].Id)
	err = HandleHttpRequest(http.MethodGet, commentsByPostUrl, &comments, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting comments by post, err:%s", err.Error())
	}

	foundIndex := -1
	for i, c := range comments {
		if c.Id == comment.Id {
			foundIndex = i
		}
	}

	if foundIndex < 0 || comments[foundIndex].Contents != Contents {
		t.Fatalf("Invalid found comment index, found index:%d, comment id:%d", foundIndex, comment.Id)
	}

	deleteUrl := fmt.Sprintf("/v1/post_comment/%d", comment.Id)
	err = HandleHttpRequest(http.MethodDelete, deleteUrl, nil, "")
	if err != nil {
		t.Fatalf("Failed to handle request for deleting comment, err:%s", err.Error())
	}

	comments = nil
	err = HandleHttpRequest(http.MethodGet, commentsByPostUrl, &comments, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting comments by post, err:%s", err.Error())
	}

	foundIndex = -1
	for i, c := range comments {
		if c.Id == comment.Id {
			foundIndex = i
		}
	}

	if foundIndex < 0 || comments[foundIndex].IsDeleted != true {
		t.Fatalf("Invalid found comment index, found index:%d, comment id:%d", foundIndex, comment.Id)
	}
}
