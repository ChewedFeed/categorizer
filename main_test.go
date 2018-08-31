package categorizer

import (
  "testing"
  "github.com/bugfixes/go-bugfixes"
  "github.com/joho/godotenv"
)

func TestGetTitles(t *testing.T) {
  err := godotenv.Load()
  if err != nil {
    bugfixes.Error("Env", err)
  }

  GetTitles()
}
