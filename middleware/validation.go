package middleware

import (
	"github.com/serrano90/cqrs"
)

//NewValidationMiddleware is a function to create a validation middleware to any command
func NewValidationMiddleware() cqrs.CommandHandlerMiddleware {
	return func(next cqrs.CommandHandler) cqrs.CommandHandler {
		return cqrs.CommandHandlerFunc(func(cmd cqrs.Command) (interface{}, error) {
			if c, ok := cmd.(cqrs.ValidationCommand); ok {
				err := c.Validate()
				if err != nil {
					return nil, err
				}
			}
			return next.Execute(cmd)
		})
	}
}
