package database

import (
	"MusicFinder/src/database"
	"strconv"
	"testing"
)

func TestGetDatabaseHandler(t *testing.T) {
	handler := database.GetDatabaseHandler()
	if handler == nil {
		t.Errorf("GetDatabaseHandler() = nil, want Database.handler")
	}
}

func TestGetAlreadyCreatedDatabaseHandlerWith(t *testing.T) {
	handler := database.GetDatabaseHandler()
	handler = database.GetDatabaseHandler()
	if handler == nil {
		t.Errorf("GetDatabaseHandler() = nil, want Database.handler")
	}
}

func TestCreateDbHandler(t *testing.T) {
	if got, err := database.CreateDbHandler(); got != true && err == nil {
		t.Errorf("CreateDbHandler() = %q, want %q", strconv.FormatBool(got), strconv.FormatBool(true))
	}
}
