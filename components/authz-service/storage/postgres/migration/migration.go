package migration

import (
	"context"
	"database/sql"
	"net/url"

	"github.com/chef/automate/lib/logger"
	"github.com/golang-migrate/migrate"
	"github.com/pkg/errors"

	"github.com/chef/automate/components/authz-service/storage/postgres/datamigration"
)

// Config holds the information needed to connect to the database (PGURL), to
// find the migration SQL files (Path), and log debug messages (Logger).
type Config struct {
	PGURL  *url.URL
	Path   string
	Logger logger.Logger
}

const (
	// PreForceUpgradeMigration is last schema migration before force-upgrade.
	// All of our golang force-upgrade code assumes we are on this schema version.
	PreForceUpgradeMigration = 73
)

// Migrate executes all migrations we have
func (c *Config) Migrate(dataMigConf datamigration.Config) error {
	pgURL := c.PGURL.String()
	migrationsPath := c.Path
	l := c.Logger
	migrationsTable := ""
	ctx := context.Background()

	l.Infof("Running db migrations from %q", migrationsPath)
	purl, err := addMigrationsTable(pgURL, migrationsTable)
	if err != nil {
		return errors.Wrap(err, "parse PG URL")
	}
	m, err := migrate.New(addScheme(migrationsPath), purl)
	if err != nil {
		return errors.Wrap(err, "init migrator")
	}

	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return errors.Wrap(err, "init migrator - error getting migration version")
	}

	if dirty {
		// force to prior version to reattempt migration
		err := m.Force(int(version) - 1)
		if err != nil {
			return errors.Wrap(err, "force to working schema version")
		}
		l.Infof("Forced to previous version: %v to reattempt migration", int(version)-1)
	} else {
		l.Infof("Current schema version: %v", version)
	}

	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return errors.Wrap(err, "opening database connection")
	}

	if version < PreForceUpgradeMigration {
		err = m.Migrate(PreForceUpgradeMigration)
		if err != nil && err != migrate.ErrNoChange {
			return errors.Wrap(err, "migration up to IAM V2 force upgrade failed")
		}
	}

	notOnV2, isDirty, err := migrateFromScratch(ctx, db)
	if err != nil {
		return errors.Wrap(err, "failed to retrieve migration_status")
	}
	if notOnV2 {
		if isDirty {
			// get IAM db in original, clean state to avoid conflicts
			err = resetIAMDb(ctx, db)
			if err != nil {
				return errors.Wrap(err, "reset IAM V2 database")
			}
			if err := dataMigConf.Reset(); err != nil {
				return errors.Wrap(err, "reset v2 data migrations")
			}
		}
		err = recordMigrationStatus(ctx, enumInProgress, db)
		if err != nil {
			return errors.Wrapf(err, "failed to set IAM v2 migration_status to %s", enumInProgress)
		}
		err = migrateToV2(ctx, db)
		if err != nil {
			statusErr := recordMigrationStatus(ctx, enumFailed, db)
			if err != nil {
				return errors.Wrapf(statusErr, "failed to set IAM v2 migration_status to %s:%s", enumFailed, err.Error())
			}
			return errors.Wrap(err, "IAM v2 force-upgrade failed")
		}
		err = recordMigrationStatus(ctx, enumSuccessful, db)
		if err != nil {
			return errors.Wrapf(err, "failed to set IAM v2 migration_status to %s", enumSuccessful)
		}
	}

	// idempotent
	err = dataMigConf.Migrate()
	if err != nil {
		return errors.Wrap(err, "IAM data migrations failed")
	}

	// perform remaining migrations
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "migrations failed")
	}

	l.Infof("Completed db migrations")

	err = db.Close()
	if err != nil {
		return errors.Wrap(err, "close migration db connection")
	}

	// The first error is trying to Close() the source. For our file source,
	// that's always nil
	_, err = m.Close()
	return errors.Wrap(err, "close migrations connection")
}

func addMigrationsTable(u, table string) (string, error) {
	pgURL, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	if table != "" {
		q := pgURL.Query()
		q.Set("x-migrations-table", table)
		pgURL.RawQuery = q.Encode()
	}
	return pgURL.String(), nil
}

func addScheme(p string) string {
	u := url.URL{}
	u.Scheme = "file"
	u.Path = p
	return u.String()
}
