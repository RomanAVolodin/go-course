package handlers

import (
	"github.com/RomanAVolodin/go-course/pkg/config"
	"github.com/RomanAVolodin/go-course/pkg/render"
	"net/http"
)

// Repo the repository useed by the handlers
var Repo *Repository

// Repository is the repository tyoe
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository {
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the bout page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the bout page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
