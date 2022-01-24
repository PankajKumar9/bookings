package config

//IMPORT ONLY WHAT IS ABSOLUTELY NECESSARY TO AVOID IMPORT CYCLE

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

//AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
