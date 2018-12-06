package ndrf

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
)

func GetId(c *gin.Context) (idInt int, err error) {
	id := c.Param("id")
	idInt, err = strconv.Atoi(id)
	return
}

func CreateView(route *gin.RouterGroup, db *gorm.DB) *View {
	return &View{db, route}
}

type Serializerser interface {
	Ojser
	ActionReturner
	ActionSqler
	ResponseWorker
	Beforeer
}

type View struct {
	db *gorm.DB
	route *gin.RouterGroup
}

func (this *View) start() {

}

func (this *View) Inset(serializers Serializerser, mapping *Mapping, methods []string, path string)  {
	serializers.UseDb(this.db)
	docPath := this.route.BasePath() + path
	for _,value := range methods {
		switch value {
		case "GET":
			if mapping.PageOpen == true {

				GetPageAddPaths(docPath, serializers.GetRequestOj(), serializers.GetModels())

				this.route.GET(path, func(c *gin.Context) {
					requestOj := serializers.GetRequestOj()
					err := c.Bind(requestOj)
					if err != nil {
						c.JSON(http.StatusOK, serializers.ErrWork(err))
						return
					}
					mappingForm := new(MappingForm)
					err = c.Bind(mappingForm)
					if err != nil {
						c.JSON(http.StatusOK, serializers.ErrWork(err))
						return
					}
					err = serializers.GetPageBefore(mapping, mappingForm, requestOj)
					if err != nil {
						c.JSON(http.StatusOK, serializers.ErrWork(err))
						return
					}
					models := serializers.GetModels()
					err = serializers.GetPageSql(mapping, mappingForm, requestOj, models)
					if err != nil {
						c.JSON(http.StatusOK, serializers.ErrWork(err))
						return
					}
					err, data := serializers.GetPageR(models)
					if err != nil {
						c.JSON(http.StatusOK, serializers.ErrWork(err))
						return
					}
					c.JSON(http.StatusOK, serializers.OkWork(data))
				})
				continue
			}

			GetAddPaths(docPath, serializers.GetRequestOj(), serializers.GetModels())

			this.route.GET(path, func(c *gin.Context) {
				requestOj := serializers.GetRequestOj()
				models := serializers.GetModels()
				err := c.Bind(requestOj)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				err = serializers.GetBefore(requestOj)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				err = serializers.GetSql(requestOj, models)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				err, data := serializers.GetR(models)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				c.JSON(http.StatusOK, serializers.OkWork(data))
			})
		case "GETID":

			GetIdPaths(docPath, serializers.GetModel())

			this.route.GET(path + "/:id", func(c *gin.Context) {
				id, err := GetId(c)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				err = serializers.GetIdBefore(id)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				model := serializers.GetModel()
				err = serializers.GetIdSql(id, model)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				err, data := serializers.GetIdR(model)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				c.JSON(http.StatusOK, serializers.OkWork(data))
			})
		case "PUT":

			if _, ok := serializers.(ComplexPuter); ok {

				PutAddPaths(docPath, serializers.(ComplexPuter).ComplexPutOj(), serializers.GetModel())

				this.route.PUT(path + "/:id", func(c *gin.Context) {
					id, err := GetId(c)
					if err != nil {
						c.JSON(http.StatusOK, serializers.ErrWork(err))
						return
					}
					requestOj := serializers.(ComplexPuter).ComplexPutOj()
					err = c.Bind(requestOj)
					err = serializers.PutIdBefore(id, requestOj)
					if err != nil {
						c.JSON(http.StatusOK, serializers.ErrWork(err))
						return
					}
					model := serializers.GetModel()
					err = serializers.(ComplexPuter).ComplexPutMapping(requestOj, model)
					if err != nil {
						c.JSON(http.StatusOK, serializers.ErrWork(err))
						return
					}
					err = serializers.PutIdSql(id,requestOj, model)
					if err != nil {
						c.JSON(http.StatusOK, serializers.ErrWork(err))
						return
					}
					err, data := serializers.PutR(model)
					if err != nil {
						c.JSON(http.StatusOK, serializers.ErrWork(err))
						return
					}
					c.JSON(http.StatusOK, serializers.OkWork(data))
				})
				continue
			}

			PutAddPaths(docPath, serializers.GetRequestOj(), serializers.GetModel())
			this.route.PUT(path + "/:id", func(c *gin.Context) {
				id, err := GetId(c)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				requestOj := serializers.GetRequestOj()
				err = c.Bind(requestOj)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				err = serializers.PutIdBefore(id, requestOj)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				model := serializers.GetModel()
				err = serializers.PutIdSql(id, requestOj, model)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				err, data := serializers.PutR(model)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				c.JSON(http.StatusOK, serializers.OkWork(data))
			})
		case "POST":

			if _,ok := serializers.(ComplexPoster); ok {

				PostAddPaths(docPath, serializers.(ComplexPoster).ComplexPostOj(), serializers.GetModel())

				this.route.POST(path, func(c *gin.Context) {
					requestOj := serializers.(ComplexPoster).ComplexPostOj()
					err := c.Bind(requestOj)
					if err != nil {
						c.JSON(http.StatusOK, serializers.ErrWork(err))
						return
					}
					err = serializers.PostBefore(requestOj)
					if err != nil {
						c.JSON(http.StatusOK, serializers.ErrWork(err))
						return
					}
					model := serializers.GetModel()
					err = serializers.(ComplexPoster).ComplexPostMapping(requestOj, model)
					if err != nil {
						c.JSON(http.StatusOK, serializers.ErrWork(err))
						return
					}
					err = serializers.PostSql(model)
					if err != nil {
						c.JSON(http.StatusOK, serializers.ErrWork(err))
						return
					}
					err, data := serializers.PostR(model)
					if err != nil {
						c.JSON(http.StatusOK, serializers.ErrWork(err))
						return
					}
					c.JSON(http.StatusOK, serializers.OkWork(data))
				})
				continue
			}
			PostAddPaths(docPath, serializers.GetRequestOj(), serializers.GetModel())
			this.route.POST(path, func(c *gin.Context) {
				requestOj := serializers.GetRequestOj()
				err := c.Bind(requestOj)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				err = serializers.PostBefore(requestOj)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				err = serializers.PostSql(requestOj)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				err, data := serializers.PostR(requestOj)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				c.JSON(http.StatusOK, serializers.OkWork(data))
			})
		case "DELETE":

			DeleteIdPaths(docPath, serializers.GetModel())

			this.route.DELETE(path+"/:id", func(c *gin.Context) {
				id, err := GetId(c)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				err = serializers.DeleteIdBefore(id)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				model := serializers.GetModel()
				err = serializers.DeleteIdSql(id,model)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				err, data := serializers.DeleteR(model)
				if err != nil {
					c.JSON(http.StatusOK, serializers.ErrWork(err))
					return
				}
				c.JSON(http.StatusOK, serializers.OkWork(data))
			})
		}
	}

}

