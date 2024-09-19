package dpath

type JsonTranslator struct {
	frame TranslationFrame
}

func NewJsonTranslator() *JsonTranslator {
	translator := &JsonTranslator{}
	translator.frame = TranslationFrame{
		Extractor:   translator,
		Initializer: translator,
		Integrator:  translator,
	}
	return translator
}

func (j JsonTranslator) InitialFrame() TranslationFrame {
	return j.frame
}

func (j JsonTranslator) Extract(key any, subject any) (any, bool, TranslationFrame) {
	jsonSliceData, isJsonSlice := subject.([]any)
	jsonMapData, isJsonMap := subject.(map[string]any)
	if isJsonSlice {
		if keyInt, ok := key.(int); !ok {
			return nil, false, j.frame
		} else if keyInt >= len(jsonSliceData) || keyInt < 0 {
			return nil, false, j.frame
		} else {
			return jsonSliceData[keyInt], true, j.frame
		}
	} else if isJsonMap {
		if keyString, ok := key.(string); !ok {
			return nil, false, j.frame
		} else if result, ok := jsonMapData[keyString]; !ok {
			return nil, false, j.frame
		} else {
			return result, true, j.frame
		}
	} else {
		return nil, false, j.frame
	}
}

func (j JsonTranslator) Initialize(property any, subject any) (bool, TranslationFrame) {
	_, isJsonSlice := subject.([]any)
	mapSubject, isJsonMap := subject.(map[string]any)
	if isJsonSlice {
		keyPosition, ok := property.(int)
		if !ok || keyPosition < 0 {
			return false, j.frame
		}
		return false, j.frame
	} else if isJsonMap {
		keyString, ok := property.(string)
		if !ok {
			return false, j.frame
		}
		_, isMapTarget := mapSubject[keyString].(map[string]any)
		if !isMapTarget {
			mapSubject[keyString] = make(map[string]any)
		}
		return true, j.frame
	} else {
		return false, j.frame
	}
}

func (j JsonTranslator) Integrate(property any, subject any, value any) (bool, TranslationFrame) {
	sliceSubject, isJsonSlice := subject.([]any)
	mapSubject, isJsonMap := subject.(map[string]any)
	if isJsonSlice {
		if keyInt, ok := property.(int); !ok {
			return false, j.frame
		} else if keyInt >= len(sliceSubject) || keyInt < 0 {
			return false, j.frame
		} else {
			sliceSubject[keyInt] = value
			return true, j.frame
		}
	} else if isJsonMap {
		if keyString, ok := property.(string); !ok {
			return false, j.frame
		} else {
			mapSubject[keyString] = value
			return true, j.frame
		}
	} else {
		return false, j.frame
	}
}
