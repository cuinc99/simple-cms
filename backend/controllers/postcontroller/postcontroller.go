package postcontroller

import (
	"net/http"

	"github.com/cuinc99/simple-cms/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var posts []models.Post

	models.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func Show(c *gin.Context) {
	var post models.Post

	id := c.Param("id")

	if err := models.DB.First(&post, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Post not found",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func Create(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	models.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func Update(c *gin.Context) {
	var post models.Post

	id := c.Param("id")

	if err := c.ShouldBindJSON(&post); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if models.DB.Model(&post).Where("id = ?", id).Updates(&post).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Can't update post",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post updated",
		"post":    post,
	})
}

func Delete(c *gin.Context) {
	var post models.Post

	input := map[string]interface{}{"id": "0"}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id := input["id"]

	if models.DB.Delete(&post, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Can't delete post",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post Deleted",
	})
}
