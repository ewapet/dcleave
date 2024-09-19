package dpath

type EventID uint

const (
	EventStart EventID = iota
	EventBeforeInitialize
	EventBeforeExtract
	EventBeforeIntegrate
	EventEnd
)

type Event struct {
	currentFrame TranslationFrame
	currentPath  []any
	initialFrame TranslationFrame
	subject      any
}

func NewEventStart(frame TranslationFrame, subject any) Event {
	return Event{
		currentFrame: frame,
		currentPath:  make([]any, 0),
		initialFrame: frame,
		subject:      subject,
	}
}

func (e Event) AppendPathKey(key any) Event {
	e.currentPath = append(e.currentPath, key)
	return e
}

func (e Event) CurrentFrame() TranslationFrame {
	return e.currentFrame
}

func (e Event) CurrentPath() []any {
	return e.currentPath
}

func (e Event) InitialFrame() TranslationFrame {
	return e.initialFrame
}

func (e Event) SetCurrentFrame(nextFrame TranslationFrame) Event {
	e.currentFrame = nextFrame
	return e
}

func (e Event) Subject() any {
	return e.subject
}
