package cqrs

//CommandHandlerMiddleware is a function that middleware
type CommandHandlerMiddleware func(CommandHandler) CommandHandler

//AddCommandHandlerMiddleware is a function that add all middleware to command handler
func AddCommandHandlerMiddleware(h CommandHandler, middlewares ...CommandHandlerMiddleware) CommandHandler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}
