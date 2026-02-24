package main

import (
	// To retrieve json
	"fmt" // to print stuff
	"log"
	"strconv"

	"github.com/JorgeJola/indnrate-go/internal/database"
	"github.com/gin-gonic/gin" // Web Framework
)

func main() {
	database.Connect()
	
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})


	router.GET("/simresults", func(c *gin.Context) {
    cellStr := c.Query("cell")
    if cellStr == "" {
        c.JSON(400, gin.H{"error": "cell parameter is required"})
        return
    }

    cellID, err := strconv.Atoi(cellStr)
    if err != nil {
        c.JSON(400, gin.H{"error": "cell must be an integer"})
        return
    }

    sims, err := database.Query(cellID)
    if err != nil {
        c.JSON(500, gin.H{"error": "database query failed"})
        return
    }

    c.JSON(200, sims)

})
	fmt.Println("Server running on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}



