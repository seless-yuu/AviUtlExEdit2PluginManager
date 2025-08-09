package library

// Plugin represents the metadata for a single AviUtl plugin.
type Plugin struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Files       []string `json:"files"`
}
