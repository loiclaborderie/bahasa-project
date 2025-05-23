package module

type ModuleCreateRequest struct {
	Title       string `validate:"required" json:"title"`
	Description string `json:"description"`
	Visible     bool   `validate:"boolean" json:"visible"`
	CoverURL    string `json:"cover_url"`
}
