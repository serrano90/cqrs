package main

import (
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/serrano90/converttype"
	"github.com/serrano90/cqrs"
	"github.com/serrano90/cqrs/commandbus"
	"github.com/serrano90/cqrs/examples/simplecqrs/command"
	"github.com/serrano90/cqrs/examples/simplecqrs/commandhandler"
	"github.com/serrano90/cqrs/middleware"
	"github.com/valyala/fasthttp"
)

var bus cqrs.CommandBus

func init() {
	bus = commandbus.NewInMemoryCommandBus()

	//load middleware
	validationMiddleware := middleware.NewValidationMiddleware()

	//load all command handler
	usercommandhandler := cqrs.AddCommandHandlerMiddleware(commandhandler.NewUserCommandHandler(), validationMiddleware)

	//load the command handler in bus
	bus.AddHandler(usercommandhandler, &command.UserCommand{})
}

func userHandler(ctx *fasthttp.RequestCtx) {
	args := ctx.QueryArgs()
	user := command.NewUserCommand(converttype.String(string(args.Peek("name"))))
	_, err := bus.Handler(user)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", userHandler)
	log.Fatal(fasthttp.ListenAndServe(":8000", router.Handler))
}
