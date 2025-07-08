package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/ajohnston1219/eatme/api/models"
	"github.com/google/uuid"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"golang.org/x/crypto/bcrypt"
	_ "modernc.org/sqlite"
)

type sqlRunner interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

type SQLiteStore struct {
	run sqlRunner
}

func NewSQLiteStore(dsn string) (*SQLiteStore, error) {
	driver := "libsql"
	if strings.HasPrefix(dsn, "file:") {
		log.Println("Using SQLite for local file")
		driver = "sqlite"
	}
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	if _, err := db.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		return nil, fmt.Errorf("failed to enable foreign keys: %w", err)
	}
	if err := migrate(db); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}
	return &SQLiteStore{db}, nil
}

func NewSQLiteStoreWithDB(db *sql.DB) (*SQLiteStore, error) {
	store := &SQLiteStore{db}
	if err := migrate(db); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}
	return store, nil
}

func (s *SQLiteStore) CreateUser(ctx context.Context, email, password string) (models.User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	id := uuid.NewString()

	_, err := s.run.ExecContext(ctx, `
		INSERT INTO users (id, email, password) VALUES (?, ?, ?);`,
		id, email, hash)
	if err != nil {
		if isUniqueViolation(err, "users.email") {
			return models.User{}, ErrEmailExists
		}
		return models.User{}, fmt.Errorf("failed to create user: %w", err)
	}
	return models.User{ID: id, Email: email}, nil
}

func (s *SQLiteStore) SaveProfile(ctx context.Context, userID string, p models.Profile) error {
	cuisines, err := json.Marshal(p.Cuisines)
	if err != nil {
		return fmt.Errorf("failed to marshal cuisines: %w", err)
	}
	diet, err := json.Marshal(p.Diet)
	if err != nil {
		return fmt.Errorf("failed to marshal diet: %w", err)
	}
	equipment, err := json.Marshal(p.Equipment)
	if err != nil {
		return fmt.Errorf("failed to marshal equipment: %w", err)
	}
	allergies, err := json.Marshal(p.Allergies)
	if err != nil {
		return fmt.Errorf("failed to marshal allergies: %w", err)
	}

	_, err = s.run.ExecContext(ctx, `
		INSERT INTO profiles (
		  user_id, setup_step, name, skill,
		  cuisines, diet, equipment, allergies
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(user_id) DO UPDATE SET
		  setup_step = excluded.setup_step,
		  name       = excluded.name,
		  skill      = excluded.skill,
		  cuisines   = excluded.cuisines,
		  diet       = excluded.diet,
		  equipment  = excluded.equipment,
		  allergies  = excluded.allergies
	`, userID, p.SetupStep, p.Name, p.Skill, cuisines, diet, equipment, allergies)
	if err != nil {
		return fmt.Errorf("failed to save profile: %w", err)
	}
	return nil
}

func (s *SQLiteStore) GetProfile(ctx context.Context, userID string) (models.Profile, error) {
	var p models.Profile
	var cuisines, diet, equipment, allergies []byte

	err := s.run.QueryRowContext(ctx, `
		SELECT setup_step, name, skill,
			COALESCE(cuisines, '[]'),
			COALESCE(diet, '[]'),
			COALESCE(equipment, '[]'),
			COALESCE(allergies, '[]')
		FROM profiles WHERE user_id = ?;
	`, userID).Scan(
		&p.SetupStep, &p.Name, &p.Skill,
		&cuisines, &diet, &equipment, &allergies,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return p, ErrNotFound
		}
		return p, fmt.Errorf("failed to get profile: %w", err)
	}

	if err := json.Unmarshal(cuisines, &p.Cuisines); err != nil {
		return p, fmt.Errorf("failed to unmarshal cuisines: %w", err)
	}
	if err := json.Unmarshal(diet, &p.Diet); err != nil {
		return p, fmt.Errorf("failed to unmarshal diet: %w", err)
	}
	if err := json.Unmarshal(equipment, &p.Equipment); err != nil {
		return p, fmt.Errorf("failed to unmarshal equipment: %w", err)
	}
	if err := json.Unmarshal(allergies, &p.Allergies); err != nil {
		return p, fmt.Errorf("failed to unmarshal allergies: %w", err)
	}

	return p, nil
}

