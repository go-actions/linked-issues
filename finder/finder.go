package finder

type IssueFinder interface {
	Find() ([]string, error)
}

func NewIssueFinder(f IssueFinder) IssueFinder {
	return f
}
