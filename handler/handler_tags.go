package handler

import (
	"API3/data/request"
	"API3/data/response"
	"API3/helper"
	"API3/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

type TagsHandler struct {
	tagsService service.TagsService
}

func NewTagsHandler(service service.TagsService) *TagsHandler {
	return &TagsHandler{
		tagsService: service,
	}
}

// Create Handler

func (p *TagsHandler) Create(ctx *gin.Context) {
	log.Info().Msg("create tags")
	createTagsRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagsRequest)
	helper.ErrorPanic(err)

	p.tagsService.Create(createTagsRequest)
	webResponse := response.Response{
		Status: "Ok",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

//Update Handler

func (p *TagsHandler) Update(ctx *gin.Context) {
	log.Info().Msg("update tags")
	updateTagsRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagsRequest.Id = id

	p.tagsService.Update(updateTagsRequest)

	webResponse := response.Response{
		Status: "Ok",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// Delete Handler
func (p *TagsHandler) Delete(ctx *gin.Context) {
	log.Info().Msg("delete tags")
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	p.tagsService.Delete(id)

	webResponse := response.Response{
		Status: "Ok",
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

//FindById Handler

func (p *TagsHandler) FindByID(ctx *gin.Context) {
	log.Info().Msg("find by id tags")
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	tagResponse := p.tagsService.FindById(id)

	webResponse := response.Response{

		Data: tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

//FindALL

func (p *TagsHandler) FindAll(ctx *gin.Context) {
	log.Info().Msg("find all tags")
	tagResponse := p.tagsService.FindAll()
	webResponse := response.Response{
		Data: tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