func (s *SQLiteStore) GetGlobalRecipe(ctx context.Context, id string) (models.GlobalRecipe, error) {
	var recipe models.GlobalRecipe

	err := s.run.QueryRowContext(ctx, `
		SELECT id, title, description, total_time_minutes, servings,
			COALESCE(ingredients, '[]'),
			COALESCE(steps, '[]'),
			source_type,
			created_at,
			updated_at
		FROM global_recipes WHERE id = ?;
	`, id).Scan(
		&recipe.ID, &recipe.RecipeBody.Title, &recipe.RecipeBody.Description,
		&recipe.RecipeBody.TotalTimeMinutes, &recipe.RecipeBody.Servings,
		&recipe.RecipeBody.Ingredients, &recipe.RecipeBody.Steps,
		&recipe.SourceType, &recipe.CreatedAt, &recipe.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return recipe, ErrNotFound
		}
		return recipe, fmt.Errorf("failed to get global recipe: %w", err)
	}

	return recipe, nil
}

func (s *SQLiteStore) SaveGlobalRecipe(ctx context.Context, recipe models.GlobalRecipe) error {
	ingredients, err := json.Marshal(recipe.RecipeBody.Ingredients)
	if err != nil {
		return fmt.Errorf("failed to marshal ingredients: %w", err)
	}
	steps, err := json.Marshal(recipe.RecipeBody.Steps)
	if err != nil {
		return fmt.Errorf("failed to marshal steps: %w", err)
	}

	_, err = s.run.ExecContext(ctx, `
		INSERT INTO global_recipes (id, title, description, total_time_minutes, servings, ingredients, steps, source_type)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			title              = excluded.title,
			description        = excluded.description,
			total_time_minutes = excluded.total_time_minutes,
			servings           = excluded.servings,
			ingredients        = excluded.ingredients,
			steps              = excluded.steps,
			source_type        = excluded.source_type,
			updated_at         = CURRENT_TIMESTAMP;
	`, recipe.ID, recipe.RecipeBody.Title, recipe.RecipeBody.Description,
		recipe.RecipeBody.TotalTimeMinutes, recipe.RecipeBody.Servings, ingredients, steps, recipe.SourceType)
	if err != nil {
		return fmt.Errorf("failed to save global recipe: %w", err)
	}
	return nil
}

func (s *SQLiteStore) CreateSuggestionThread(ctx context.Context, userID string, thread models.SuggestionThread) error {
	_, err := s.run.ExecContext(ctx, `
		INSERT INTO recipe_suggestion_threads (id, user_id, original_prompt)
		VALUES (?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			user_id            = excluded.user_id,
			original_prompt    = excluded.original_prompt,
			updated_at         = CURRENT_TIMESTAMP;
	`, thread.ID, userID, thread.OriginalPrompt)
	if err != nil {
		return fmt.Errorf("failed to create suggestion thread: %w", err)
	}
	return nil
}

func (s *SQLiteStore) AppendToSuggestionThread(ctx context.Context, threadID string, suggestion models.RecipeSuggestion) error {
	recipeJson, err := json.Marshal(suggestion.Suggestion)
	if err != nil {
		return fmt.Errorf("failed to marshal recipe: %w", err)
	}
	_, err = s.run.ExecContext(ctx, `
		INSERT INTO recipe_suggestions (id, thread_id, recipe_json, response_text, accepted)
		VALUES (?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			thread_id          = excluded.thread_id,
			recipe_json        = excluded.recipe_json,
			response_text      = excluded.response_text,
			accepted           = excluded.accepted,
			updated_at         = CURRENT_TIMESTAMP;
	`, suggestion.ID, threadID, recipeJson, suggestion.ResponseText, suggestion.Accepted)
	if err != nil {
		return fmt.Errorf("failed to append to suggestion thread: %w", err)
	}
	return nil
}

