package main



//  original 
import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"strconv"
	"fmt"
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

	router.GET("/inicio", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	
	
	router.GET("v1/multiplica/:numero1/:numero2", getMultiplicaByID)
	
	router.Run(":" + port)


}
//termin ortiginall

 



// multiplica represents data about multiplicacion.
type multiplica struct {
    Status     string  `json:"status"`
    Resultado  string  `json:"resultado"`

}

var resultados = []multiplica{
    {Status: "ok", Resultado: "Blue Train"},

}


// getMultiplicaByID responds with the stauts and the result as JSON.
func getMultiplicaByID(c *gin.Context) {
	 elemento1 := c.Param("numero1")
	 elemento2 := c.Param("numero2")
	var s1final float64 = 0
	var s2final float64 = 0
	
	 
	if s1, err := strconv.ParseFloat(elemento1, 64); err == nil {
             fmt.Println(s1) // 3.1415927410125732
		s1final =s1;
	}else{
		resultados[0].Status ="NOK";
	        resultados[0].Resultado = "first parameter is expected numeric";
	        c.IndentedJSON(http.StatusOK, resultados)
		return
	}
       if s2, err := strconv.ParseFloat(elemento2, 64); err == nil {
         fmt.Println(s2) // 3.14159265
	       s2final =s2
	}else{
         	resultados[0].Status ="NOK";
	        resultados[0].Resultado = "second parameter is expected numeric";
	        c.IndentedJSON(http.StatusOK, resultados)
		return
	}
	
	resultado := s1final* s2final;
	 fmt.Println(resultado) 
	sresultado := fmt.Sprintf("%f", resultado)
	
	resultados[0].Resultado = sresultado
	d:=resultados[0].Resultado 
	fmt.Println(d+"si")

//	response:= json.NewEncoder(w).Encode(map[string]string{"status": "OK"})	
	
        c.IndentedJSON(http.StatusOK, resultados)
}
