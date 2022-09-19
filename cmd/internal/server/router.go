package server

import (
	"net/http"
	"blog.com/packages/cmd/internal/store"

	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
	r := gin.Default()

	//Enables automatic redirection if the current route can't handler
	//for the path with (without) the trailing slash e
	r.RedirectTrailingSlash = true

	api := r.Group("/api")
	{
		api.POST("/signup", gin.Bind(store.User{}) ,signUp)
		api.POST("/signin",  gin.Bind(store.User{}),signIn)

	}

	authorized := api.Group("/")
	authorized.Use(authorization)
	{
		authorized.GET("/posts",indexPosts)
		authorized.POST("/posts",createPost)
		authorized.PUT("/posts",updatePost)
		authorized.DELETE("/posts/:id",deletePost)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{})
	})

	return r

}
