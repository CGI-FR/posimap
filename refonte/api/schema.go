package api

type Schema interface {
	Build() Record
}
