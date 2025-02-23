package repository

import (
    "seaventures/src/config"
    "seaventures/src/models"
)

func CreateBlog(blog *models.Blog) error {
    return config.DB.Create(blog).Error
}

func GetAllBlogs() ([]models.Blog, error) {
    var blogs []models.Blog
    err := config.DB.Find(&blogs).Error
    return blogs, err
}

func GetBlogByID(id string) (*models.Blog, error) {
    var blog models.Blog
    err := config.DB.Where("blog_id = ?", id).First(&blog).Error
    return &blog, err
}

func UpdateBlog(id string, updatedBlog *models.Blog) error {
    var blog models.Blog
    err := config.DB.Where("blog_id = ?", id).First(&blog).Error
    if err != nil {
        return err
    }
    return config.DB.Model(&blog).Updates(updatedBlog).Error
}

func DeleteBlog(id string) error {
    return config.DB.Where("blog_id = ?", id).Delete(&models.Blog{}).Error
}