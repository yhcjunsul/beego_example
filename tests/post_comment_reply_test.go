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

func TestPostCommentReply(t *testing.T) {
	const contents = "test"

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

	var comments []models.PostComment
	commentsByPostUrl := fmt.Sprintf("/v1/post/%d/post_comments", posts[0].Id)
	err = HandleHttpRequest(http.MethodGet, commentsByPostUrl, &comments, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting comments by post, err:%s", err.Error())
	}

	var reply models.PostCommentReply
	body := `{"contents":"` + contents + `"}`
	replyCreateUrl := fmt.Sprintf("/v1/post_comment/%d/post_comment_reply", comments[0].Id)
	err = HandleHttpRequest(http.MethodPost, replyCreateUrl, &reply, body)
	if err != nil {
		t.Fatalf("Failed to handle request for creating reply, err:%s", err.Error())
	}

	var replies []models.PostCommentReply
	repliesByCommentUrl := fmt.Sprintf("/v1/post_comment/%d/post_comment_replies", comments[0].Id)
	err = HandleHttpRequest(http.MethodGet, repliesByCommentUrl, &replies, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting replies by comment, err:%s", err.Error())
	}

	foundIndex := -1
	for i, r := range replies {
		if r.Id == reply.Id {
			foundIndex = i
		}
	}

	if foundIndex < 0 || replies[foundIndex].Contents != contents {
		t.Fatalf("Invalid found reply index, found index:%d, reply id:%d", foundIndex, reply.Id)
	}

	deleteUrl := fmt.Sprintf("/v1/post_comment_reply/%d", reply.Id)
	err = HandleHttpRequest(http.MethodDelete, deleteUrl, nil, "")
	if err != nil {
		t.Fatalf("Failed to handle request for deleting reply, err:%s", err.Error())
	}

	replies = nil
	err = HandleHttpRequest(http.MethodGet, repliesByCommentUrl, &replies, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting replies by comment, err:%s", err.Error())
	}

	foundIndex = -1
	for i, r := range replies {
		if r.Id == reply.Id {
			foundIndex = i
		}
	}

	if foundIndex < 0 || replies[foundIndex].IsDeleted != true {
		t.Fatalf("Invalid found reply index, found index:%d, reply id:%d", foundIndex, reply.Id)
	}
}
