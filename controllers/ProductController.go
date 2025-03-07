package controllers

import (
    "net/http"
    "backend/config"
    "backend/models"
    "github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := config.DB.Exec("INSERT INTO products (name, description, price, image) VALUES (?, ?, ?, ?)",
        product.Name, product.Description, product.Price, product.Image)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan produk"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil ditambahkan!"})
}
