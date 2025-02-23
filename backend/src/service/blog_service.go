package service

import (
    "seaventures/src/models"
    "seaventures/src/repository"
)

func CreateBlog(blog *models.Blog) error {
    return repository.CreateBlog(blog)
}

func GetAllBlogs() ([]models.Blog, error) {
    return repository.GetAllBlogs()
}

func GetBlogByID(id string) (*models.Blog, error) {
    return repository.GetBlogByID(id)
}

func UpdateBlog(id string, blog *models.Blog) error {
    return repository.UpdateBlog(id, blog)
}

func DeleteBlog(id string) error {
    return repository.DeleteBlog(id)
}