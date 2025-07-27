package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/ajohnston1219/eatme/api/internal/models"
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

func (s *SQLiteStore) GetUser(ctx context.Context, userID string) (models.User, error) {
	var user models.User
	err := s.run.QueryRowContext(ctx, `
		SELECT id, email
		FROM users WHERE id = ?;
	`, userID).Scan(&user.ID, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, ErrNotFound
		}
		return models.User{}, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func (s *SQLiteStore) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := s.run.QueryRowContext(ctx, `
		SELECT id, email
		FROM users WHERE email = ?;
	`, email).Scan(&user.ID, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, ErrNotFound
		}
		return models.User{}, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func (s *SQLiteStore) CheckPassword(ctx context.Context, userID string, password string) error {
	var hash []byte
	err := s.run.QueryRowContext(ctx, `
		SELECT password
		FROM users WHERE id = ?;
	`, userID).Scan(&hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFound
		}
		return fmt.Errorf("failed to get user: %w", err)
	}
	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}

func (s *SQLiteStore) SaveProfile(ctx context.Context, userID string, p models.Profile) error {
	cuisines, err := json.Marshal(p.Cuisines)
	if err != nil {
		return fmt.Errorf("failed to marshal cuisines: %w", err)
	}
	diets, err := json.Marshal(p.Diets)
	if err != nil {
		return fmt.Errorf("failed to marshal diets: %w", err)
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
		  cuisines, diets, equipment, allergies
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(user_id) DO UPDATE SET
		  setup_step = excluded.setup_step,
		  name       = excluded.name,
		  skill      = excluded.skill,
		  cuisines   = excluded.cuisines,
		  diets      = excluded.diets,
		  equipment  = excluded.equipment,
		  allergies  = excluded.allergies
	`, userID, p.SetupStep, p.Name, p.Skill, cuisines, diets, equipment, allergies)
	if err != nil {
		return fmt.Errorf("failed to save profile: %w", err)
	}
	return nil
}

func (s *SQLiteStore) GetProfile(ctx context.Context, userID string) (models.Profile, error) {
	var p models.Profile
	var cuisines, diets, equipment, allergies []byte

	err := s.run.QueryRowContext(ctx, `
		SELECT setup_step, name, skill,
			COALESCE(cuisines, '[]'),
			COALESCE(diets, '[]'),
			COALESCE(equipment, '[]'),
			COALESCE(allergies, '[]')
		FROM profiles WHERE user_id = ?;
	`, userID).Scan(
		&p.SetupStep, &p.Name, &p.Skill,
		&cuisines, &diets, &equipment, &allergies,
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
	if err := json.Unmarshal(diets, &p.Diets); err != nil {
		return p, fmt.Errorf("failed to unmarshal diets: %w", err)
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
		SELECT id, title, description, total_time_minutes, servings, image_url,
			COALESCE(ingredients, '[]'),
			COALESCE(steps, '[]'),
			source_type,
			created_at,
			updated_at
		FROM global_recipes WHERE id = ?;
	`, id).Scan(
		&recipe.ID, &recipe.RecipeBody.Title, &recipe.RecipeBody.Description,
		&recipe.RecipeBody.ImageURL, &recipe.RecipeBody.TotalTimeMinutes,
		&recipe.RecipeBody.Servings, &recipe.RecipeBody.Ingredients,
		&recipe.RecipeBody.Steps, &recipe.SourceType,
		&recipe.CreatedAt, &recipe.UpdatedAt,
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
		INSERT INTO global_recipes (id, title, description, total_time_minutes, servings, image_url, ingredients, steps, source_type)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
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
		recipe.RecipeBody.TotalTimeMinutes, recipe.RecipeBody.Servings, recipe.RecipeBody.ImageURL,
		ingredients, steps, recipe.SourceType)
	if err != nil {
		return fmt.Errorf("failed to save global recipe: %w", err)
	}
	return nil
}

func (s *SQLiteStore) CreateThread(ctx context.Context, userID string, thread models.Thread) error {
	_, err := s.run.ExecContext(ctx, `
		INSERT INTO threads (id, user_id, thread_type)
		VALUES (?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			user_id            = excluded.user_id,
			thread_type        = excluded.thread_type,
			updated_at         = CURRENT_TIMESTAMP;
	`, thread.ID, userID, thread.Type)
	if err != nil {
		return fmt.Errorf("failed to create thread: %w", err)
	}
	for i, event := range thread.Events {
		eventId := uuid.NewString()
		_, err = s.run.ExecContext(ctx, `
			INSERT INTO thread_events (id, thread_id, event_index, event_type, payload)
			VALUES (?, ?, ?, ?, ?)
		`, eventId, thread.ID, i, event.Type, event.Payload)
		if err != nil {
			return fmt.Errorf("failed to append to thread: %w", err)
		}
	}
	return nil
}

