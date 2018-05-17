package api

import "github.com/asaskevich/govalidator"

type SubscribeForm struct {
	Email string `json:"email" valid:"email,required"`
	Path  string `json:"path"`
}

func (f *SubscribeForm) validate() error {
	_, err := govalidator.ValidateStruct(f)
	return err
}
