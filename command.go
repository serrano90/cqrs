package cqrs

//Command
type Command interface {
	//TypeOf return a string descriptor of the command name
	TypeOf() string
}

//ValidationCommand
type ValidationCommand interface {
	//Validate return a error is command defined validation is incorrect
	Validate() error
}
