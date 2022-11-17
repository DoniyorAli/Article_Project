package handlers

import (
	"net/http"
	"regexp"
	"strconv"

	"UacademyGo/Article/models"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// * ================== Create Article ======================
// CreateArticle godoc
// @Summary     Create article
// @Description Create a new article
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       article body models.CreateModelArticle true "article body" //? True false nimaga kerak ahir modulda required borku
// @Success     201 {object} models.JSONRespons{data=models.Article} //? interfeysni overright qivoradi
// @Failure     400 {object} models.JSONErrorRespons                 //? yani        bizani    sructuramizni interfeysni orniga qoyvoradi
// @Router      /v2/article [post]
func (h *Handler) CreateArticle(ctx *gin.Context){
	var body models.CreateModelArticle
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{Error: err.Error()})
		return
	}

	id := uuid.New()

	err := h.Stg.AddNewArticle(id.String(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	article, err := h.Stg.GetArticleById(id.String())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.JSONRespons{
		Message: "Article successfully created",
		Data:    article,
	})
}

// * ==================== GetArticleById ====================
// GetArticleById godoc
// @Summary     get article by id
// @Description get a new article
// @Tags        articles
// @Accept      json
// @Param       id path string true "Article ID"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.GetByIDArticleModel}
// @Failure     404 {object} models.JSONErrorRespons
// @Router      /v2/article/{id} [get]
func (h *Handler) GetArticleById(ctx *gin.Context) {
	idStr := ctx.Param("id")

//TODO UUID validation
	// // if len(idStr) == 37 && 				//? ---> 4c0fb410-76e1-4ef8-8317-e4f4a78f3d76
	// isTrue := GetUUIDValidator(idStr)
	// if isTrue == true{
		 
	// }
	
	article, err := h.Stg.GetArticleById(idStr) //? qanaqa qilip inMga bervorvotti
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}


	
	ctx.JSON(http.StatusOK, models.JSONRespons{
		Message: "passed successfully",
		Data: article,
	})
}

// * ==================== GetArticleList ====================
// GetArticleList godoc
// @Summary     List articles
// @Description get articles
// @Tags        articles
// @Accept      json
// @Produce     json
// @Param       offset query    int    false "0"
// @Param       limit  query    int    false "10"
// @Param       search query    string false "smth"
// @Success     200    {object} models.JSONRespons{data=[]models.Article}
// @Router      /v2/article [get]
func (h *Handler) GetArticleList(ctx *gin.Context) {
	offsetStr := ctx.DefaultQuery("offset", "0")
	limitStr := ctx.DefaultQuery("limit", "10")
	searchStr := ctx.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	
	articleList, err := h.Stg.GetArticleList(offset, limit, searchStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, models.JSONRespons{
		Message: "OK",
		Data:    articleList,
	})
}

// * ==================== UpdateArticle ====================
// UpdateArticle godoc
// @Summary     Update article
// @Description Update a new article
// @Tags        articles
// @Accept      json
// @Param       article body models.UpdateArticleResponse true "updating article"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.Article}
// @Failure     400 {object} models.JSONErrorRespons
// @Router      /v2/article [put]
func (h *Handler) UpdateArticle(ctx *gin.Context) {
	var body models.UpdateArticleResponse
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{Error: err.Error()})
		return
	}

	
	err :=  h.Stg.UpdateArticle(body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Message: "error",
			Error: err.Error(),
		})
		return
	}

	article, err := h.Stg.GetArticleById(body.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONRespons{
		Message: "Article successfully updated",
		Data: article,
	})
}

// * ==================== DeleteArticle ====================
// DeleteArticle godoc
// @Summary     Delete article
// @Description delete article
// @Tags        articles
// @Accept      json
// @Param       id path string true "Article ID"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.Article}
// @Failure     400 {object} models.JSONErrorRespons
// @Router      /v2/article/{id} [delete]
func (h *Handler) DeleteArticle(ctx *gin.Context) {
	idStr := ctx.Param("id")

	
	article, err := h.Stg.GetArticleById(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Message: "article already deleted or not found or you entered wrong ID",
			Error: err.Error(),
		})
		return
	}

	err = h.Stg.DeleteArticle(article.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNotFound, models.JSONRespons{
		Message: "Article suucessfully deleted",
		Data:    article,
	})
}

// * ==================== PingPong ====================
func (h *Handler) Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

//*===========================================================
//! Checking the UUID ...
func GetUUIDValidator(text string) bool {
    r, _ := regexp.Compile("/[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89aAbB][a-f0-9]{3}-[a-f0-9]{12}/")
    return r.Match([]byte(text))
}
