package utils

import (
	"fmt"

	"github.com/Marchiz3D/gwho/config"
)

func GetAlias() ([]string, error) {
	conf, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	if len(conf.Profiles) == 0 {
		return nil, fmt.Errorf("No Git profiles found. Use 'gwho add <alias>' to create one.")
	}

	aliases := make([]string, 0, len(conf.Profiles))
	for alias := range conf.Profiles {
		aliases = append(aliases, alias)
	}

	return aliases, nil
}
