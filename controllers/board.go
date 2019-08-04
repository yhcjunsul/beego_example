package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/yhcjunsul/beego_example/models"

	"github.com/astaxie/beego"
)

type BoardController struct {
	beego.Controller
}

func (this *BoardController) URLMapping() {
	this.Mapping("CreateBoard", this.CreateBoard)
	this.Mapping("GetBoardsByCategory", this.GetBoardsByCategory)
	this.Mapping("GetAllBoards", this.GetAllBoards)
	this.Mapping("DeleteBoard", this.DeleteBoard)
}

// @Title Create board
// @Summary /board_category/{board_category_id}/board Create board
// @Description Create board using name and category id
// @Param   name				body	string	true	"Name of board category"
// @Param	board_category_id 	path	int		true	"category id"
// @Success 200
// @Failure 400 Bad Request
// @Failure 404 Not found
// @Accept json
// @router /*/board [post]
func (this *BoardController) CreateBoard() {
	board := models.Board{}

	splat := this.Ctx.Input.Param(":splat")
	// splat == "board_category/:id"
	splat_splits := strings.Split(splat, "/")

	if len(splat_splits) != 2 {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	if splat_splits[0] != "board_category" {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	category_id, err := strconv.Atoi(splat_splits[1])
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &board); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	board.BoardCategory, err = models.FindCategoryById(category_id)
	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Not found"))
		return
	}

	if err = models.AddBoard(&board); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	beego.Info("new board, board name:%s", board.Name)
}

// @Title Get boards by category
// @Summary /board_category/{board_category_id}/boards Get boards by category
// @Param board_category_id		path	int		true	"category id"
// @Success 200 {array} models.Board
// @Failure 404 Not found
// @Accept json
// @router /*/boards [get]
func (this *BoardController) GetBoardsByCategory() {
	splat := this.Ctx.Input.Param(":splat")
	// splat == "board_category/:id"
	splat_splits := strings.Split(splat, "/")

	if len(splat_splits) != 2 {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	if splat_splits[0] != "board_category" {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	category_id, err := strconv.Atoi(splat_splits[1])
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	category := models.BoardCategory{Id: category_id}
	boards, err := models.GetBoardsByCategory(&category)
	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Not found category"))
		return
	}

	this.Data["json"] = boards
	beego.Info(this.Data["json"])
	this.ServeJSON()
}

// @Title Get all boards
// @Summary Get all boards
// @Success 200 {object}array models.Board
// @Failure 404 Not found
// @Accept json
// @router /boards [get]
func (this *BoardController) GetAllBoards() {
	boards, err := models.GetAllBoards()

	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("get all  boards error:" + err.Error()))
		return
	}

	this.Data["json"] = boards

	beego.Info(this.Data["json"])
	this.ServeJSON()
}

// @Title Delete board
// @Summary Delete board by ID
// @Success 200
// @Failure 404 Not found
// @Accept json
// @router /board/:id:int [delete]
func (this *BoardController) DeleteBoard() {
	id_param := this.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(id_param)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad request"))
	}

	if err := models.SetBoardDeleteFlag(id, true); err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Delete boards error:" + err.Error()))
		return
	}
}
