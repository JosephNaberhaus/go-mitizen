package commit

import (
	"encoding/json"
	"fmt"
	"github.com/JosephNaberhaus/go-mitizen/git"
	"github.com/JosephNaberhaus/go-mitizen/prompt"
	"github.com/JosephNaberhaus/go-mitizen/util"
	"io/ioutil"
	"log"
	"path/filepath"
)

const configName = "config.gz.json"

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

func overrideConfig() error {
	repositoryRoot, err := git.GetRepositoryRoot()
	if err != nil {
		return err
	}

	existed, err := overrideIfExists(filepath.Join(repositoryRoot, configName))
	if err != nil {
		return err
	}

	if !existed {
		_, err := overrideIfExists(filepath.Join("~", configName))
		return err
	}

	return nil
}

func overrideIfExists(configPath string) (existed bool, err error) {
	log.Printf("Looking for config at: %s", configPath)

	exists, err := util.FileExists(configPath)
	if err != nil {
		return false, fmt.Errorf("error checking if config file exists: %w", err)
	}

	if exists {
		log.Println("Config found")
		content, err := ioutil.ReadFile(configPath)
		if err != nil {
			return false, fmt.Errorf("error reading config at %s: %w", configPath, err)
		}

		err = json.Unmarshal(content, &config)
		if err != nil {
			return false, fmt.Errorf("error parsing config: %w", err)
		}

		return true, nil
	} else {
		log.Println("Config not found")
	}

	return false, nil
}