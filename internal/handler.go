package internal

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("internal/template/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.POST("/", h.home)
	return router
}

func (h *Handler) home(c *gin.Context) {
	uid := c.PostForm("id")
	order, ok := h.repo.CacheRepo.GetByUid(uid)
	if !ok {
		c.HTML(http.StatusBadRequest, "error.html", nil)
		return
	}
	c.HTML(http.StatusOK, "order.html", order)
}
