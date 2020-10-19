package someobject

import "errors"

type SomeObject struct {
	SomeString string
	SomeInt int
}

func (s *SomeObject) DoSomething() (error, *string){
	if s.SomeInt > 10 {
		return errors.New("error"), nil
	}
	return nil, &s.SomeString
}