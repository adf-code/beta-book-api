package migration

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

// readSQLFile reads a file and returns its content as string
func readSQLFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", path, err)
	}
	return string(content), nil
}

// runSQL executes raw SQL from a file
func runSQL(db *sql.DB, path string) {
	sqlContent, err := readSQLFile(path)
	if err != nil {
		log.Fatalf("❌ Error reading SQL: %v", err)
	}

	_, err = db.Exec(sqlContent)
	if err != nil {
		log.Fatalf("❌ Error executing SQL in %s: %v", path, err)
	}
}

// extractVersion extracts the version (timestamp) from a file name
// Example: "20250725100000_add_books_table.up.sql" => "20250725100000"
func extractVersion(filePath string) string {
	base := filepath.Base(filePath)
	parts := strings.Split(base, "_")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

// getAppliedVersions reads versions from schema_migrations table
func getAppliedVersions(db *sql.DB) map[string]bool {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (
        version TEXT PRIMARY KEY,
        applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`)
	if err != nil {
		log.Fatalf("❌ Failed to ensure schema_migrations table: %v", err)
	}

	rows, err := db.Query("SELECT version FROM schema_migrations")
	if err != nil {
		log.Fatalf("❌ Failed to query schema_migrations: %v", err)
	}
	defer rows.Close()

	versions := make(map[string]bool)
	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			log.Fatalf("❌ Failed to scan migration version: %v", err)
		}
		versions[version] = true
	}

	return versions
}

// listSQLFiles returns sorted list of SQL files by extension (.up.sql or .down.sql)
func listSQLFiles(dir string, suffix string) []string {
	files, err := filepath.Glob(filepath.Join(dir, "*"+suffix))
	if err != nil {
		log.Fatalf("❌ Failed to list migration files: %v", err)
	}

	return files
}
