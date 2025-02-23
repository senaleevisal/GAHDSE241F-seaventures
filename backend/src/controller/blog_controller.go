package controller

import (
    "net/http"

    "seaventures/src/models"
    "seaventures/src/service"

    "github.com/gin-gonic/gin"
)

func CreateBlog(c *gin.Context) {
    var blog models.Blog
    if err := c.ShouldBindJSON(&blog); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := service.CreateBlog(&blog); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Blog created successfully"})
}

func GetBlogs(c *gin.Context) {
    blogs, err := service.GetAllBlogs()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, blogs)
}

func GetBlogByID(c *gin.Context) {
    id := c.Param("id")
    blog, err := service.GetBlogByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, blog)
}

func UpdateBlog(c *gin.Context) {
    id := c.Param("id")
    var blog models.Blog
    if err := c.ShouldBindJSON(&blog); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := service.UpdateBlog(id, &blog); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Blog updated successfully"})
}

func DeleteBlog(c *gin.Context) {
    id := c.Param("id")
    if err := service.DeleteBlog(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}