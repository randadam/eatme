package thread

import "errors"

var (
	ErrThreadNotFound                       = errors.New("thread not found")
	ErrThreadNotAssociatedWithRecipeVersion = errors.New("thread not associated with recipe version")
	ErrInvalidThreadEventType               = errors.New("invalid thread event type")
	ErrInvalidThreadEventPayload            = errors.New("invalid thread event payload")
	ErrSuggestionNotFound                   = errors.New("suggestion not found")
)
