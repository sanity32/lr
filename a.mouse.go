package lr

type MouseMoveArgs struct {
	X      int
	Y      int
	Smooth bool
}

type MouseClickArgs struct {
	Button string
	Double bool
}
