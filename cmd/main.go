package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Acepta cualquier método y ruta
	r.Any("/*path", func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		fmt.Println("📩 Nuevo request recibido:")
		fmt.Println("Método:", c.Request.Method)
		fmt.Println("Ruta:", c.Request.URL.Path)
		fmt.Println("Query:", c.Request.URL.RawQuery)
		fmt.Println("Headers:", c.Request.Header)
		fmt.Println("Body:", string(body))
		fmt.Println("------------------------")

		c.String(http.StatusOK, "OK")
	})

	fmt.Println("🚀 API escuchando en http://localhost:9001")
	r.Run(":9001")
}
