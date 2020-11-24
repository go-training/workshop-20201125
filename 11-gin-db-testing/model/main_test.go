package model

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/rs/zerolog/log"
)

var fixtures *testfixtures.Loader

func TestMain(m *testing.M) {
	MainTest(m, "..")
}

// MainTest a reusable TestMain(..) function for unit tests that need to use a
// test database. Creates the test database, and sets necessary settings.
func MainTest(m *testing.M, root string) {
	var err error
	if err := SetupDatabase(); err != nil {
		log.Fatal().Err(err).Msg("can't connect database")
	}
	fixturesDir := filepath.Join(root, "model", "fixtures")

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("can't initial database")
	}
	fixtures, err = testfixtures.New(
		testfixtures.Database(sqlDB),
		testfixtures.Dialect("sqlite"),
		testfixtures.Directory(fixturesDir),
		testfixtures.DangerousSkipTestDatabaseCheck(),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("can't setup database")
	}

	os.Exit(m.Run())
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		log.Fatal().Err(err).Msg("can't load fixtures")
	}
}
