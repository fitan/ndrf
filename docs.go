package ndrf

import (
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/gin-gonic/gin"
	"net/http"
	"gopkg.in/yaml.v2"
	"github.com/swaggo/swag"
)

func OpenDoc(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/api/doc", func(c *gin.Context) {
		s, err := yaml.Marshal(Swagger)
		if err != nil {
			panic(err)
		}
		c.String(http.StatusOK, string(s))
	})
}


type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	doc, err := yaml.Marshal(Swagger)

	if err != nil {
		panic(err)
	}
	return string(doc)
}

func init() {
	swag.Register(swag.Name, &s{})
}