func (s *SQLiteStore) AcceptSuggestion(ctx context.Context, threadID string, suggestion models.RecipeSuggestion) error {
	_, err := s.run.ExecContext(ctx, `
		UPDATE recipe_suggestions
		SET accepted = true
		WHERE id = ?;
	`, suggestion.ID)
	if err != nil {
		return fmt.Errorf("failed to accept suggestion: %w", err)
	}
	return nil
}

func (s *SQLiteStore) GetSuggestionThread(ctx context.Context, threadID string) (models.SuggestionThread, error) {
	var thread models.SuggestionThread
	err := s.run.QueryRowContext(ctx, `
		SELECT 
			id, original_prompt,
			created_at, updated_at
		FROM recipe_suggestion_threads WHERE id = ?;
	`, threadID).Scan(
		&thread.ID, &thread.OriginalPrompt,
		&thread.CreatedAt, &thread.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return thread, ErrNotFound
		}
		return thread, fmt.Errorf("failed to get suggestion thread: %w", err)
	}

	var suggestions []models.RecipeSuggestion
	rows, err := s.run.QueryContext(ctx, `
		SELECT 
			id, thread_id, recipe_json, response_text, accepted,
			created_at, updated_at
		FROM recipe_suggestions WHERE thread_id = ?;
	`, threadID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return thread, ErrNotFound
		}
		return thread, fmt.Errorf("failed to get suggestion thread: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var suggestion models.RecipeSuggestion
		var recipeJson []byte
		err := rows.Scan(
			&suggestion.ID, &suggestion.ThreadID, &recipeJson, &suggestion.ResponseText, &suggestion.Accepted,
			&suggestion.CreatedAt, &suggestion.UpdatedAt,
		)
		if err != nil {
			return thread, fmt.Errorf("failed to scan recipe suggestion: %w", err)
		}
		err = json.Unmarshal(recipeJson, &suggestion.Suggestion)
		if err != nil {
			return thread, fmt.Errorf("failed to unmarshal recipe suggestion: %w", err)
		}
		suggestions = append(suggestions, suggestion)
	}
	thread.Suggestions = suggestions

	return thread, nil
}

func (s *SQLiteStore) GetUserRecipe(ctx context.Context, userID string, recipeID string) (models.UserRecipe, error) {
	var recipe models.UserRecipe
	err := s.run.QueryRowContext(ctx, `
		SELECT 
			ur.id, ur.user_id, ur.global_recipe_id,
			ur.title, ur.description, ur.is_favorite,
			ur.latest_version_id, ur.created_at, ur.updated_at,
			rv.total_time_minutes, rv.servings,
			COALESCE(rv.ingredients, '[]'),
			COALESCE(rv.steps, '[]')
		FROM user_recipes ur
		JOIN recipe_versions rv ON ur.latest_version_id = rv.id
		WHERE ur.id = ? AND ur.user_id = ?;
	`, recipeID, userID).Scan(
		&recipe.ID, &recipe.UserID, &recipe.GlobalRecipeID, &recipe.Title, &recipe.Description, &recipe.IsFavorite,
		&recipe.LatestVersionID, &recipe.CreatedAt, &recipe.UpdatedAt,
		&recipe.RecipeBody.TotalTimeMinutes, &recipe.RecipeBody.Servings, &recipe.RecipeBody.Ingredients, &recipe.RecipeBody.Steps,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return recipe, ErrNotFound
		}
		return recipe, fmt.Errorf("failed to get user recipe: %w", err)
	}

	return recipe, nil
}

