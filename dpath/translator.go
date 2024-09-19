package dpath

type Translator interface {
	InitialFrame() TranslationFrame
}

type TranslationFrame struct {
	Extractor
	Initializer
	Integrator
}

type Extractor interface {
	Extract(location any, subject any) (any, bool, TranslationFrame)
}

type Initializer interface {
	Initialize(location any, subject any) (bool, TranslationFrame)
}

type Integrator interface {
	Integrate(location any, subject any, value any) (bool, TranslationFrame)
}
