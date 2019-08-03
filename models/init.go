package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(BoardCategory), new(Board), new(Post), new(PostComment), new(PostCommentReply))
}

func InitTestSetting() {
	category := BoardCategory{Name: "category"}
	AddCategory(&category)

	numberSlice := []string{"first_", "second_", "third_"}

	for i, numStr := range numberSlice {
		board := Board{Name: numStr + "board", BoardCategory: &category}
		AddBoard(&board)

		post_board, _ := FindBoardById(i + 1)

		post := Post{Title: numStr + "post", Contents: numStr + "contents", Ip: "127.0.0.1", Board: post_board}
		AddPost(&post)

		comment := PostComment{Contents: numStr + "comment", Ip: "127.0.0.1", Post: &post}
		AddPostComment(&comment)
	}
}
