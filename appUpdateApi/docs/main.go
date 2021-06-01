package docs

import (
	"fmt"

	_ "github.com/tedmax100/gin-angular/docs"
)

func main() {
	port := 8080
	router := router.SetupRouter(port)
	router.Run(fmt.Sprintf(":%d", port))
}
