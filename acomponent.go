package wasmGoAframe

type Component interface {
	Attributes() (map[string]interface{}, error)
}

type AComponent struct {
	Component
}
