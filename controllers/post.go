package controllers

import (
	"net/http"
	"strconv"

	"github.com/yhcjunsul/beego_example/models"
	"github.com/yhcjunsul/beego_example/utils"

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
// @Summary Create post
// @Param   title		body	string	true	"Title of post"
// @Param   contents	body	string	true	"Contents of post"
// @Param  	board_id	path	int		true	"board id"
// @Success 200
// @Failure 400 Bad request, invalid board id
// @Failure 400 Bad request, invalid body contents
// @Failure 404 Not found board
// @Failure 500 Internal server error
// @Accept json
// @router /board/:boardId:int/post [post]
func (this *PostController) CreatePost() {
	post := models.Post{}

	board_id, err := strconv.Atoi(this.Ctx.Input.Param(":boardId"))
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid board id")
		return
	}

	if err := utils.UnmarshalRequestJson(this.Ctx.Input.RequestBody, &post); err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid body contents")
		return
	}

	post.Board, err = models.FindBoardById(board_id)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusNotFound, "Not found board")
		return
	}

	ip := this.Ctx.Input.IP()
	post.Ip = ip

	if err = models.AddPost(&post); err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	beego.Info("new post, post title:%s", post.Title)
}

// @Title Get posts by board
// @Summary Get posts by board
// @Param board_id 	path	int	true	"board id"
// @Success 200 {array} models.Post
// @Failure 400 Bad request, invalid board id
// @Failure 404 Not found board
// @Accept json
// @router /board/:boardId:int/posts [get]
func (this *PostController) GetPostsByBoard() {
	board_id, err := strconv.Atoi(this.Ctx.Input.Param(":boardId"))
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid board id")
		return
	}

	board := models.Board{Id: board_id}
	posts, err := models.GetPostsByBoard(&board)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusNotFound, "Not found board")
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
// @Failure 400 Bad request, invalid post id
// @Failure 404 Not found post
// @Failure 500 Internal server error
// @Accept json
// @router /post/:id:int [get]
func (this *PostController) GetPostById() {
	id_param := this.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(id_param)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid post id")
		return
	}

	post, err := models.FindPostById(id)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusNotFound, "Not found post")
		return
	}

	err = models.IncreasePostViewCount(id)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	this.Data["json"] = post

	beego.Info(this.Data["json"])
	this.ServeJSON()
}

// @Title Delete post
// @Summary Delete post by ID
// @Success 200
// @Failure 400 Bad request, invalid post id
// @Failure 500 Internal server error
// @Accept json
// @router /post/:id:int [delete]
func (this *PostController) DeletePost() {
	id_param := this.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(id_param)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid post id")
		return
	}

	if err := models.SetPostDeleteFlag(id, true); err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}
}
