package challenge02

type Context interface {
}

type emptyCtx int

type backgroundCtx struct{ emptyCtx }

func (backgroundCtx) String() string {
	return "context.Background"
}

type todoCtx struct{ emptyCtx }

func (todoCtx) String() string {
	return "context.TODO"
}

func Background() Context {
	return backgroundCtx{}
}

func TODO() Context {
	return todoCtx{}
}
