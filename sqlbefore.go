package ndrf

type TableNameer interface {
	TableName() string
}

type Beforeer interface {
	GetBefore(requestOj interface{}) (err error)
	GetPageBefore(mapping *Mapping,mappingForm *MappingForm,requestOj interface{}) (err error)
	GetIdBefore(id int) (err error)
	PutIdBefore(id int, requestOj interface{}) (err error)
	PostBefore(requestOj interface{}) (err error)
	DeleteIdBefore(id int) (err error)
}

type Before struct {

}

func (this *Before)  GetBefore(requestOj interface{}) (err error) {
	return nil
}

func (this *Before) GetPageBefore(mapping *Mapping,mappingForm *MappingForm,requestOj interface{}) (err error) {
	return nil
}

func (this *Before) GetIdBefore(id int) (err error)  {
	return nil
}

func (this *Before) PutIdBefore(id int, requestOj interface{}) (err error)  {
	return nil
}

func (this *Before) PostBefore(requestOj interface{}) (err error) {
	return nil
}

func (this *Before) DeleteIdBefore(id int) (err error)  {
	return nil
}






