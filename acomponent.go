package goAframe

type AComponent interface {
	Attributes() (map[string]interface{}, error)
}
