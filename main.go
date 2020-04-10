package main

import (
	"github.com/gin-gonic/gin" 
	"github.com/subosito/gotenv" 
	"github.com/ariespirnando/covid19/controller" 
	"log"
	"os"
	"fmt"
)

func init()  {
	err := gotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
    } 
    if(os.Getenv("GIN_MODE")=="release"){
        gin.SetMode(gin.ReleaseMode)
	} 
}

func main(){   
	fmt.Printf("Welcome API \n") 
    router := gin.Default()  
	url := router.Group("/covid19/") 
	{	
		url.GET("/allprovinsi", controller.Allprovinsi)  
		url.GET("/kasusperprovinsi", controller.Kasusperprovinsi)   
		url.GET("/kasusprovinsi", controller.Kasusprovinsi)
		url.GET("/cronjob", controller.Cronjob)   
	}
    router.Run(os.Getenv("ROUTER"))
}