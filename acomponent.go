package wasmGoAframe

type AComponent interface {
	Attributes() (map[string]interface{}, error)
}
