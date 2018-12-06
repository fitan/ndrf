package ndrf


//type ResponseChecker interface {
//	Ok(data interface{}, c *gin.Context)
//	Err(err interface{}, c *gin.Context)
//}
//type ResponseCheck struct {
//}
//
//func (this *ResponseCheck) Ok(data interface{}, c *gin.Context) {
//	c.JSON(http.StatusOK, data)
//	return
//}
//
//func (this *ResponseCheck) Err(data interface{}, c *gin.Context) {
//	c.JSON(http.StatusOK, data)
//	return
//}

type ResponseWorker interface {
	OkWork(data interface{}) interface{}
	ErrWork(err error) interface{}
}
type ResponseWork struct {
}

type workOj struct {
	Status string
	Data interface{}
	Msg string
}

func (this *ResponseWork) OkWork(data interface{}) interface{} {
	return &workOj{"ok", data, ""}
}

func (this *ResponseWork) ErrWork(err error) interface{} {
	return &workOj{"err", "", err.Error()}
}




