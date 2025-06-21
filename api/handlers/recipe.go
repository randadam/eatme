package handlers

import "net/http"

// @Summary Generate a recipe
// @Description Generate a new recipe (not yet implemented)
// @Tags recipes
// @Produce plain
// @Success 200 {string} string "Recipe generation message"
// @Router /generate [get]
func GenerateRecipeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Recipe generation endpoint - not yet implemented."))
}
