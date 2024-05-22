package auth

import (
	"article-go/internal/domain"
	"article-go/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	authService service.AuthService
}

func NewHandler(authService service.AuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}

func (h *Handler) SignUp(ctx *gin.Context) {
	signUpInp := new(domain.SignUpInput)

	if err := ctx.BindJSON(signUpInp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.authService.SignUp(ctx, signUpInp.Username, signUpInp.Email, signUpInp.Password); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) SignIn(ctx *gin.Context) {
	signInInp := new(domain.SignInInput)

	if err := ctx.BindJSON(signInInp); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := h.authService.SignIn(ctx, signInInp.UsernameOrEmail, signInInp.Password)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, domain.SignInResponse{
		Token: token,
	})
}
