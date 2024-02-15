package pkg

type MouseEvent struct {
	X              int
	Y              int
	ClientX        int
	ClientY        int
	PreventDefault func()
	Object         *ObjectDOM
}
