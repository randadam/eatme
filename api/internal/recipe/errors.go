package recipe

import "errors"

var (
	ErrRecipeNotFound           = errors.New("recipe not found")
	ErrRecipeVersionNotFound    = errors.New("recipe version not found")
	ErrSuggestionThreadNotFound = errors.New("suggestion thread not found")
	ErrSuggestionNotFound       = errors.New("suggestion not found")
)
