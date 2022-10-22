package handlers

import (
	"net/http"

	"UacademyGo/Article/models"
	"UacademyGo/Article/storage/inmemory"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// * ================== Create Author =========================
// CreateAuthor godoc
// @Summary     Create author
// @Description Create a new author
// @Tags        authors
// @Accept      json
// @Param       author body models.CreateModelAuthor true "author body" //? True false nimaga kerak ahir modulda required borku
// @Produce     json
// @Success     201 {object} models.JSONRespons{data=models.Author} //? interfeysni overright qivoradi
// @Failure     400 {object} models.JSONErrorRespons                //? yani        bizani    sructuramizni interfeysni orniga qoyvoradi
// @Router      /v2/author [post]
func (h Handler) CreateAuthor(ctx *gin.Context){
	inM := inmemory.InMemory{}
	var body models.CreateModelAuthor
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{Error: err.Error()})
		return
	}

	id := uuid.New()

	err := inM.AddAuthor(id.String(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	author, err := inM.GetAuthorById(id.String())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.JSONRespons{
		Message: "Author | GetList",
		Data:    author,
	})
}

// * ==================== Get Author By Id ====================
// GetAuthorById godoc
// @Summary     get author by id
// @Description get a new author
// @Tags        authors
// @Accept      json
// @Param       id path string true "Article ID"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.Author}
// @Failure     404 {object} models.JSONErrorRespons
// @Router      /v2/author/{id} [get]
func (h Handler) GetAuthorById(ctx *gin.Context) {
	inM := inmemory.InMemory{}
	idStr := ctx.Param("id")

//TODO UUID validation
	// // if len(idStr) == 37 && 				//? ---> 4c0fb410-76e1-4ef8-8317-e4f4a78f3d76
	// isTrue := GetUUIDValidator(idStr)
	// if isTrue == true{
		 
	// }

	author, err := inM.GetAuthorById(idStr) //? qanaqa qilip storagega bervorvotti
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}


	
	ctx.JSON(http.StatusOK, models.JSONRespons{
		Message: "passed successfully",
		Data: author,
	})
}

// * ==================== Get Article List ====================
// GetArticleList godoc
// @Summary     List authors
// @Description get authors
// @Tags        authors
// @Accept      json
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=[]models.Author}
// @Router      /v2/author [get]
func (h Handler) GetAuthorList(ctx *gin.Context) {
	inM := inmemory.InMemory{}
	authorList, err := inM.GetAuthorList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, models.JSONRespons{
		Message: "OK",
		Data:    authorList,
	})
}

// * ==================== Update Author =======================
// UpdateAuthor godoc
// @Summary     Update author
// @Description Update a new author
// @Tags        authors
// @Accept      json
// @Param       author body models.UpdateAuthorResponse true "updating author"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.Author}
// @Failure     400 {object} models.JSONErrorRespons
// @Router      /v2/author [put]
func (h Handler) UpdateAuthor(ctx *gin.Context) {
	inM := inmemory.InMemory{}
	var body models.UpdateAuthorResponse
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{Error: err.Error()})
		return
	}

	err :=  inM.UpdateAuthor(body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Message: "Storage error",
			Error: err.Error(),
		})
		return
	}

	author, err := inM.GetAuthorById(body.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.JSONRespons{
		Message: "Author successfully updated",
		Data: author,
	})
}

// * ==================== Delete Author =======================
// DeleteAuthor godoc
// @Summary     Delete author
// @Description delete author
// @Tags        authors
// @Accept      json
// @Param       id path string true "Author ID"
// @Produce     json
// @Success     200 {object} models.JSONRespons{data=models.Author}
// @Failure     400 {object} models.JSONErrorRespons
// @Router      /v2/author/{id} [delete]
func (h Handler) DeleteAuthor(ctx *gin.Context) {
	inM := inmemory.InMemory{}
	idStr := ctx.Param("id")

	author, err := inM.GetAuthorById(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.JSONErrorRespons{
			Message: "author already deleted or not found or you entered wrong ID",
			Error: err.Error(),
		})
		return
	}

	err = inM.DeleteAuthor(author.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.JSONErrorRespons{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNotFound, models.JSONRespons{
		Message: "Author suucessfully deleted",
		Data:    author,
	})
}

// // * ==================== PingPong ====================
// func Pong(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": "pong",
// 	})
// }

// //*==========================================================
// //! Checking the UUID ...
// func GetUUIDValidator(text string) bool {
//     r, _ := regexp.Compile("/[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89aAbB][a-f0-9]{3}-[a-f0-9]{12}/")
//     return r.Match([]byte(text))
// }
