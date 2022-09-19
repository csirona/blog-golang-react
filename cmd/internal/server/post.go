package server

import(
	"net/http"
	"strconv"
	"time"
	"blog.com/packages/cmd/internal/store";
	"github.com/gin-gonic/gin"

)

func createPost(c *gin.Context) {
	post := new(store.Post)
	if err := c.Bind(post); err!= nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	user, err := currentUser(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	if err := store.AddPost(user,post); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "Post created successfully.",
		"data":post,
	})
}

func indexPosts(c *gin.Context) {
	user, err := currentUser(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	if err := store.FetchUserPosts(user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":"Posts fetched successfully",
		"data": user.Posts,
	})
}

func updatePost(c 	*gin.Context) {
	jsonPost := new(store.Post)
	if err := c.Bind(jsonPost); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	user, err := currentUser(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	dbPost, err := store.FetchPost(jsonPost.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	if user.ID != dbPost.UserID {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error":"Not authorized"})
		return
	}
	jsonPost.ModifiedAt = time.Now()
	if err := store.UpdatePost(jsonPost); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":"Post updated successufully.",
		"data": jsonPost,
	})
}

func deletePost(c 	*gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":"Not valid ID."})
		return
	}
	user, err := currentUser(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	post, err := store.FetchPost(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	if user.ID != post.UserID {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error":"Not authorized"})
		return
	}
	if err := store.DeletePost(post); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":"Post deleted successufully.",
	})
}
