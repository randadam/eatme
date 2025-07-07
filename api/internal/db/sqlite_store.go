package db

import (
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

type SQLiteStore struct{ *sql.DB }

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

func (s *SQLiteStore) CreateUser(email, password string) (models.User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	id := uuid.NewString()

	_, err := s.Exec(`
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

func (s *SQLiteStore) SaveProfile(userID string, p models.Profile) error {
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

	_, err = s.Exec(`
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

func (s *SQLiteStore) GetProfile(userID string) (models.Profile, error) {
	var p models.Profile
	var cuisines, diet, equipment, allergies []byte

	err := s.QueryRow(`
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

func (s *SQLiteStore) SaveMealPlan(userID string, mealPlan models.MealPlan) error {
	recipes, err := json.Marshal(mealPlan.Recipes)
	if err != nil {
		return fmt.Errorf("failed to marshal recipes: %w", err)
	}

	_, err = s.Exec(`
		INSERT INTO meal_plans (id, user_id, recipes) VALUES (?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET recipes = excluded.recipes;
	`, mealPlan.ID, userID, recipes)
	if err != nil {
		return fmt.Errorf("failed to save meal plan: %w", err)
	}
	return nil
}

func (s *SQLiteStore) GetMealPlan(userID string, mealPlanID string) (models.MealPlan, error) {
	var mealPlan models.MealPlan
	var recipes []byte

	err := s.QueryRow(`
		SELECT recipes FROM meal_plans WHERE id = ? AND user_id = ?;
	`, mealPlanID, userID).Scan(&recipes)
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

	const meal_plans = `
	CREATE TABLE IF NOT EXISTS meal_plans (
		id         TEXT PRIMARY KEY,
		user_id    TEXT REFERENCES users(id) ON DELETE CASCADE,
		recipes    JSON NOT NULL DEFAULT '[]'
	);`

	if _, err := db.Exec(users); err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}
	if _, err := db.Exec(profiles); err != nil {
		return fmt.Errorf("failed to create profiles table: %w", err)
	}
	if _, err := db.Exec(meal_plans); err != nil {
		return fmt.Errorf("failed to create meal_plans table: %w", err)
	}
	return nil
}

func isUniqueViolation(err error, field string) bool {
	return err != nil && strings.Contains(err.Error(), "UNIQUE constraint failed: "+field)
}
