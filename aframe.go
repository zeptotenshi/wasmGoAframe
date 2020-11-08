package wasmGoAframe

import (
	"errors"
	"syscall/js"
)

type Aframe struct {
	global js.Value
	scene  js.Value
	three  js.Value

	entities map[string]*AEntity

	skyboxes map[string]int
}

var aframe *Aframe

func init() {
	aframe = &Aframe{
		entities: map[string]*AEntity{},
		skyboxes: map[string]int{},
	}

	global()
	scene()
	three()
}

func global() {
	a := js.Global().Get("AFRAME")
	if a.IsUndefined() || a.IsNull() {
		return
	}

	aframe.global = a
}

func scene() {
	doc := js.Global().Get("document")
	s := doc.Call("querySelector", "a-scene")
	if s.IsUndefined() || s.IsNull() {
		s = doc.Call("createElement", "a-scene")
		doc.Get("body").Call("appendChild", s)
	}

	aframe.scene = s
}

func three() {
	t := js.Global().Get("THREE")
	if t.IsUndefined() || t.IsNull() {
		return
	}

	aframe.three = t
}

// func (a *Aframe) CreateEntity() AEntity {
// }

func GetElementFromSceneById(_id string) (js.Value, error) {
	if e, ok := aframe.entities[_id]; ok {
		return e.Value(), nil
	}

	if aframe.global.IsNull() || aframe.global.IsUndefined() {
		global()
		if aframe.global.IsNull() || aframe.global.IsUndefined() {
			return js.ValueOf(nil), errors.New("AFRAME global not found")
		}
	}

	if aframe.scene.IsUndefined() || aframe.scene.IsNull() {
		scene()
		if aframe.scene.IsUndefined() || aframe.scene.IsNull() {
			return js.ValueOf(nil), errors.New("scene element not found")
		}
	}

	el := aframe.scene.Call("querySelector", "#"+_id)
	if el.IsUndefined() {
		return js.ValueOf(nil), errors.New(_id + " element undefined")
	} else if el.IsNull() {
		return js.ValueOf(nil), errors.New(_id + " element null")
	}

	return el, nil
}

func CacheEntity(e *AEntity) error {
	if _, ok := aframe.entities[e.ID()]; ok {
		return errors.New("entity already in cache")
	}

	aframe.entities[e.ID()] = e
	return nil
}

func GetCachedEntity(_id string) (*AEntity, error) {
	v, ok := aframe.entities[_id]
	if !ok {
		return nil, errors.New("entity with given id not found in cache")
	}

	return v, nil
}
