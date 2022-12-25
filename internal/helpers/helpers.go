package helpers

import "github.com/ShehabEl-DeenAlalkamy/residencia/internal/config"

var app *config.AppConfig

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}