func (s *SQLiteStore) GetAllUserRecipes(ctx context.Context, userID string) ([]models.UserRecipe, error) {
	var recipes []models.UserRecipe
	rows, err := s.run.QueryContext(ctx, `
		SELECT 
			ur.id, ur.user_id, ur.global_recipe_id,
			ur.title, ur.description, ur.is_favorite,
			ur.latest_version_id, ur.created_at, ur.updated_at,
			rv.total_time_minutes, rv.servings,
			COALESCE(rv.ingredients, '[]'),
			COALESCE(rv.steps, '[]')
		FROM user_recipes ur
		JOIN recipe_versions rv ON ur.latest_version_id = rv.id
		WHERE ur.user_id = ?;
	`, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return recipes, ErrNotFound
		}
		return recipes, fmt.Errorf("failed to get user recipes: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var recipe models.UserRecipe
		err := rows.Scan(
			&recipe.ID, &recipe.UserID, &recipe.GlobalRecipeID, &recipe.Title, &recipe.Description, &recipe.IsFavorite,
			&recipe.LatestVersionID, &recipe.CreatedAt, &recipe.UpdatedAt,
			&recipe.RecipeBody.TotalTimeMinutes, &recipe.RecipeBody.Servings, &recipe.RecipeBody.Ingredients, &recipe.RecipeBody.Steps,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user recipe: %w", err)
		}
		recipes = append(recipes, recipe)
	}
	return recipes, nil
}

func (s *SQLiteStore) SaveUserRecipe(ctx context.Context, recipe models.UserRecipe) error {
	_, err := s.run.ExecContext(ctx, `
		INSERT INTO user_recipes (id, user_id, global_recipe_id, title, description, total_time_minutes, servings, is_favorite, latest_version_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			user_id            = excluded.user_id,
			global_recipe_id   = excluded.global_recipe_id,
			title              = excluded.title,
			description        = excluded.description,
			total_time_minutes = excluded.total_time_minutes,
			servings           = excluded.servings,
			is_favorite        = excluded.is_favorite,
			latest_version_id  = excluded.latest_version_id,
			updated_at         = CURRENT_TIMESTAMP;
	`, recipe.ID, recipe.UserID, recipe.GlobalRecipeID, recipe.Title, recipe.Description,
		recipe.RecipeBody.TotalTimeMinutes, recipe.RecipeBody.Servings, recipe.IsFavorite,
		recipe.LatestVersionID)
	if err != nil {
		return fmt.Errorf("failed to save user recipe: %w", err)
	}
	return nil
}

func (s *SQLiteStore) DeleteUserRecipe(ctx context.Context, userID string, recipeID string) error {
	_, err := s.run.ExecContext(ctx, `
		DELETE FROM user_recipes WHERE id = ? AND user_id = ?;`, recipeID, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user recipe: %w", err)
	}
	return nil
}

func (s *SQLiteStore) UpdateUserRecipeVersion(ctx context.Context, userID string, recipeID string, version models.RecipeVersion) error {
	_, err := s.run.ExecContext(ctx, `
		UPDATE user_recipes
		SET 
		    title              = ?,
		    description        = ?,
			total_time_minutes = ?,
			servings           = ?,
			latest_version_id  = ?
		WHERE id = ? AND user_id = ?;
	`, version.Title, version.Description, version.TotalTimeMinutes, version.Servings, version.ID, recipeID, userID)
	if err != nil {
		return fmt.Errorf("failed to update user recipe version: %w", err)
	}
	return nil
}

func (s *SQLiteStore) GetRecipeVersion(ctx context.Context, recipeVersionID string) (models.RecipeVersion, error) {
	var recipeVersion models.RecipeVersion

	err := s.run.QueryRowContext(ctx, `
		SELECT 
			id, user_recipe_id, parent_id,
			total_time_minutes, servings,
			COALESCE(ingredients, '[]'),
			COALESCE(steps, '[]'),
			notes, created_at
		FROM recipe_versions WHERE id = ?;
	`, recipeVersionID).Scan(
		&recipeVersion.ID, &recipeVersion.UserRecipeID, &recipeVersion.ParentID,
		&recipeVersion.TotalTimeMinutes, &recipeVersion.Servings, &recipeVersion.Ingredients, &recipeVersion.Steps,
		&recipeVersion.Notes, &recipeVersion.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return recipeVersion, ErrNotFound
		}
		return recipeVersion, fmt.Errorf("failed to get recipe version: %w", err)
	}

	return recipeVersion, nil
}

