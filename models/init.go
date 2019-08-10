package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(BoardCategory), new(Board), new(Post), new(PostComment), new(PostCommentReply), new(Report), new(ReportReason))
}

func InitTestSetting() {
	category := BoardCategory{Name: "category"}
	AddCategory(&category)

	foundCategory, _ := FindCategoryById(1)

	numberSlice := []string{"first_", "second_", "third_"}

	for i, numStr := range numberSlice {
		board := Board{Name: numStr + "board", BoardCategory: foundCategory}
		AddBoard(&board)

		id_index := i + 1

		post_board, _ := FindBoardById(id_index)

		post := Post{Title: numStr + "post", Contents: numStr + "contents", Ip: "127.0.0.1", Board: post_board}
		AddPost(&post)

		foundPost, _ := FindPostById(id_index)

		comment := PostComment{Contents: numStr + "comment", Ip: "127.0.0.1", Post: foundPost}
		AddPostComment(&comment)

		foundComment, _ := FindPostCommentById(id_index)

		reply := PostCommentReply{Contents: numStr + "reply", Ip: "127.0.0.1", PostComment: foundComment}
		AddPostCommentReply(&reply)

		foundReply, _ := FindPostCommentReplyById(id_index)

		reason := ReportReason{Contents: numStr + "reason"}
		AddReportReason(&reason)

		foundReason, _ := FindReportReasonById(id_index)

		postReport := Report{Detail: numStr + "report", Ip: "127.0.0.1", ReportReason: foundReason, Post: foundPost}
		AddReport(&postReport)

		commentReport := Report{Detail: numStr + "report", Ip: "127.0.0.1", ReportReason: foundReason, PostComment: foundComment}
		AddReport(&commentReport)

		replyReport := Report{Detail: numStr + "report", Ip: "127.0.0.1", ReportReason: foundReason, PostCommentReply: foundReply}
		AddReport(&replyReport)
	}
}
