package main

type FreeDayError struct {
	Query string
}

func (e FreeDayError) Error() string {
	return e.Query
}
