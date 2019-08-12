package controllers

import (
	"net/http"
	"strconv"

	"github.com/yhcjunsul/beego_example/models"
	"github.com/yhcjunsul/beego_example/utils"

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
// @Summary Create board
// @Description Create board using name and category id
// @Param   name				body	string	true	"Name of board category"
// @Param	board_category_id 	path	int		true	"category id"
// @Success 200
// @Failure 400 Bad Request, invalid category id
// @Failure 400 Bad request, invalid body contents
// @Failure 404 Not found category
// @Failure 500 internal server error
// @Accept json
// @router /board_category/:categoryId:int/board [post]
func (this *BoardController) CreateBoard() {
	board := models.Board{}

	category_id, err := strconv.Atoi(this.Ctx.Input.Param(":categoryId"))
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad Request, invalid category id")
		return
	}

	if err := utils.UnmarshalRequestJson(this.Ctx.Input.RequestBody, &board); err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad Request, invalid body contents")
		return
	}

	board.BoardCategory, err = models.FindCategoryById(category_id)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusNotFound, "Not found category")
		return
	}

	if err = models.AddBoard(&board); err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	beego.Info("new board, board name:%s", board.Name)
}

// @Title Get boards by category
// @Summary Get boards by category
// @Param board_category_id		path	int		true	"category id"
// @Success 200 {array} models.Board
// @Failure 400 Bad request, invalid category id
// @Failure 404 Not found category
// @Accept json
// @router /board_category/:categoryId:int/boards [get]
func (this *BoardController) GetBoardsByCategory() {
	category_id, err := strconv.Atoi(this.Ctx.Input.Param(":categoryId"))
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid category id")
		return
	}

	category := models.BoardCategory{Id: category_id}
	boards, err := models.GetBoardsByCategory(&category)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusNotFound, "Not found category")
		return
	}

	this.Data["json"] = boards
	beego.Info(this.Data["json"])
	this.ServeJSON()
}

// @Title Get all boards
// @Summary Get all boards
// @Success 200 {object}array models.Board
// @Failure 500 Internal server error
// @Accept json
// @router /boards [get]
func (this *BoardController) GetAllBoards() {
	boards, err := models.GetAllBoards()

	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	this.Data["json"] = boards

	beego.Info(this.Data["json"])
	this.ServeJSON()
}

// @Title Delete board
// @Summary Delete board by ID
// @Success 200
// @Failure 400 Bad request, invalid board id
// @Failure 500 Internal server error
// @Accept json
// @router /board/:id:int [delete]
func (this *BoardController) DeleteBoard() {
	id_param := this.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(id_param)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid board id")
		return
	}

	if err := models.SetBoardDeleteFlag(id, true); err != nil {
		utils.SetErrorStatus(this.Ctx, 500, "Internal server error")
		return
	}
}