func (s *SQLiteStore) AddRecipeVersion(ctx context.Context, recipeVersion models.RecipeVersion) error {
	ingredients, err := json.Marshal(recipeVersion.Ingredients)
	if err != nil {
		return fmt.Errorf("failed to marshal ingredients: %w", err)
	}
	steps, err := json.Marshal(recipeVersion.Steps)
	if err != nil {
		return fmt.Errorf("failed to marshal steps: %w", err)
	}

	_, err = s.run.ExecContext(ctx, `
		INSERT INTO recipe_versions (id, user_recipe_id, parent_id, total_time_minutes, servings, ingredients, steps, notes, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			user_recipe_id     = excluded.user_recipe_id,
			parent_id          = excluded.parent_id,
			total_time_minutes = excluded.total_time_minutes,
			servings           = excluded.servings,
			ingredients        = excluded.ingredients,
			steps              = excluded.steps,
			notes              = excluded.notes;
	`, recipeVersion.ID, recipeVersion.UserRecipeID, recipeVersion.ParentID,
		recipeVersion.TotalTimeMinutes, recipeVersion.Servings, ingredients, steps, recipeVersion.Notes, recipeVersion.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to save recipe version: %w", err)
	}
	return nil
}

func (s *SQLiteStore) SaveMealPlan(ctx context.Context, userID string, mealPlan models.MealPlan) error {
	recipes, err := json.Marshal(mealPlan.Recipes)
	if err != nil {
		return fmt.Errorf("failed to marshal recipes: %w", err)
	}

	_, err = s.run.ExecContext(ctx, `
		INSERT INTO meal_plans (id, user_id, recipes) VALUES (?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET recipes = excluded.recipes;
	`, mealPlan.ID, userID, recipes)
	if err != nil {
		return fmt.Errorf("failed to save meal plan: %w", err)
	}
	return nil
}

func (s *SQLiteStore) GetAllPlans(ctx context.Context, userID string) ([]models.MealPlan, error) {
	var mealPlans []models.MealPlan
	var recipes []byte

	rows, err := s.run.QueryContext(ctx, `
		SELECT id, user_id, recipes FROM meal_plans WHERE user_id = ?;
	`, userID)
	if err != nil {
		return mealPlans, fmt.Errorf("failed to get meal plans: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var mealPlan models.MealPlan
		if err := rows.Scan(&mealPlan.ID, &mealPlan.UserID, &recipes); err != nil {
			return mealPlans, fmt.Errorf("failed to scan meal plan: %w", err)
		}
		if err := json.Unmarshal(recipes, &mealPlan.Recipes); err != nil {
			return mealPlans, fmt.Errorf("failed to unmarshal recipes: %w", err)
		}
		mealPlans = append(mealPlans, mealPlan)
	}
	return mealPlans, nil
}

func (s *SQLiteStore) GetMealPlan(ctx context.Context, userID string, mealPlanID string) (models.MealPlan, error) {
	var mealPlan models.MealPlan
	var recipes []byte

	err := s.run.QueryRowContext(ctx, `
		SELECT id, user_id, recipes FROM meal_plans WHERE id = ? AND user_id = ?;
	`, mealPlanID, userID).Scan(&mealPlan.ID, &mealPlan.UserID, &recipes)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return mealPlan, ErrNotFound
		}
		return mealPlan, fmt.Errorf("failed to get meal plan: %w", err)
	}

	if err := json.Unmarshal(recipes, &mealPlan.Recipes); err != nil {
		return mealPlan, fmt.Errorf("failed to unmarshal recipes: %w", err)
	}
	return mealPlan, nil
}

