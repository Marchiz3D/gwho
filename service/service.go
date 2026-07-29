package service

import (
	"fmt"
	"os/exec"

	"github.com/Marchiz3D/gwho/config"
)

func UseProfile(alias string, isGlobal bool) (config.Profile, error) {
	profiles, err := GetAllProfiles()
	if err != nil {
		return config.Profile{}, err
	}

	profile, ok := profiles[alias]
	if !ok {
		return config.Profile{}, fmt.Errorf("Profile '%s' not found", alias)
	}

	var err1, err2 error
	if isGlobal {
		err1 = exec.Command("git", "config", "--global", "user.name", profile.Name).Run()
		err2 = exec.Command("git", "config", "--global", "user.email", profile.Email).Run()
	} else {
		err1 = exec.Command("git", "config", "user.name", profile.Name).Run()
		err2 = exec.Command("git", "config", "user.email", profile.Email).Run()
	}
	if err1 != nil || err2 != nil {
		return config.Profile{}, fmt.Errorf("Error applying profile. Make sure you are in a git repository")
	}
	return profile, nil

}

func GetAllProfiles() (map[string]config.Profile, error) {
	conf, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	if len(conf.Profiles) == 0 {
		return nil, fmt.Errorf("No Git profiles found. Use 'gwho add <alias>' to create one.")
	}

	return conf.Profiles, nil
}

func CreateProfile(alias string, keys config.Profile) error {
	conf, err := config.LoadConfig()
	if err != nil {
		return err
	}

	if _, ok := conf.Profiles[alias]; ok {
		return fmt.Errorf("Alias '%s' already exists", alias)
	}

	conf.Profiles[alias] = config.Profile{
		Name:  keys.Name,
		Email: keys.Email,
	}

	if err := conf.SaveConfig(); err != nil {
		return err
	}

	return nil
}

func UpdateProfile(alias string, keys config.Profile) error {
	conf, err := config.LoadConfig()
	if err != nil {
		return err
	}

	if _, ok := conf.Profiles[alias]; !ok {
		return fmt.Errorf("Alias '%s' does not exist", alias)
	}

	conf.Profiles[alias] = config.Profile{
		Name:  keys.Name,
		Email: keys.Email,
	}

	if err := conf.SaveConfig(); err != nil {
		return err
	}

	return nil
}

func DeleteProfile(alias string) error {
	conf, err := config.LoadConfig()
	if err != nil {
		return err
	}

	if _, ok := conf.Profiles[alias]; !ok {
		return fmt.Errorf("Alias '%s' does not exist", alias)
	}

	delete(conf.Profiles, alias)

	if err := conf.SaveConfig(); err != nil {
		return err
	}

	return nil
}
