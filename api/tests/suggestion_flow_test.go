package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/ajohnston1219/eatme/api/models"
	"github.com/stretchr/testify/require"
)

func TestSuggestionFlow(t *testing.T) {
	ml := &MLStub{
		Responses: []models.SuggestChatResponse{
			{ResponseText: "How about Beef Stroganoff?", NewRecipe: makeFakeRecipe("Beef Stroganoff")},
			{ResponseText: "Maybe Beef & Mushroom Tacos?", NewRecipe: makeFakeRecipe("Beef & Mushroom Tacos")},
		},
	}
	ts, store := NewTestServer(t, ml)
	defer ts.Close()

	userID := "user-123"
	authToken := "Bearer " + userID
	createUser(store, userID)
	prompt := "beef & mushrooms"

	// 1️⃣ start thread
	body, _ := json.Marshal(models.SuggestChatRequest{Message: prompt})
	req, err := http.NewRequest("POST", ts.URL+"/chat/suggest", bytes.NewReader(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// check response
	var suggestionResponse models.SuggestChatResponse
	json.NewDecoder(resp.Body).Decode(&suggestionResponse)
	fmt.Printf("suggestionResponse: %+v\n", suggestionResponse)
	require.NotEmpty(t, suggestionResponse.ThreadID)
	require.NotEmpty(t, suggestionResponse.NewRecipe.Title)
	require.Equal(t, "How about Beef Stroganoff?", suggestionResponse.ResponseText)
	require.Equal(t, "Beef Stroganoff", suggestionResponse.NewRecipe.Title)

	// check thread
	thread, err := store.GetSuggestionThread(context.Background(), suggestionResponse.ThreadID)
	require.NoError(t, err)
	require.Equal(t, prompt, thread.OriginalPrompt)
	require.Equal(t, "How about Beef Stroganoff?", thread.Suggestions[0].ResponseText)
	require.Equal(t, 1, len(thread.Suggestions))
	require.Equal(t, "Beef Stroganoff", thread.Suggestions[0].Suggestion.Title)
	require.Equal(t, false, thread.Suggestions[0].Accepted)

	// 2️⃣ reject first -> next
	req, err = http.NewRequest("GET", ts.URL+"/chat/suggest/"+suggestionResponse.ThreadID+"/next", nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", authToken)
	resp, err = http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var next models.SuggestChatResponse
	json.NewDecoder(resp.Body).Decode(&next)
	require.Equal(t, "Beef & Mushroom Tacos", next.NewRecipe.Title)
	require.Equal(t, "Maybe Beef & Mushroom Tacos?", next.ResponseText)

	// check thread
	thread, err = store.GetSuggestionThread(context.Background(), suggestionResponse.ThreadID)
	require.NoError(t, err)
	require.Equal(t, prompt, thread.OriginalPrompt)
	require.Equal(t, 2, len(thread.Suggestions))
	require.Equal(t, "Beef Stroganoff", thread.Suggestions[0].Suggestion.Title)
	require.Equal(t, "How about Beef Stroganoff?", thread.Suggestions[0].ResponseText)
	require.Equal(t, false, thread.Suggestions[0].Accepted)
	require.Equal(t, "Beef & Mushroom Tacos", thread.Suggestions[1].Suggestion.Title)
	require.Equal(t, "Maybe Beef & Mushroom Tacos?", thread.Suggestions[1].ResponseText)
	require.Equal(t, false, thread.Suggestions[1].Accepted)

	// 3️⃣ accept
	acceptURL := ts.URL + "/chat/suggest/" + suggestionResponse.ThreadID + "/accept"
	req, err = http.NewRequest("POST", acceptURL, nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", authToken)
	resp, err = http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	var accepted models.UserRecipe
	json.NewDecoder(resp.Body).Decode(&accepted)
	require.Equal(t, "Beef & Mushroom Tacos", accepted.Title)

	// check thread
	thread, err = store.GetSuggestionThread(context.Background(), suggestionResponse.ThreadID)
	require.NoError(t, err)
	require.Equal(t, 2, len(thread.Suggestions))
	require.Equal(t, "Beef Stroganoff", thread.Suggestions[0].Suggestion.Title)
	require.Equal(t, "How about Beef Stroganoff?", thread.Suggestions[0].ResponseText)
	require.Equal(t, false, thread.Suggestions[0].Accepted)
	require.Equal(t, "Beef & Mushroom Tacos", thread.Suggestions[1].Suggestion.Title)
	require.Equal(t, "Maybe Beef & Mushroom Tacos?", thread.Suggestions[1].ResponseText)
	require.Equal(t, true, thread.Suggestions[1].Accepted)

	// 4️⃣ verify DB: verify a recipe was created
	recipe, err := store.GetUserRecipe(context.Background(), userID, accepted.ID)
	require.NoError(t, err)
	require.Equal(t, "Beef & Mushroom Tacos", recipe.Title)
}
