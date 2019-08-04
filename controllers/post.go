package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/yhcjunsul/beego_example/models"

	"github.com/astaxie/beego"
)

type PostController struct {
	beego.Controller
}

func (this *PostController) URLMapping() {
	this.Mapping("CreatePost", this.CreatePost)
	this.Mapping("GetPostsByBoard", this.GetPostsByBoard)
	this.Mapping("GetPostById", this.GetPostById)
	this.Mapping("DeletePost", this.DeletePost)
}

// @Title Create post
// @Summary /board/{board_id}/post Create post
// @Param   title		body	string	true	"Title of post"
// @Param   contents	body	string	true	"Contents of post"
// @Param 	ip			body	string	true	"ip of user"
// @Param  	board_id	path	int		true	"board id"
// @Success 200
// @Failure 400 Bad Request
// @Failure 404 Not found
// @Accept json
// @router /*/post [post]
func (this *PostController) CreatePost() {
	post := models.Post{}

	splat := this.Ctx.Input.Param(":splat")
	// splat == "board/:id"
	splat_splits := strings.Split(splat, "/")

	if len(splat_splits) != 2 {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	if splat_splits[0] != "board" {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	board_id, err := strconv.Atoi(splat_splits[1])
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &post); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	post.Board, err = models.FindBoardById(board_id)
	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Not found"))
		return
	}

	if err = models.AddPost(&post); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	beego.Info("new post, post title:%s", post.Title)
}

// @Title Get posts by board
// @Summary /board/{board_id}/posts Get posts by board
// @Param board_id 	path	int	true	"board id"
// @Success 200 {array} models.Post
// @Failure 404 Not found
// @Accept json
// @router /*/posts [get]
func (this *PostController) GetPostsByBoard() {
	splat := this.Ctx.Input.Param(":splat")
	// splat == "board/:id"
	splat_splits := strings.Split(splat, "/")

	if len(splat_splits) != 2 {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	if splat_splits[0] != "board" {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	board_id, err := strconv.Atoi(splat_splits[1])
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	board := models.Board{Id: board_id}
	posts, err := models.GetPostsByBoard(&board)
	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Not found board"))
		return
	}

	this.Data["json"] = posts
	beego.Info(this.Data["json"])
	this.ServeJSON()
}

// @Title Get post by id
// @Summary Get post by id
// @Param id	path	int		true	"id of post"
// @Success 200 {object} models.Post
// @Failure 404 Not found
// @Accept json
// @router /post/:id:int [get]
func (this *PostController) GetPostById() {
	id_param := this.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(id_param)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad request"))
	}

	post, err := models.FindPostById(id)
	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("get post error:" + err.Error()))
		return
	}

	err = models.IncreasePostViewCount(id)
	if err != nil {
		this.Ctx.Output.SetStatus(500)
		this.Ctx.Output.Body([]byte("get post error:" + err.Error()))
		return
	}

	this.Data["json"] = post

	beego.Info(this.Data["json"])
	this.ServeJSON()
}

// @Title Delete post
// @Summary Delete post by ID
// @Success 200
// @Failure 404 Not found
// @Accept json
// @router /post/:id:int [delete]
func (this *PostController) DeletePost() {
	id_param := this.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(id_param)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad request"))
	}

	if err := models.SetPostDeleteFlag(id, true); err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Delete post error:" + err.Error()))
		return
	}
}
