package dpath

import "github.com/ewapet/event-dispatcher/dispatcher"

type PathFinder struct {
	translator Translator
	dispatcher dispatcher.EventDispatcher[EventID, Event]
}

func New(translator Translator) *PathFinder {
	return &PathFinder{
		translator: translator,
		dispatcher: dispatcher.New[EventID, Event](nil),
	}
}

func NewJsonPathFinder() *PathFinder {
	return New(NewJsonTranslator())
}

func (e PathFinder) GetPath(locationPath []any, subject any) (any, bool) {
	if len(locationPath) == 0 {
		return nil, false
	}
	var nextResult = subject
	var ok = true

	nextExtractor := e.translator.(Extractor)

	for i := 0; i < len(locationPath); i++ {
		currentPathEntry := locationPath[i]
		nextResult, ok, nextExtractor = nextExtractor.Extract(currentPathEntry, nextResult)
		if ok == false {
			return nil, false
		}
	}

	return nextResult, ok
}

func (e PathFinder) SetPath(locationPath []any, subject any, value any) bool {
	if len(locationPath) == 0 {
		return false
	}
	var nextResult = subject
	var ok = true

	currentFrame := e.translator.InitialFrame()
	currentEvent := NewEventStart(currentFrame, subject)
	e.dispatcher.Dispatch(EventStart, currentEvent)

	for i := 0; i < len(locationPath); i++ {
		currentPathEntry := locationPath[i]
		currentEvent = currentEvent.AppendPathKey(currentPathEntry)

		if (i + 1) >= len(locationPath) {
			e.dispatcher.Dispatch(EventBeforeIntegrate, currentEvent)
			ok, currentFrame = currentFrame.Integrator.Integrate(currentPathEntry, nextResult, value)
			currentEvent = currentEvent.SetCurrentFrame(currentFrame)
			if !ok {
				return false
			}
		} else {
			e.dispatcher.Dispatch(EventBeforeInitialize, currentEvent)
			ok, nextFrame := currentFrame.Initializer.Initialize(currentPathEntry, nextResult)
			currentEvent = currentEvent.SetCurrentFrame(currentFrame)
			if !ok {
				return false
			}
			e.dispatcher.Dispatch(EventBeforeExtract, currentEvent)
			extractedResult, ok, nextFrame := nextFrame.Extractor.Extract(currentPathEntry, nextResult)
			currentEvent = currentEvent.SetCurrentFrame(nextFrame)
			if !ok {
				panic("failed to extract value after initializing")
			}
			nextResult = extractedResult
			currentFrame = nextFrame
		}
	}
	e.dispatcher.Dispatch(EventEnd, currentEvent)
	return ok
}
