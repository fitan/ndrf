package ndrf

type Ojser interface {
	GetModel() interface{}
	GetModels() interface{}
	GetRequestOj() interface{}
}



type MappingForm struct {
	Page int `form:"page"`
	Limit int `form:"limit"`
}