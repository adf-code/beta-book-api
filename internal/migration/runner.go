package migration

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"sort"
)

func MigrateUp(db *sql.DB, dir string) {
	applied := getAppliedVersions(db)
	files, _ := filepath.Glob(filepath.Join(dir, "*.up.sql"))
	sort.Strings(files)

	for _, f := range files {
		version := extractVersion(f)
		if applied[version] {
			continue
		}

		fmt.Println("🟢 Applying:", f)
		runSQL(db, f)
		_, _ = db.Exec("INSERT INTO schema_migrations (version) VALUES ($1)", version)
	}
}

func MigrateDown(db *sql.DB, dir string) {
	applied := getAppliedVersions(db)
	files, _ := filepath.Glob(filepath.Join(dir, "*.down.sql"))
	sort.Sort(sort.Reverse(sort.StringSlice(files)))

	for _, f := range files {
		version := extractVersion(f)
		if !applied[version] {
			continue
		}

		fmt.Println("🔻 Rolling back:", f)
		runSQL(db, f)
		_, _ = db.Exec("DELETE FROM schema_migrations WHERE version = $1", version)
		break
	}
}

// Fungsi tambahan: runSQL, extractVersion, getAppliedVersions bisa kamu pisah ke utils.go
