package db

import (
	"database/sql"
	"encoding/json"
	"errors"
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
		return nil, err
	}
	if _, err := db.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		return nil, err
	}
	if err := migrate(db); err != nil {
		return nil, err
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
		return models.User{}, err
	}
	return models.User{ID: id, Email: email}, nil
}

func (s *SQLiteStore) SaveProfile(userID string, p models.Profile) error {
	cuisines, _ := json.Marshal(p.Cuisines)
	diet, _ := json.Marshal(p.Diet)
	equipment, _ := json.Marshal(p.Equipment)
	allergies, _ := json.Marshal(p.Allergies)

	_, err := s.Exec(`
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
		  allergies  = excluded.allergies;
	`, userID, p.SetupStep, p.Name, p.Skill, cuisines, diet, equipment, allergies)
	return err
}

func (s *SQLiteStore) GetProfile(userID string) (models.Profile, error) {
	var p models.Profile
	var cuisines, diet, equipment, allergies []byte

	err := s.QueryRow(`
		SELECT setup_step, name, skill, cuisines, diet, equipment, allergies
		FROM profiles WHERE user_id = ?;
	`, userID).Scan(
		&p.SetupStep, &p.Name, &p.Skill,
		&cuisines, &diet, &equipment, &allergies,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return p, ErrNotFound
		}
		return p, err
	}

	_ = json.Unmarshal(cuisines, &p.Cuisines)
	_ = json.Unmarshal(diet, &p.Diet)
	_ = json.Unmarshal(equipment, &p.Equipment)
	_ = json.Unmarshal(allergies, &p.Allergies)
	return p, nil
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

	if _, err := db.Exec(users); err != nil {
		return err
	}
	if _, err := db.Exec(profiles); err != nil {
		return err
	}
	return nil
}

func isUniqueViolation(err error, field string) bool {
	return err != nil && strings.Contains(err.Error(), "UNIQUE constraint failed: "+field)
}
