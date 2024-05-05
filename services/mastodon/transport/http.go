package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tomef96/coop/mastodon/domain"
)

func Router(postsService domain.PostsService) *gin.Engine {
	r := gin.Default()

	r.GET("/posts", func(ctx *gin.Context) {
		posts, err := postsService.HandleListPosts(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, posts)
		}
	})

	r.POST("/posts", func(ctx *gin.Context) {
		var post domain.NewPost
		if err := ctx.ShouldBindJSON(&post); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !post.Valid() {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "title and body must be defined"})
			return
		}

		err := postsService.HandleNewPost(ctx, post)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "an unexpected error occured",
				"error":   err.Error(),
			})
		} else {
			ctx.JSON(200, gin.H{
				"message": "post successfully handled",
			})
		}
	})

	return r
}
