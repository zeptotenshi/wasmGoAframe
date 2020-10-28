package goAframe

import (
	"errors"
	"syscall/js"
)

type Aframe struct {
	js.Value
	scene    js.Value
	entities map[string]*AEntity
}

var aframe *Aframe

func init() {
	Aframe := &Aframe{
		entities: map[string]AEntity{},
	}

	aframe()
	scene()
}

func aframe() {
	a := js.Global().Get("AFRAME")
	if a.IsUndefined() || s.IsNull() {
		return
	}

	aframe.Value = a
}

func scene() {
	doc := js.Global().Get("document")
	s := doc.Call("querySelector", "a-scene")
	if s.IsUndefined() || s.IsNull() {
		s = doc.Call("createElement", "a-scene")
		doc.Get("body").Call("appendChild", s)
	}

	Aframe.scene = s
}

// func (a *Aframe) CreateEntity() AEntity {
// }

func GetEntityFromSceneByID(_id string) (js.Value, error) {
	if e, ok := aframe.entities[_id]; ok {
		return e.Value(), nil
	}

	el := aframe.scene.Call("querySelector", "#"+_id)
	if el.IsUndefined() {
		return js.ValueOf(nil), errors.New(_id + " element undefined")
	} else if el.IsNull() {
		return js.ValueOf(nil), errors.New(_id + " element null")
	}

	return el, nil
}

func CacheEntity(e AEntity) error {
	if _, ok := aframe.entities[e.ID()]; ok {
		return errors.New("entity already in cache")
	}

	a.entities[e.ID()] = e
	return nil
}

func GetCachedEntity(_id string) (AEntity, error) {
	v, ok := aframe.entities[_id]
	if !ok {
		return nil, errors.New("entity with given id not found in cache")
	}

	return v, nil
}
