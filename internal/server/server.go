package server

import (
	"article-go/internal/delivery"
	"article-go/internal/service"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run(port string, authService service.AuthService, articleService service.ArticleService) error {
	router := gin.Default()

	router.Use(gin.Recovery(), gin.Logger())

	delivery.RegisterAuthHTTPEndpoints(router, authService)
	delivery.RegisterArticleHTTPEndpoints(router, authService, articleService)

	httpServer := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return httpServer.Shutdown(ctx)
}
