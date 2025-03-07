package controllers

import (
    "net/http"
    "backend/config"
    "backend/models"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    _, err := config.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, hashedPassword)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal register"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Registrasi berhasil!"})
}
