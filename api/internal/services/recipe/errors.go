package recipe

import "errors"

var (
	ErrRecipeNotFound           = errors.New("recipe not found")
	ErrSuggestionThreadNotFound = errors.New("suggestion thread not found")
	ErrSuggestionNotFound       = errors.New("suggestion not found")
)
