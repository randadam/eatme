package thread

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/ajohnston1219/eatme/api/internal/models"
)

func TestReduceThreadEvents(t *testing.T) {
	threadID := "test_thread_id"
	events := createThreadEvents(t,
		withEvent(models.ThreadEventTypePromptSet, models.PromptSetEvent{Prompt: "test_prompt"}),
		withEvent(models.ThreadEventTypePromptEdited, models.PromptEditedEvent{Prompt: "test_prompt_edited"}),
		withEvent(models.ThreadEventTypeSuggestionGenerated, models.SuggestionGeneratedEvent{
			SuggestionID: "test_suggestion_id_1",
			Recipe: models.RecipeBody{
				Title:       "test_recipe",
				Description: "test_recipe_description",
				Ingredients: []models.Ingredient{{Name: "test_ingredient"}},
				Steps:       []models.Step{"test_step"},
			},
		}),
		withEvent(models.ThreadEventTypeSuggestionGenerated, models.SuggestionGeneratedEvent{
			SuggestionID: "test_suggestion_id_2",
			Recipe: models.RecipeBody{
				Title:       "test_recipe_2",
				Description: "test_recipe_2_description",
				Ingredients: []models.Ingredient{{Name: "test_ingredient_2"}},
				Steps:       []models.Step{"test_step_2"},
			},
		}),
		withEvent(models.ThreadEventTypeSuggestionRejected, models.SuggestionRejectedEvent{SuggestionID: "test_suggestion_id_1"}),
		withEvent(models.ThreadEventTypeSuggestionAccepted, models.SuggestionAcceptedEvent{SuggestionID: "test_suggestion_id_2"}),
		withEvent(models.ThreadEventTypeRecipeModified, models.RecipeModifiedEvent{
			Recipe: models.RecipeBody{
				Title:       "test_recipe_3",
				Description: "test_recipe_3_description",
				Ingredients: []models.Ingredient{{Name: "test_ingredient_3"}},
				Steps:       []models.Step{"test_step_3"},
			},
		}),
		withEvent(models.ThreadEventTypeQuestionAnswered, models.QuestionAnsweredEvent{Question: "test_question", Answer: "test_answer"}),
	)
	thread, err := ReduceThreadEvents(threadID, events, nil)
	if err != nil {
		t.Fatalf("failed to reduce thread events: %v", err)
	}
	if thread.ID != threadID {
		t.Fatalf("expected thread ID %s, got %s", threadID, thread.ID)
	}
	if thread.OriginalPrompt != "test_prompt" {
		t.Fatalf("expected original prompt %s, got %s", "test_prompt", thread.OriginalPrompt)
	}
	if thread.CurrentPrompt != "test_prompt_edited" {
		t.Fatalf("expected current prompt %s, got %s", "test_prompt_edited", thread.CurrentPrompt)
	}
	if len(thread.Suggestions) != 2 {
		t.Fatalf("expected 2 suggestions, got %d", len(thread.Suggestions))
	}
	if thread.Suggestions[0].ID != "test_suggestion_id_1" {
		t.Fatalf("expected suggestion ID %s, got %s", "test_suggestion_id_1", thread.Suggestions[0].ID)
	}
	if thread.Suggestions[0].Accepted != false {
		t.Fatalf("expected suggestion accepted to be true, got %t", thread.Suggestions[0].Accepted)
	}
	if thread.Suggestions[0].Rejected != true {
		t.Fatalf("expected suggestion rejected to be false, got %t", thread.Suggestions[0].Rejected)
	}
	if thread.Suggestions[1].ID != "test_suggestion_id_2" {
		t.Fatalf("expected suggestion ID %s, got %s", "test_suggestion_id_2", thread.Suggestions[1].ID)
	}
	if thread.Suggestions[1].Accepted != true {
		t.Fatalf("expected suggestion accepted to be true, got %t", thread.Suggestions[1].Accepted)
	}
	if thread.Suggestions[1].Rejected != false {
		t.Fatalf("expected suggestion rejected to be false, got %t", thread.Suggestions[1].Rejected)
	}
	if thread.CurrentRecipe == nil || thread.CurrentRecipe.Title != "test_recipe_2" {
		t.Fatalf("expected recipe title %s, got %s", "test_recipe_2", thread.CurrentRecipe.Title)
	}
	if thread.CurrentRecipe.Description != "test_recipe_2_description" {
		t.Fatalf("expected recipe description %s, got %s", "test_recipe_2_description", thread.CurrentRecipe.Description)
	}
	if thread.CurrentRecipe.Ingredients[0].Name != "test_ingredient_2" {
		t.Fatalf("expected recipe ingredient name %s, got %s", "test_ingredient_2", thread.CurrentRecipe.Ingredients[0].Name)
	}
	if thread.CurrentRecipe.Steps[0] != "test_step_2" {
		t.Fatalf("expected recipe step %s, got %s", "test_step_2", thread.CurrentRecipe.Steps[0])
	}
	if thread.ModifiedRecipe == nil || thread.ModifiedRecipe.Title != "test_recipe_3" {
		t.Fatalf("expected modified recipe title %s, got %s", "test_recipe_3", thread.ModifiedRecipe.Title)
	}
	if thread.ModifiedRecipe.Description != "test_recipe_3_description" {
		t.Fatalf("expected modified recipe description %s, got %s", "test_recipe_3_description", thread.ModifiedRecipe.Description)
	}
	if thread.ModifiedRecipe.Ingredients[0].Name != "test_ingredient_3" {
		t.Fatalf("expected modified recipe ingredient name %s, got %s", "test_ingredient_3", thread.ModifiedRecipe.Ingredients[0].Name)
	}
	if thread.ModifiedRecipe.Steps[0] != "test_step_3" {
		t.Fatalf("expected modified recipe step %s, got %s", "test_step_3", thread.ModifiedRecipe.Steps[0])
	}
	if len(thread.ChatHistory) != 2 {
		t.Fatalf("expected 2 chat history entries, got %d", len(thread.ChatHistory))
	}
	if thread.ChatHistory[0].Source != "user" || thread.ChatHistory[0].Message != "test_question" {
		t.Fatalf("expected chat history entry source %s and message %s, got %s and %s", "user", "test_question", thread.ChatHistory[0].Source, thread.ChatHistory[0].Message)
	}
	if thread.ChatHistory[1].Source != "assistant" || thread.ChatHistory[1].Message != "test_answer" {
		t.Fatalf("expected chat history entry source %s and message %s, got %s and %s", "assistant", "test_answer", thread.ChatHistory[1].Source, thread.ChatHistory[1].Message)
	}

	events = createThreadEvents(t,
		withEvent(models.ThreadEventTypeRecipeModificationAccepted, models.RecipeModificationAcceptedEvent{}),
	)
	thread, err = ReduceThreadEvents(threadID, events, thread)
	if err != nil {
		t.Fatalf("failed to reduce thread events: %v", err)
	}
	if thread.CurrentRecipe == nil || thread.CurrentRecipe.Title != "test_recipe_3" {
		t.Fatalf("expected current recipe title %s, got %s", "test_recipe_3", thread.CurrentRecipe.Title)
	}
	if thread.ModifiedRecipe != nil {
		t.Fatalf("expected modified recipe to be nil, got %v", thread.ModifiedRecipe)
	}

	events = createThreadEvents(t,
		withEvent(models.ThreadEventTypeRecipeModified, models.RecipeModifiedEvent{Recipe: models.RecipeBody{Title: "test_recipe_4", Description: "test_recipe_4_description", Ingredients: []models.Ingredient{{Name: "test_ingredient_4"}}, Steps: []models.Step{"test_step_4"}}}),
		withEvent(models.ThreadEventTypeRecipeModificationRejected, models.RecipeModificationRejectedEvent{}),
	)
	thread, err = ReduceThreadEvents(threadID, events, thread)
	if err != nil {
		t.Fatalf("failed to reduce thread events: %v", err)
	}
	if thread.CurrentRecipe == nil || thread.CurrentRecipe.Title != "test_recipe_3" {
		t.Fatalf("expected current recipe title %s, got %s", "test_recipe_3", thread.CurrentRecipe.Title)
	}
	if thread.ModifiedRecipe != nil {
		t.Fatalf("expected modified recipe to be nil, got %v", thread.ModifiedRecipe)
	}
}

type threadEventOpt func(t *testing.T, event *models.ThreadEvent)

func withEvent(eventType models.ThreadEventType, payload any) threadEventOpt {
	return func(t *testing.T, event *models.ThreadEvent) {
		payload, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}
		event.Type = eventType
		event.Payload = payload
		event.Timestamp = time.Now()
	}
}

func createThreadEvents(t *testing.T, events ...threadEventOpt) []models.ThreadEvent {
	var threadEvents []models.ThreadEvent
	for _, event := range events {
		var e models.ThreadEvent
		event(t, &e)
		threadEvents = append(threadEvents, e)
	}
	return threadEvents
}
