package test

import (
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
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

func TestPostReport(t *testing.T) {
	const Detail = "test"

	reasons, err := models.GetAllReportResons()
	if err != nil {
		t.Fatalf("Failed to get all report reasons")
	}

	var boards []models.Board
	err = HandleHttpRequest(http.MethodGet, "/v1/boards", &boards, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting all boards, err:%s", err.Error())
	}

	var posts []models.Post
	postsByBoardUrl := fmt.Sprintf("/v1/board/%d/posts", boards[0].Id)
	err = HandleHttpRequest(http.MethodGet, postsByBoardUrl, &posts, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting posts by board, err:%s", err.Error())
	}

	var report models.Report
	body := `{"detail":"` + Detail + `","report_reason_id":` + strconv.Itoa(reasons[0].Id) + `,"post_id":` + strconv.Itoa(posts[0].Id) + `}`
	err = HandleHttpRequest(http.MethodPost, "/v1/report/post", &report, body)
	if err != nil {
		t.Fatalf("Failed to handle request for creating post report, err:%s", err.Error())
	}

	var reports []models.Report
	reportsByPostUrl := fmt.Sprintf("/v1/reports/post/%d", posts[0].Id)
	err = HandleHttpRequest(http.MethodGet, reportsByPostUrl, &reports, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting reports by post, err:%s", err.Error())
	}

	foundIndex := -1
	for i, r := range reports {
		if r.Id == report.Id {
			foundIndex = i
		}
	}

	if foundIndex < 0 || reports[foundIndex].Detail != Detail {
		t.Fatalf("Invalid found report index, found index:%d, report id:%d", foundIndex, report.Id)
	}
}

func TestPostCommentReport(t *testing.T) {
	const Detail = "test"

	reasons, err := models.GetAllReportResons()
	if err != nil {
		t.Fatalf("Failed to get all report reasons")
	}

	var boards []models.Board
	err = HandleHttpRequest(http.MethodGet, "/v1/boards", &boards, "")
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

	var report models.Report
	body := `{"detail":"` + Detail + `","report_reason_id":` + strconv.Itoa(reasons[0].Id) + `,"post_comment_id":` + strconv.Itoa(comments[0].Id) + `}`
	err = HandleHttpRequest(http.MethodPost, "/v1/report/post_comment", &report, body)
	if err != nil {
		t.Fatalf("Failed to handle request for creating comment report, err:%s", err.Error())
	}

	var reports []models.Report
	reportsByCommentUrl := fmt.Sprintf("/v1/reports/post_comment/%d", comments[0].Id)
	err = HandleHttpRequest(http.MethodGet, reportsByCommentUrl, &reports, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting reports by comment, err:%s", err.Error())
	}

	foundIndex := -1
	for i, r := range reports {
		if r.Id == report.Id {
			foundIndex = i
		}
	}

	if foundIndex < 0 || reports[foundIndex].Detail != Detail {
		t.Fatalf("Invalid found report index, found index:%d, report id:%d", foundIndex, report.Id)
	}
}

func TestPostCommentReplyReport(t *testing.T) {
	const Detail = "test"

	reasons, err := models.GetAllReportResons()
	if err != nil {
		t.Fatalf("Failed to get all report reasons")
	}

	var boards []models.Board
	err = HandleHttpRequest(http.MethodGet, "/v1/boards", &boards, "")
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

	var replies []models.PostCommentReply
	repliesByCommentsUrl := fmt.Sprintf("/v1/post_comment/%d/post_comment_replies", comments[0].Id)
	err = HandleHttpRequest(http.MethodGet, repliesByCommentsUrl, &replies, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting replies by comment, err:%s", err.Error())
	}

	var report models.Report
	body := `{"detail":"` + Detail + `","report_reason_id":` + strconv.Itoa(reasons[0].Id) + `,"post_comment_reply_id":` + strconv.Itoa(replies[0].Id) + `}`
	err = HandleHttpRequest(http.MethodPost, "/v1/report/post_comment_reply", &report, body)
	if err != nil {
		t.Fatalf("Failed to handle request for creating reply report, err:%s", err.Error())
	}

	var reports []models.Report
	reportsByReplyUrl := fmt.Sprintf("/v1/reports/post_comment_reply/%d", replies[0].Id)
	err = HandleHttpRequest(http.MethodGet, reportsByReplyUrl, &reports, "")
	if err != nil {
		t.Fatalf("Failed to handle request for getting reports by reply, err:%s", err.Error())
	}

	foundIndex := -1
	for i, r := range reports {
		if r.Id == report.Id {
			foundIndex = i
		}
	}

	if foundIndex < 0 || reports[foundIndex].Detail != Detail {
		t.Fatalf("Invalid found report index, found index:%d, report id:%d", foundIndex, report.Id)
	}
}
