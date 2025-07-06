package testutil

import (
	"path/filepath"
	"testing"

	"github.com/ajohnston1219/eatme/api/internal/db"
)

func NewTestSQLiteStore(t *testing.T) *db.SQLiteStore {
	t.Helper()

	tempDir := t.TempDir()
	dsn := "file:" + filepath.Join(tempDir, "test.db")

	store, err := db.NewSQLiteStore(dsn)
	if err != nil {
		t.Fatalf("failed to init SQLiteStore: %v", err)
	}
	return store
}
