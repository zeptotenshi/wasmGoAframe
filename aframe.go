package aframe

import (
	"errors"
	"fmt"
	"syscall/js"
)

type AEntity interface {
	SetAttribute(string, map[string]interface{})
	GetProperty(string) (js.Value, error)
	GetAttribute(string) (js.Value, error)
	ID() string
	Tag() string
	Value() js.Value
	String() string
}

type AComponent interface {
	Attributes() (map[string]interface{}, error)
}

type Aframe struct {
	aframe   js.Value
	scene    js.Value
	entities map[string]AEntity
}

func NewScene() *Aframe {
	r := &Aframe{
		entities: map[string]AEntity{},
	}

	doc := js.Global().Get("document")
	s := doc.Call("querySelector", "a-scene")
	if s.IsUndefined() || s.IsNull() {
		s = doc.Call("createElement", "a-scene")
		doc.Get("body").Call("appendChild", s)
	}
	r.scene = s

	a := js.Global().Get("AFRAME")
	if a.IsUndefined() || s.IsNull() {
		fmt.Printf("global variable 'AFRAME' not found")
		return r
	}
	r.aframe = a

	return r
}

// func (a *Aframe) CreateEntity() AEntity {
// }

func (a *Aframe) GetEntityFromSceneByID(_id string) (js.Value, error) {
	if e, ok := a.entities[_id]; ok {
		return e.Value(), nil
	}

	el := a.scene.Call("querySelector", "#"+_id)
	if el.IsUndefined() {
		return js.ValueOf(nil), errors.New(_id + " element undefined")
	} else if el.IsNull() {
		return js.ValueOf(nil), errors.New(_id + " element null")
	}

	return el, nil
}

func (a *Aframe) CacheEntity(e AEntity) error {
	if _, ok := a.entities[e.ID()]; ok {
		return errors.New("entity already in cache")
	}

	a.entities[e.ID()] = e
	return nil
}

func (a *Aframe) RegisterComponent() error {
	if a.aframe.IsUndefined() || a.aframe.IsNull() {
		return errors.New("'AFRAME' global var not found")
	}

	a.aframe.Call("registerComponent", map[string]interface{}{
		"hello": "world",
	})
	return nil
}

func (a *Aframe) Aframe() js.Value {
	return a.aframe
}

func (a *Aframe) GetCachedEntity(_id string) (AEntity, error) {
	v, ok := a.entities[_id]
	if !ok {
		return nil, errors.New("entity with given id not found in cache")
	}

	return v, nil
}

func (a *Aframe) SetPosition(_id string, _x, _y, _z float64) error {
	v, ok := a.entities[_id]
	if !ok {
		return errors.New("entity with given id not found in scene cache")
	}

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

// func (a *Aframe) RemoveEntity(v js.Value) error {
// 	if v.IsUndefined() || v.IsNull() {
// 		return nil
// 	}

// 	n := v.Get("id")
// 	if !n.IsUndefined() && !n.IsNull() {
// 		if n.String() != "" {
// 			if _, ok := a.entities[n.String()]; ok {
// 				delete(a.entities, n.String())
// 			}
// 		}
// 	}
// 	a.scene.Call("removeChild", v)
// 	return nil
// }

// func (a *Aframe) getEntity(id string) (e *Entity, error) {

// }

// func (a *Aframe) CacheEntity(e *Entity) {
// 	a.entities[e.ID] = e
// }

// func (a *Aframe) addEntityToScene(id string, e Entity) {
// 	a.entities[id] = e
// }

// func (a Aframe) registerComponent(name string, fn func()) {

// }
