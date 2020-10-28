package goAframe

import (
	"syscall/js"
)

type Entity interface {
	SetAttribute(string, map[string]interface{})
	GetProperty(string) (js.Value, error)
	GetAttribute(string) (js.Value, error)
	ID() string
	Tag() string
	Value() js.Value
	String() string
}

type AEntity struct {
	Entity
}

func (e *AEntity) SetPosition(_x, _y, _z float64) error {
	obj3D := v.Value().Get("object3D")
	if obj3D.IsUndefined() {
		return errors.New("entity.object3D is undefined")
	}
	if obj3D.IsNull() {
		return errors.New("entity.object3D is null")
	}

	pos := obj3D.Get("position")
	if pos.IsUndefined() {
		return errors.New("entity.object3D.position is undefined")
	}
	if pos.IsNull() {
		return errors.New("entity.object3D.position is null")
	}

	pos.Set("x", _x)
	pos.Set("y", _y)
	pos.Set("z", _z)

	return nil
}