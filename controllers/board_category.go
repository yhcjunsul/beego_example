package controllers

import (
	"net/http"
	"strconv"

	"github.com/yhcjunsul/beego_example/models"
	"github.com/yhcjunsul/beego_example/utils"

	"github.com/astaxie/beego"
)

type BoardCategoryController struct {
	beego.Controller
}

func (this *BoardCategoryController) URLMapping() {
	this.Mapping("CreateBoardCategory", this.CreateBoardCategory)
	this.Mapping("GetAllBoardCategories", this.GetAllBoardCategories)
	this.Mapping("DeleteBoardCategory", this.DeleteBoardCategory)
}

// @Title Create board category
// @Summary Create board category
// @Description Create board category using name
// @Param   name	body	string	true	"Name of board category"
// @Success 200 {object} models.BoardCategory
// @Failure 400 Bad request, invalid body contents
// @Failure 500 Internal server error
// @Accept json
// @router /board_category [post]
func (this *BoardCategoryController) CreateBoardCategory() {
	category := models.BoardCategory{}

	if err := utils.UnmarshalRequestJson(this.Ctx.Input.RequestBody, &category); err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid body contents")
		return
	}

	if err := models.AddCategory(&category); err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	this.Data["json"] = category
	this.ServeJSON()

	beego.Info("new category, category id:", category.Id, ", category name:", category.Name)
}

// @Title Get all board category
// @Summary Get all board category
// @Success 200 {array} models.BoardCategory
// @Failure 500 Internal server error
// @Accept json
// @router /board_categories [get]
func (this *BoardCategoryController) GetAllBoardCategories() {
	categories, err := models.GetAllCategories()

	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}

	this.Data["json"] = categories

	beego.Info(this.Data["json"])
	this.ServeJSON()
}

// @Title Delete board category
// @Summary Delete board category by ID
// @Success 200
// @Failure 400 Bad request, invalid category id
// @Failure 500 Internal server error
// @Accept json
// @router /board_category/:id:int [delete]
func (this *BoardCategoryController) DeleteBoardCategory() {
	id_param := this.Ctx.Input.Param(":id")

	id, err := strconv.Atoi(id_param)
	if err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusBadRequest, "Bad request, invalid category id")
		return
	}

	if err := models.DeleteCategoy(id); err != nil {
		utils.SetErrorStatus(this.Ctx, http.StatusInternalServerError, "Internal server error")
		return
	}
}
