package config

import (
	"encoding/json"
	"fmt"
	"github.com/JosephNaberhaus/go-mitizen/git"
	"github.com/JosephNaberhaus/go-mitizen/util"
	"github.com/JosephNaberhaus/prompt"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Config struct {
	ForceSubjectLowerCase *bool
	ForceScopeLowerCase   *bool
	AllowBlankLinesInBody *bool
	MaxHeaderLength       *int
	MaxLineLength         *int
	Types                 []prompt.SelectionOption
}

func OverrideConfig() error {
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

		config := Config{}
		err = json.Unmarshal(content, &config)
		if err != nil {
			return false, fmt.Errorf("error parsing config: %w", err)
		}

		applyConfig(config)

		return true, nil
	} else {
		log.Println("Config not found")
	}

	return false, nil
}

func applyConfig(config Config) {
	if config.ForceScopeLowerCase != nil {
		ForceSubjectLowerCase = *config.ForceSubjectLowerCase
	}

	if config.ForceScopeLowerCase != nil {
		ForceScopeLowerCase = *config.ForceScopeLowerCase
	}

	if config.AllowBlankLinesInBody != nil {
		AllowBlankLinesInBody = *config.AllowBlankLinesInBody
	}

	if config.MaxHeaderLength != nil {
		MaxHeaderLength = *config.MaxHeaderLength
	}

	if config.MaxLineLength != nil {
		MaxLineLength = *config.MaxLineLength
	}

	if config.Types != nil {
		Types = config.Types
	}
}
