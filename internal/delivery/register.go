package delivery

import (
	"article-go/internal/delivery/article"
	"article-go/internal/delivery/auth"
	"article-go/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAuthHTTPEndpoints(router *gin.Engine, authService service.AuthService) {
	handler := auth.NewHandler(authService)

	authEndpoints := router.Group("/auth")

	{
		authEndpoints.POST("/sign-up", handler.SignUp)
		authEndpoints.POST("/sign-in", handler.SignIn)
	}
}

func RegisterArticleHTTPEndpoints(router *gin.Engine, authService service.AuthService, articleService service.ArticleService) {
	handler := article.NewHandler(articleService)

	articleEndPoints := router.Group("/article")

	articleEndPoints.Use(func(c *gin.Context) {
		if method := c.Request.Method; method == http.MethodPost || method == http.MethodPut || method == http.MethodDelete {
			auth.NewMiddleware(authService)(c)
		} else {
			c.Next()
		}
	})

	{
		articleEndPoints.GET("", handler.GetArticle)
		articleEndPoints.GET("/:id", handler.GetArticleByID)
		articleEndPoints.POST("", handler.CreateArticle)
		articleEndPoints.PUT("/:id", handler.UpdateArticle)
		articleEndPoints.DELETE("/:id", handler.DeleteArticle)
	}
}
