package controllers

import (
    "net/http"
    "backend/config"
    "backend/models"
    "github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := config.DB.Exec("INSERT INTO orders (user_id, product_id, quantity) VALUES (?, ?, ?)",
        order.UserID, order.ProductID, order.Quantity)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat pesanan"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Pesanan berhasil dibuat!"})
}


func GetOrders(c *gin.Context) {
    rows, err := config.DB.Query(`
        SELECT 
            o.id, o.user_id, u.username, 
            o.product_id, p.name AS product_name, p.price,
            o.quantity, o.order_date
        FROM online_store_orders o
        JOIN online_store_users u ON o.user_id = u.id
        JOIN online_store_products p ON o.product_id = p.id
        ORDER BY o.order_date DESC
    `)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data pesanan"})
        return
    }
    defer rows.Close()

    var orders []gin.H
    for rows.Next() {
        var order struct {
            ID          int     `json:"id"`
            UserID      int     `json:"user_id"`
            Username    string  `json:"username"`
            ProductID   int     `json:"product_id"`
            ProductName string  `json:"product_name"`
            Price       float64 `json:"price"`
            Quantity    int     `json:"quantity"`
            OrderDate   string  `json:"order_date"`
        }

        err := rows.Scan(&order.ID, &order.UserID, &order.Username, &order.ProductID, &order.ProductName, &order.Price, &order.Quantity, &order.OrderDate)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses data pesanan"})
            return
        }

        orders = append(orders, gin.H{
            "id":          order.ID,
            "user_id":     order.UserID,
            "username":    order.Username,
            "product_id":  order.ProductID,
            "product_name": order.ProductName,
            "price":       order.Price,
            "quantity":    order.Quantity,
            "order_date":  order.OrderDate,
        })
    }

    c.JSON(http.StatusOK, gin.H{"orders": orders})
}

