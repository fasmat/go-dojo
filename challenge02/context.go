package challenge02

type Context interface {
}

type emptyCtx struct{}

func (e *emptyCtx) String() string {
	switch e {
	case background:
		return "context.Background"
	case todo:
		return "context.TODO"
	}
	return "unknown empty Context"
}

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

func TODO() Context {
	return nil
}

func Background() Context {
	return nil
}
