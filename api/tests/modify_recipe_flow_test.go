package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/ajohnston1219/eatme/api/internal/models"
	"github.com/stretchr/testify/require"
)

func TestModifyRecipeFlow(t *testing.T) {
	initialRecipe := makeFakeRecipe(
		"Beef Bolonese",
		WithIngredients([]models.Ingredient{{Name: "Beef"}, {Name: "Tomato"}}),
		WithSteps([]string{"Step 1", "Step 2"}),
		WithTotalTimeMinutes(60),
		WithServings(4),
	)
	modifiedRecipe := makeFakeRecipe(
		"Chicken Bolonese",
		WithIngredients([]models.Ingredient{{Name: "Chicken"}, {Name: "Tomato"}}),
		WithSteps([]string{"Updated Step 1", "Updated Step 2"}),
		WithTotalTimeMinutes(45),
		WithServings(2),
	)
	ml := &MLStub{
		ModifyResponses: []models.ModifyChatResponse{
			{ResponseText: "Here you go", NewRecipe: modifiedRecipe},
		},
	}
	ts, store := NewTestServer(t, ml)
	defer ts.Close()

	email := "user@example.com"
	user, err := createUser(store, email)
	require.NoError(t, err)
	authToken := "Bearer " + user.ID

	newRecipe, err := createRecipe(store, user.ID, initialRecipe)
	require.NoError(t, err)

	// Modify Recipe
	body, _ := json.Marshal(models.ModifyChatRequest{Message: "make this with chicken"})
	req, err := http.NewRequest("PUT", ts.URL+"/chat/modify/recipes/"+newRecipe.ID, bytes.NewReader(body))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// Check Response
	var response models.ModifyChatResponse
	json.NewDecoder(resp.Body).Decode(&response)
	require.Equal(t, "Here you go", response.ResponseText)
	require.Equal(t, "Chicken Bolonese", response.NewRecipe.Title)
	require.Equal(t, "Chicken", response.NewRecipe.Ingredients[0].Name)
	require.Equal(t, "Tomato", response.NewRecipe.Ingredients[1].Name)
	require.Equal(t, "Updated Step 1", response.NewRecipe.Steps[0])
	require.Equal(t, "Updated Step 2", response.NewRecipe.Steps[1])
	require.Equal(t, 45, response.NewRecipe.TotalTimeMinutes)
	require.Equal(t, 2, response.NewRecipe.Servings)

	// Check Recipe
	req, err = http.NewRequest("GET", ts.URL+"/recipes/"+newRecipe.ID, nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", authToken)
	resp, err = http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var recipe models.UserRecipe
	json.NewDecoder(resp.Body).Decode(&recipe)
	require.Equal(t, "Chicken Bolonese", recipe.Title)
	require.Equal(t, "Chicken", recipe.Ingredients[0].Name)
	require.Equal(t, "Tomato", recipe.Ingredients[1].Name)
	require.Equal(t, "Updated Step 1", recipe.Steps[0])
	require.Equal(t, "Updated Step 2", recipe.Steps[1])
	require.Equal(t, 45, recipe.TotalTimeMinutes)
	require.Equal(t, 2, recipe.Servings)
}
