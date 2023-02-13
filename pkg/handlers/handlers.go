package handlers

import (
	"net/http"

	"github.com/ethanjmarchand/rewtwo/pkg/config"
	"github.com/ethanjmarchand/rewtwo/pkg/models"
	"github.com/ethanjmarchand/rewtwo/pkg/render"
)

// this is the repository patter. This allows a way to swap components
// Declare a new variable type *repository (Which is a struct with a *config.AppConfig)
var Repo *Repository

// Define the Repository type, that has another type of *confic.AppConfig insite of it.
type Repository struct {
	App *config.AppConfig
}

// This takes a memory address of config.AppConfig, and returns a *Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers takes a mem address of a type Repository, and assigns it to the var Repo (Of type *Repository)
func NewHandlers(r *Repository) {
	Repo = r
}

// Home - This is the handler function that is called for "/" route
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About - This is the handler function that is called for "/about" route
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := map[string]string{}
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
