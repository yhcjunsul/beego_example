package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/yhcjunsul/beego_example/models"

	"github.com/astaxie/beego"
)

type BoardCategoryController struct {
	beego.Controller
}

// @Title Create board category
// @Summary Create board category
// @Description Create board category using name
// @Param   name	body	string	true	"Name of board category"
// @Success 200
// @Failure 400 Bad Request, Duplicate name
// @Failure 422 Unprocessable Entity
// @Accept json
// @router /board_category [post]
func (this *BoardCategoryController) CreateBoardCategory() {
	category := models.BoardCategory{}

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &category); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad Request"))
		return
	}

	if err := models.AddCategory(&category); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	beego.Info("new category, category name:%s", category.Name)
}

// @Title Get all board category
// @Summary Get all board category
// @Success 200 {object}array models.BoardCategory
// @Failure 404 Not found
// @Accept json
// @router /board_categories [get]
func (this *BoardCategoryController) GetAllBoardCategories() {
	categories, err := models.GetAllCategories()

	if err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("get all categories error:" + err.Error()))
		return
	}

	this.Data["json"] = categories

	beego.Info(this.Data["json"])
	this.ServeJSON()
}

// @Title Delete board category
// @Summary Delete board category by ID
// @Success 200
// @Failure 404 Not found
// @Accept json
// @router /board_category/:id [delete]
func (this *BoardCategoryController) DeleteBoardCategory() {
	id_param := this.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(id_param)
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("Bad request"))
		return
	}

	if err := models.DeleteCategoy(id); err != nil {
		this.Ctx.Output.SetStatus(404)
		this.Ctx.Output.Body([]byte("Delete categories error:" + err.Error()))
		return
	}
}
