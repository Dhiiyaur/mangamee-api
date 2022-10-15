package helper

import "mangamee-api/src/config"

func IsProductionEnv(config config.Config) bool {
	return config.Server.Env == "PROD"
}
