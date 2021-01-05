package commit

import (
	"github.com/JosephNaberhaus/go-mitizen/prompt"
)

type Config struct {
	ForceSubjectLowerCase   bool
	ForceScopeLowerCase     bool
	AllowBlankLinesInBody   bool
	MaxHeaderLength         int
	MaxLineLength           int
	Types                   []*prompt.SelectionOption
}

var config = Config{
	ForceSubjectLowerCase: true,
	ForceScopeLowerCase: true,
	AllowBlankLinesInBody: true,
	MaxHeaderLength: 100,
	MaxLineLength: 100,
	Types: []*prompt.SelectionOption{
		{
			Name:        "feat",
			Description: "A new feature",
		},
		{
			Name:        "fix",
			Description: "A bug fix",
		},
		{
			Name:        "improvement",
			Description: "An improvement to a current feature",
		},
		{
			Name:        "docs",
			Description: "Documentation only changes",
		},
		{
			Name:        "style",
			Description: "Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)",
		},
		{
			Name:        "refactor",
			Description: "A code change that neither fixes a bug nor adds a feature",
		},
		{
			Name:        "perf",
			Description: "A code change that improves performance",
		},
		{
			Name:        "test",
			Description: "Adding missing tests or correcting existing tests",
		},
		{
			Name:        "build",
			Description: "Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)",
		},
		{
			Name:        "ci",
			Description: "Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)",
		},
		{
			Name:        "chore",
			Description: "Other changes that don't modify src or test files",
		},
		{
			Name:        "revert",
			Description: "Reverts a previous commit",
		},
	},
}