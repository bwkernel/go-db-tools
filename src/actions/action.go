package actions

type Action interface{
	Handle(args []string)(error)
}
