package main

import (
	// To retrieve json
	"fmt" // to print stuff
	"log"
	"strconv"

	"github.com/JorgeJola/indnratebackend/internal/database"
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
	// Optional parameters (Nitrogen and grain price)
	nitroPriceStr := c.Query("nitro_price")
	
	var nitroPrice float64

	if nitroPriceStr != "" {
		nitroPrice,err = strconv.ParseFloat(nitroPriceStr,64)
		if err != nil{
			c.JSON(400, gin.H{"error" : "Nitrogen price should be a number (Float)"})
			return
		}
	} else {
		nitroPrice = 0.4
	}

	grainPriceStr := c.Query("grain_price")

	var grainPrice float64

	if grainPriceStr != "" {
		grainPrice,err = strconv.ParseFloat(grainPriceStr,64)
		if err !=nil{
			c.JSON(400, gin.H{"error":"Grain price should be a number (Float)"})
			return
		}
	} else{
		grainPrice = 4
	}
	// Quering database
    sims, err := database.Query(cellID, nitroPrice, grainPrice)
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



