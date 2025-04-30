package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"coffeeshop/coffee"
)

var log_path = "logs/logs.txt"
var sugar *zap.SugaredLogger

func createFile() {
	var _, err = os.Stat(log_path)
	if os.IsNotExist(err) {
		err = os.MkdirAll("logs", os.ModePerm)

		if err != nil {
			return
		}

		var file, err = os.Create(log_path)
		if err != nil {
			return
		}

		defer file.Close()
	}
}

func NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{log_path}
	return config.Build()
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to CoffeeShop!",
	})
}

func getCoffeeList(c *gin.Context) {
	coffees, err := coffee.GetCoffees()
	if err != nil {
		sugar.Error("Error getting the coffee list: ", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"coffees": coffees,
	})
}

func getCoffeeListApi(c *gin.Context) {
	coffees, err := coffee.GetCoffees()
	if err != nil {
		sugar.Error("Error getting the coffee list: ", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"coffees": coffees,
	})
}

func homePage(c *gin.Context) {
	coffees, err := coffee.GetCoffees()
	if err != nil {
		sugar.Error("Error getting the coffee list: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting the coffee list",
		})
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"coffees": coffees.List,
	})
}

func setLogOutput() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func init() {
	createFile()
	logger, err := NewLogger()
	if err != nil {
		log.Fatal(err)
	} else {
		sugar = logger.Sugar()
	}

}

func main() {
	var PORT = os.Getenv("PORT")

	// using gin
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/ping", ping)
	r.GET("/coffeeListApi", getCoffeeListApi)
	r.GET("/home", homePage)

	sugar.Info("Starting server at port %s...", PORT)
	//r.Run(fmt.Sprintf(":%s", PORT))
	r.Run(":8080")
}
