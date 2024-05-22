package auth

import (
	"article-go/internal/domain"
	"article-go/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Middleware struct {
	authService service.AuthService
}

func NewMiddleware(authService service.AuthService) gin.HandlerFunc {
	return (&Middleware{
		authService: authService,
	}).Handle
}

func (m *Middleware) Handle(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	if authHeader == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := m.authService.ParseToken(ctx.Request.Context(), headerParts[1])

	if err != nil {
		status := http.StatusInternalServerError

		if errors.Is(err, domain.ErrInvalidAccessToken) {
			status = http.StatusUnauthorized
		}

		ctx.AbortWithStatus(status)
		return
	}

	ctx.Set(service.CtxUserKey, user)
}
