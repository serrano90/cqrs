package command

import (
	"errors"
	"reflect"

	"github.com/serrano90/cqrs"
)

func NewUserCommand(name *string) cqrs.Command {
	return &UserCommand{
		name: name,
	}
}

type UserCommand struct {
	name *string `valid:"required"`
}

func (this *UserCommand) TypeOf() string {
	return reflect.TypeOf(this).String()
}

func (this *UserCommand) Validate() error {
	if this.name == nil || *this.name == "" {
		return errors.New("Bad Request Body")
	}
	// _, err := govalidator.ValidateStruct(this)
	// if err != nil {
	// 	return err
	// }
	return nil
}
