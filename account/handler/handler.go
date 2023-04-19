package handler

import (
	"net/http"
	"os"

	"github.com/Marif226/words-learner/account/model"
	"github.com/gin-gonic/gin"
)

// Handler struct holds required services
type Handler struct {
	UserService model.UserService
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R *gin.Engine
	UserService model.UserService
}

// NewHandler initializes the handler with http routes and services
func NewHandler(c *Config) {
	// Create a handler
	h := &Handler{
		UserService: c.UserService,
	}

	// Create an account group
	g := c.R.Group(os.Getenv("ACCOUNT_API_URL"))

	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"hello" : "world",
		})
	})

	g.GET("/me", h.Me)
	g.POST("/signup", h.Signup)
	g.POST("/signin", h.Signin)
	g.POST("/signout", h.Signout)
	g.POST("/tokens", h.Tokens)
	g.POST("/image", h.Image)
	g.DELETE("/image", h.DeleteImage)
	g.PUT("/details", h.Details)
}

// Signup handler
func (h *Handler) Signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello" : "it's signup",
	})
}

// Signin handler
func (h *Handler) Signin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"hello" : "it's signin",
	})
}

// Signout handler
func (h *Handler) Signout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"hello" : "it's signout",
	})
}

// Tokens handler
func (h *Handler) Tokens(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"hello" : "it's tokens",
	})
}

// Image handler
func (h *Handler) Image(c *gin.Context) {
c.JSON(http.StatusOK, gin.H {
		"hello" : "it's images",
	})
}

// DeleteImage handler
func (h *Handler) DeleteImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"hello" : "it's deleteImage",
	})
}

// Details handler
func (h *Handler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"hello" : "it's details",
	})
}