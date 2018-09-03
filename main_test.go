package categorizer

import (
	"github.com/bugfixes/go-bugfixes"
	"github.com/joho/godotenv"
	"testing"
)

func TestGetTitles(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		bugfixes.Error("Env - GetTitles", err)
	}

	GetTitles()
}

func TestParseTitles(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		bugfixes.Error("Env - ParseTitles", err)
	}

	AllTitles = AllItems{}
	ParseTitles()
}

func TestRunApp(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		bugfixes.Error("Env - RunApp", err)
	}

	RunApp()
}
