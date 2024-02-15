package pkg

// Represents a MouseEvent object in Typescript and JS (It is very incomplete)
type MouseEvent struct {
	X              int        // X
	Y              int        // Y
	ClientX        int        // ClientX
	ClientY        int        // ClientY
	PreventDefault func()     // PreventDefault()
	Object         *ObjectDOM // Object
}
