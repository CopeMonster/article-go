package article

import (
	"article-go/internal/domain"
	"article-go/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	articleService service.ArticleService
}

func NewHandler(articleService service.ArticleService) *Handler {
	return &Handler{
		articleService: articleService,
	}
}

func (h *Handler) GetArticle(ctx *gin.Context) {
	articles, err := h.articleService.GetArticle(ctx.Request.Context())
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, domain.GetArticlesResponse{
		Articles: articles,
	})
}

func (h *Handler) GetArticleByID(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	article, err := h.articleService.GetArticleByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, domain.GetArticleResponse{
		Article: article,
	})
}

func (h *Handler) CreateArticle(ctx *gin.Context) {
	createInp := new(domain.CreateArticleInput)

	if err := ctx.BindJSON(createInp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := ctx.MustGet(service.CtxUserKey).(*domain.User)

	if err := h.articleService.CreateArticle(ctx.Request.Context(), user, createInp.Title, createInp.Text); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) UpdateArticle(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	updateInp := new(domain.UpdateArticleInput)

	if err := ctx.BindJSON(updateInp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := ctx.MustGet(service.CtxUserKey).(*domain.User)

	if err := h.articleService.UpdateArticle(ctx.Request.Context(), user, id, updateInp.Title, updateInp.Text); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) DeleteArticle(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := ctx.MustGet(service.CtxUserKey).(*domain.User)

	if err := h.articleService.DeleteArticle(ctx.Request.Context(), user, id); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
