package backend

type AMS struct {
	ID          int
	Name        string
	Country     string
	Description string
	Image       string
	Title       string
	Planets     []string
	OnFly       bool
}
