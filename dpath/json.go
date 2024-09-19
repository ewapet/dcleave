package dpath

func JsonGet(subject any, path []any) (any, bool) {
	pathFinder := NewJsonPathFinder()
	return pathFinder.GetPath(path, subject)
}

func JsonPut(subject any, value any, path []any) bool {
	pathFinder := NewJsonPathFinder()
	return pathFinder.SetPath(path, subject, value)
}