func (s *SQLiteStore) AppendToThread(ctx context.Context, threadId string, events []models.ThreadEvent) error {
	for _, event := range events {
		eventId := uuid.NewString()
		_, err := s.run.ExecContext(ctx, `
			WITH next_index AS (
				SELECT MAX(event_index) + 1 AS next_index
				FROM thread_events
				WHERE thread_id = ?
			)
			INSERT INTO thread_events (thread_id, id, event_index, event_type, payload)
			VALUES (?, ?, (SELECT next_index FROM next_index), ?, ?)
		`, threadId, threadId, eventId, event.Type, event.Payload)
		if err != nil {
			return fmt.Errorf("failed to append to thread: %w", err)
		}
	}
	return nil
}

func (s *SQLiteStore) AssociateThreadWithRecipe(ctx context.Context, threadID string, recipeID string) error {
	_, err := s.run.ExecContext(ctx, `
		UPDATE threads SET recipe_id = ? WHERE id = ?;
	`, recipeID, threadID)
	if err != nil {
		return fmt.Errorf("failed to associate recipe with thread: %w", err)
	}
	return nil
}

func (s *SQLiteStore) GetThread(ctx context.Context, threadID string) (models.Thread, error) {
	var thread models.Thread
	err := s.run.QueryRowContext(ctx, `
		SELECT 
			id,
			thread_type,
			recipe_id,
			created_at,
			updated_at
		FROM threads WHERE id = ?;
	`, threadID).Scan(
		&thread.ID, &thread.Type, &thread.RecipeID, &thread.CreatedAt, &thread.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return thread, ErrNotFound
		}
		return thread, fmt.Errorf("failed to get thread: %w", err)
	}

	var events []models.ThreadEvent
	rows, err := s.run.QueryContext(ctx, `
		SELECT 
			event_type, payload, created_at
		FROM thread_events WHERE thread_id = ?
		ORDER BY event_index ASC;
	`, threadID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return thread, ErrNotFound
		}
		return thread, fmt.Errorf("failed to get thread: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var event models.ThreadEvent
		err := rows.Scan(
			&event.Type, &event.Payload, &event.Timestamp,
		)
		if err != nil {
			return thread, fmt.Errorf("failed to scan thread event: %w", err)
		}
		events = append(events, event)
	}
	thread.Events = events

	return thread, nil
}

func (s *SQLiteStore) GetUserRecipe(ctx context.Context, userID string, recipeID string) (models.UserRecipe, error) {
	var recipe models.UserRecipe
	err := s.run.QueryRowContext(ctx, `
		SELECT 
			ur.id, ur.user_id, ur.thread_id, ur.global_recipe_id,
			ur.title, ur.description, ur.is_favorite, ur.image_url,
			ur.latest_version_id, ur.created_at, ur.updated_at,
			rv.total_time_minutes, rv.servings,
			COALESCE(rv.ingredients, '[]'),
			COALESCE(rv.steps, '[]')
		FROM user_recipes ur
		JOIN recipe_versions rv ON ur.latest_version_id = rv.id
		WHERE ur.id = ? AND ur.user_id = ?;
	`, recipeID, userID).Scan(
		&recipe.ID, &recipe.UserID, &recipe.ThreadID, &recipe.GlobalRecipeID, &recipe.Title, &recipe.Description, &recipe.IsFavorite,
		&recipe.ImageURL, &recipe.LatestVersionID, &recipe.CreatedAt, &recipe.UpdatedAt,
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
			ur.id, ur.user_id, ur.thread_id, ur.global_recipe_id,
			ur.title, ur.description, ur.is_favorite, ur.image_url,
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
			&recipe.ID, &recipe.UserID, &recipe.ThreadID, &recipe.GlobalRecipeID, &recipe.Title, &recipe.Description, &recipe.IsFavorite,
			&recipe.ImageURL, &recipe.LatestVersionID, &recipe.CreatedAt, &recipe.UpdatedAt,
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
		INSERT INTO user_recipes (id, user_id, global_recipe_id, thread_id, title, description, total_time_minutes, servings, is_favorite, image_url, latest_version_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			user_id            = excluded.user_id,
			global_recipe_id   = excluded.global_recipe_id,
			thread_id          = excluded.thread_id,
			title              = excluded.title,
			description        = excluded.description,
			total_time_minutes = excluded.total_time_minutes,
			servings           = excluded.servings,
			is_favorite        = excluded.is_favorite,
			image_url          = excluded.image_url,
			latest_version_id  = excluded.latest_version_id,
			updated_at         = CURRENT_TIMESTAMP;
	`, recipe.ID, recipe.UserID, recipe.GlobalRecipeID, recipe.ThreadID, recipe.Title, recipe.Description,
		recipe.RecipeBody.TotalTimeMinutes, recipe.RecipeBody.Servings, recipe.IsFavorite,
		recipe.ImageURL, recipe.LatestVersionID)
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
			image_url          = ?,
			latest_version_id  = ?
		WHERE id = ? AND user_id = ?;
	`, version.Title, version.Description, version.TotalTimeMinutes, version.Servings, version.ImageURL, version.ID, recipeID, userID)
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
			image_url,
			COALESCE(ingredients, '[]'),
			COALESCE(steps, '[]'),
			notes, created_at
		FROM recipe_versions WHERE id = ?;
	`, recipeVersionID).Scan(
		&recipeVersion.ID, &recipeVersion.UserRecipeID, &recipeVersion.ParentID,
		&recipeVersion.TotalTimeMinutes, &recipeVersion.Servings, &recipeVersion.ImageURL,
		&recipeVersion.Ingredients, &recipeVersion.Steps,
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
		INSERT INTO recipe_versions (id, user_recipe_id, parent_id, total_time_minutes, servings, image_url, ingredients, steps, notes, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			user_recipe_id     = excluded.user_recipe_id,
			parent_id          = excluded.parent_id,
			total_time_minutes = excluded.total_time_minutes,
			servings           = excluded.servings,
			image_url          = excluded.image_url,
			ingredients        = excluded.ingredients,
			steps              = excluded.steps,
			notes              = excluded.notes;
	`, recipeVersion.ID, recipeVersion.UserRecipeID, recipeVersion.ParentID,
		recipeVersion.TotalTimeMinutes, recipeVersion.Servings, recipeVersion.ImageURL, ingredients, steps, recipeVersion.Notes, recipeVersion.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to save recipe version: %w", err)
	}
	return nil
}

