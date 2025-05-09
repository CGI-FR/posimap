package api

type Predicate func(root Record) (bool, error)
