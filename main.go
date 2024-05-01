package main

import (
     "email_analyzer/db"
     "email_analyzer/routes"
     "github.com/gin-gonic/gin"
)

func main() {
    db.InitDB() 

    router := gin.Default() 
    routes.RegisterRoutes(router) // Pass router to the registration function

    router.Run(":8080") // Example port, adjust as needed
}