func (s *SQLiteStore) UpdateRecipeVersion(ctx context.Context, newVersion models.RecipeVersion) error {
	_, err := s.run.ExecContext(ctx, `
		UPDATE recipe_versions SET
		    user_recipe_id     = ?,
		    parent_id          = ?,
		    total_time_minutes = ?,
		    servings           = ?,
		    image_url          = ?,
		    ingredients        = ?,
		    steps              = ?,
		    notes              = ?
		WHERE id = ? AND user_recipe_id = ?;
	`, newVersion.UserRecipeID, newVersion.ParentID,
		newVersion.TotalTimeMinutes, newVersion.Servings, newVersion.ImageURL, newVersion.Ingredients, newVersion.Steps,
		newVersion.Notes, newVersion.ID, newVersion.UserRecipeID)
	if err != nil {
		return fmt.Errorf("failed to update recipe version: %w", err)
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
		diets      JSON NOT NULL DEFAULT '[]',
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
		image_url          TEXT NULL,
		ingredients        JSON NOT NULL DEFAULT '[]',
		steps              JSON NOT NULL DEFAULT '[]',
		source_type        TEXT NOT NULL DEFAULT 'generated',
		created_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	const threads = `
	CREATE TABLE IF NOT EXISTS threads (
		id                 TEXT PRIMARY KEY,
		thread_type        TEXT NOT NULL,
		recipe_id          TEXT NULL,
		user_id            TEXT REFERENCES users(id) ON DELETE CASCADE,
		created_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	const threadEvents = `
	CREATE TABLE IF NOT EXISTS thread_events (
		id                 TEXT PRIMARY KEY,
		thread_id          TEXT REFERENCES threads(id) ON DELETE CASCADE,
		event_index        INTEGER NOT NULL,
		event_type         TEXT NOT NULL,
		payload            JSON NOT NULL,
		created_at         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	const userRecipes = `
	CREATE TABLE IF NOT EXISTS user_recipes (
		id                 TEXT PRIMARY KEY,
		user_id            TEXT REFERENCES users(id) ON DELETE CASCADE,
		thread_id          TEXT REFERENCES threads(id) ON DELETE CASCADE,
		global_recipe_id   TEXT NULL REFERENCES global_recipes(id) ON DELETE SET NULL,
		title              TEXT NOT NULL,
		description        TEXT NOT NULL,
		total_time_minutes INTEGER NOT NULL,
		servings           INTEGER NOT NULL,
		image_url          TEXT NULL,
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
		image_url          TEXT NULL,
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
	if _, err := db.Exec(threads); err != nil {
		return fmt.Errorf("failed to create threads table: %w", err)
	}
	if _, err := db.Exec(threadEvents); err != nil {
		return fmt.Errorf("failed to create thread_events table: %w", err)
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