func (s *SQLiteStore) WithTx(fn func(tx Store) error) error {
	tx, err := s.run.(*sql.DB).BeginTx(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	txStore := &SQLiteStore{run: tx}
	if err := fn(txStore); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func migrate(db *sql.DB) error {
	const users = `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		email TEXT UNIQUE NOT NULL,
		password BLOB NOT NULL
	);`

	const profiles = `
	CREATE TABLE IF NOT EXISTS profiles (
		user_id    TEXT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
		setup_step TEXT NOT NULL DEFAULT 'account',
		name       TEXT NOT NULL DEFAULT '',
		skill      TEXT NOT NULL DEFAULT 'beginner',
		cuisines   JSON NOT NULL DEFAULT '[]',
		diet       JSON NOT NULL DEFAULT '[]',
		equipment  JSON NOT NULL DEFAULT '[]',
		allergies  JSON NOT NULL DEFAULT '[]'
	);`

	const globalRecipes = `
	CREATE TABLE IF NOT EXISTS global_recipes (
		id                 TEXT PRIMARY KEY,
		title              TEXT NOT NULL,
		description        TEXT NOT NULL,
		total_time_minutes INTEGER NOT NULL,
		servings           INTEGER NOT NULL,
		ingredients        JSON NOT NULL DEFAULT '[]',
		steps              JSON NOT NULL DEFAULT '[]',
		source_type        TEXT NOT NULL DEFAULT 'generated',
		created_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	const recipeSuggestionThreads = `
	CREATE TABLE IF NOT EXISTS recipe_suggestion_threads (
		id                 TEXT PRIMARY KEY,
		user_id            TEXT REFERENCES users(id) ON DELETE CASCADE,
		original_prompt    TEXT NOT NULL,
		created_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	const recipeSuggestions = `
	CREATE TABLE IF NOT EXISTS recipe_suggestions (
		id                 TEXT PRIMARY KEY,
		thread_id          TEXT REFERENCES recipe_suggestion_threads(id) ON DELETE CASCADE,
		recipe_json        JSON NOT NULL DEFAULT '{}',
		response_text      TEXT NOT NULL,
		accepted           BOOLEAN NOT NULL DEFAULT FALSE,
		created_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	const userRecipes = `
	CREATE TABLE IF NOT EXISTS user_recipes (
		id                 TEXT PRIMARY KEY,
		user_id            TEXT REFERENCES users(id) ON DELETE CASCADE,
		global_recipe_id   TEXT NULL REFERENCES global_recipes(id) ON DELETE SET NULL,
		title              TEXT NOT NULL,
		description        TEXT NOT NULL,
		total_time_minutes INTEGER NOT NULL,
		servings           INTEGER NOT NULL,
		is_favorite        BOOLEAN DEFAULT FALSE,
		latest_version_id  TEXT NULL,
		created_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	const recipeVersions = `
	CREATE TABLE IF NOT EXISTS recipe_versions (
		id                 TEXT PRIMARY KEY,
		user_recipe_id     TEXT,
		parent_id          TEXT,
		created_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		total_time_minutes INTEGER NOT NULL,
		servings           INTEGER NOT NULL,
		ingredients        JSON NOT NULL DEFAULT '[]',
		steps              JSON NOT NULL DEFAULT '[]',
		notes              TEXT NULL
	);`

	const mealPlans = `
	CREATE TABLE IF NOT EXISTS meal_plans (
		id         TEXT PRIMARY KEY,
		name       TEXT NOT NULL,
		user_id    TEXT REFERENCES users(id) ON DELETE CASCADE,
		recipes    JSON NOT NULL DEFAULT '[]'
	);`

	if _, err := db.Exec(users); err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}
	if _, err := db.Exec(profiles); err != nil {
		return fmt.Errorf("failed to create profiles table: %w", err)
	}
	if _, err := db.Exec(globalRecipes); err != nil {
		return fmt.Errorf("failed to create global_recipes table: %w", err)
	}
	if _, err := db.Exec(recipeSuggestionThreads); err != nil {
		return fmt.Errorf("failed to create recipe_suggestion_threads table: %w", err)
	}
	if _, err := db.Exec(recipeSuggestions); err != nil {
		return fmt.Errorf("failed to create recipe_suggestions table: %w", err)
	}
	if _, err := db.Exec(userRecipes); err != nil {
		return fmt.Errorf("failed to create user_recipes table: %w", err)
	}
	if _, err := db.Exec(recipeVersions); err != nil {
		return fmt.Errorf("failed to create recipe_versions table: %w", err)
	}
	if _, err := db.Exec(mealPlans); err != nil {
		return fmt.Errorf("failed to create meal_plans table: %w", err)
	}
	return nil
}

func isUniqueViolation(err error, field string) bool {
	return err != nil && strings.Contains(err.Error(), "UNIQUE constraint failed: "+field)
}
