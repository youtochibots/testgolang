package main



//  original 
import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	
//	router.GET("/v1/multiplica", netHandle(handleDBGettokenizedcards, nil))
	router.Run(":" + port)
/*
	var r = net.GetRouter()
	//route for test
	    fmt.Print("cz  init net_v1")

 	r.Handle("/v3/fetchtokenizedcards", netHandle(handleDBGettokenizedcards, nil)).Methods("GET")   //logicbusiness.go
  */   

}
//termin ortiginall


func handleDBGettokenizedcards(w http.ResponseWriter, r *http.Request) {

	fmt.Print("uno");					
}
