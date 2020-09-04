// Package controllers gin's controller
package controllers

import (
	"net/http"

	"github.com/103cuong/gorm_kit/models"
	"github.com/gin-gonic/gin"
)

// GetCats Get all cats.
func GetCats(ctx *gin.Context) {
	var cats []models.Cat
	err := models.GetCats(&cats)
	if err != nil {
		// FIXME: response status + error.
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, cats)
}

// CreateCat Create new cat.
func CreateCat(ctx *gin.Context) {
	var cat models.Cat
	err := ctx.BindJSON(&cat)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = models.CreateCat(&cat)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusCreated, cat)
}

// GetCatByID Get cat by id.
func GetCatByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var cat models.Cat
	err := models.GetCatByID(&cat, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, cat)
}

// UpdateCat Update cat.
func UpdateCat(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var cat models.Cat
	err := models.GetCatByID(&cat, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, cat)
		return
	}
	err = ctx.BindJSON(&cat)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = models.UpdateCat(&cat)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, cat)
}

// DeleteCat Delete cat.
func DeleteCat(ctx *gin.Context) {
	var cat models.Cat
	id := ctx.Params.ByName("id")
	err := models.DeleteCat(&cat, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"deleted": true})
}
