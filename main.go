package main

import (
    "backend/config"
    "backend/routes"
)

func main() {
    config.ConnectDatabase()
    r := routes.SetupRouter()
    r.Run(":8080")
}
