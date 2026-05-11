package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/olawuwo-abideen/go-crud/initializers"
	"github.com/olawuwo-abideen/go-crud/models"
)

type CreatePostInput struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type UpdatePostInput struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// PostsCreate godoc
// @Summary Create a new post
// @Description Create a post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body CreatePostInput true "Post Data"
// @Success 201 {object} models.Post
// @Router /posts [post]
func PostsCreate(c *gin.Context) {

	var body CreatePostInput

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	if err := initializers.DB.Create(&post).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error":   "Failed to create post",
		})
		return
	}

	c.JSON(201, gin.H{
		"success": true,
		"data":    post,
	})
}

// PostsIndex godoc
// @Summary Get all posts
// @Description Retrieve all posts
// @Tags posts
// @Produce json
// @Success 200 {array} models.Post
// @Router /posts [get]
func PostsIndex(c *gin.Context) {
	var posts []models.Post

	if err := initializers.DB.Find(&posts).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error":   "Failed to fetch posts",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    posts,
	})
}

// PostsShow godoc
// @Summary Get a single post
// @Description Retrieve a post by ID
// @Tags posts
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} models.Post
// @Router /posts/{id} [get]
func PostsShow(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := initializers.DB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{
			"success": false,
			"error":   "Post not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    post,
	})
}

// PostsUpdate godoc
// @Summary Update a post
// @Description Update a post by ID
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param post body UpdatePostInput true "Updated Post"
// @Success 200 {object} models.Post
// @Router /posts/{id} [put]
func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body UpdatePostInput

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	var post models.Post
	if err := initializers.DB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{
			"success": false,
			"error":   "Post not found",
		})
		return
	}

	if err := initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	}).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error":   "Failed to update post",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    post,
	})
}

// PostsDelete godoc
// @Summary Delete a post
// @Description Delete a post by ID
// @Tags posts
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {string} string "Post deleted successfully"
// @Router /posts/{id} [delete]
func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	result := initializers.DB.First(&post, "id = ?", id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"success": false,
			"error":   "Post not found",
		})
		return
	}

	if err := initializers.DB.Delete(&post).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error":   "Failed to delete post",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Post deleted successfully",
	})
}
