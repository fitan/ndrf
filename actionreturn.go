package ndrf


type ActionReturner interface {
	GetR(models interface{}) (error, interface{})
	GetPageR(models interface{}) (error, interface{})
	GetIdR(model interface{}) (error, interface{})
	PutR(model interface{}) (error, interface{})
	DeleteR(model interface{}) (error, interface{})
	PostR(model interface{}) (error, interface{})
}

type ActionReturn struct {
}


func (this *ActionReturn)  GetR(models  interface{}) (error, interface{}) {
	return nil, models
}

func (this *ActionReturn) GetPageR(models interface{}) (error, interface{}) {
	return nil, models
}

func (this *ActionReturn) GetIdR(model interface{}) (error, interface{}) {
	return nil, model
}

func (this *ActionReturn) PutR(model interface{}) (error, interface{}) {
	return nil, model
}

func (this *ActionReturn) DeleteR(model interface{}) (error, interface{}) {
	return nil, model
}

func (this *ActionReturn) PostR(model interface{}) (error, interface{}) {
	return nil, model
}
