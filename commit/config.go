package commit

import (
	"encoding/json"
	"github.com/JosephNaberhaus/go-mitizen/prompt"
)

const defaultConfig = `{
  "forceSubjectLowerCase": true,
  "forceScopeLowerCase": true,
  "allowBlankLinesInBody": true,
  "allowBlankLinesInFooter": false,
  "maxHeaderLength": 100,
  "maxLineLength": 100,
  "types": [
    {
      "name": "feat",
      "description": "A new feature"
    },
    {
      "name": "fix",
      "description": "A bug fix"
    },
    {
      "name": "improvement",
      "description": "An improvement to a current feature"
    },
    {
      "name": "docs",
      "description": "Documentation only changes"
    },
    {
      "name": "style",
      "description": "Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)"
    },
    {
      "name": "refactor",
      "description": "A code change that neither fixes a bug nor adds a feature"
    },
    {
      "name": "perf",
      "description": "A code change that improves performance"
    },
    {
      "name": "test",
      "description": "Adding missing tests or correcting existing tests"
    },
    {
      "name": "build",
      "description": "Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)"
    },
    {
      "name": "ci",
      "description": "Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)"
    },
    {
      "name": "chore",
      "description": "Other changes that don't modify src or test files"
    },
    {
      "name": "revert",
      "description": "Reverts a previous commit"
    }
  ]
}`

type Config struct {
	ForceSubjectLowerCase   bool
	ForceScopeLowerCase     bool
	AllowBlankLinesInBody   bool
	AllowBlankLinesInFooter bool
	MaxHeaderLength         int
	MaxLineLength           int
	Types                   []*prompt.SelectionOption
}

func loadDefaultConfig() *Config {
	config := new(Config)

	err := json.Unmarshal([]byte(defaultConfig), &config)
	if err != nil {
		panic(err)
	}

	return config
}