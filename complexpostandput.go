package ndrf

type ComplexPoster interface {
	ComplexPostOj() interface{}
	ComplexPostMapping(requestOj interface{}, model interface{}) error
}

type ComplexPuter interface {
	ComplexPutOj() interface{}
	ComplexPutMapping(requestOj interface{}, model interface{}) error
}


type complexPost struct {

}

func (this *complexPost) ComplexPostOj() interface{}  {
	return  nil

}

func (this *complexPost) ComplexPostMapping(requestOj interface{}, model interface{}) error {
	return nil
}

type complexPut struct {

}

func (this *complexPut) ComplexPutOj() interface{} {
	return nil
}

func (this *complexPut) ComplexPutMapping(requesOj interface{}, model interface{}) error {
	return nil
}
