package wasmGoAframe

import (
	"errors"
	"syscall/js"
)

var (
	faces []string

	THREE_BackSide          js.Value
	THREE_BoxGeometry       js.Value
	THREE_TextureLoader     js.Value
	THREE_Mesh              js.Value
	THREE_MeshBasicMaterial js.Value
)

func init() {
	faces = []string{"back", "front", "bottom", "top", "left", "right"}

	THREE_BackSide = aframe.three.Get("BackSide")
	THREE_Mesh = aframe.three.Get("Mesh")
	THREE_TextureLoader = aframe.three.Get("TextureLoader")
	THREE_MeshBasicMaterial = aframe.three.Get("MeshBasicMaterial")
	THREE_BoxGeometry = aframe.three.Get("BoxGeometry")
}

type Skybox struct {
	Name string
	uuid int

	Images map[string]string

	Length float32
	Height float32
	Depth  float32
}

func newTexture(_src string) (js.Value, error) {
	texture := js.ValueOf(nil)

	if THREE_TextureLoader.IsUndefined() {
		return texture, errors.New("'THREE_TextureLoader' js.Value undefined")
	}
	if THREE_TextureLoader.IsNull() {
		return texture, errors.New("'THREE_TextureLoader' js.Value null")
	}

	// create a new TextureLoader
	textureLoader := THREE_TextureLoader.New()

	// var err error

	// texture = textureLoader.Call("load", _src, js.FuncOf(nil), js.FuncOf(nil), js.FuncOf(func(_this js.Value, _args []js.Value) interface{} {
	// 	wp.LogValue(_args[0])
	// 	err = errors.New(_args[0].String())

	// 	return js.ValueOf(nil)
	// }))
	texture = textureLoader.Call("load", _src)
	if texture.IsUndefined() {
		return texture, errors.New("'texture' js.Value undefined")
	}
	if texture.IsNull() {
		return texture, errors.New("'texture' js.Value null")
	}

	return texture, nil
}

func newBasicMaterial(_texture js.Value) (js.Value, error) {
	material := js.ValueOf(nil)

	if THREE_MeshBasicMaterial.IsUndefined() {
		return material, errors.New("'THREE_MeshBasicMaterial' js.Value undefined")
	}
	if THREE_MeshBasicMaterial.IsNull() {
		return material, errors.New("'THREE_MeshBasicMaterial' js.Value null")
	}

	// create a new MeshBasicMaterial
	material = THREE_MeshBasicMaterial.New(map[string]interface{}{"map": _texture})
	if material.IsUndefined() {
		return material, errors.New("material js.Value undefined")
	}
	if material.IsNull() {
		return material, errors.New("material js.Value null")
	}

	return material, nil
}

func (s *Skybox) SetAsSceneSkybox() error {
	for attempts := 3; attempts >= 0; attempts-- {
		if aframe.scene.IsUndefined() || aframe.scene.IsNull() {
			if attempts == 0 {
				return errors.New("'AFRAME.scene' global js.Value undefined/null")
			}
			scene()
			continue
		}

		break
	}

	for attempts := 3; attempts >= 0; attempts-- {
		if aframe.three.IsUndefined() || aframe.three.IsNull() {
			if attempts == 0 {
				return errors.New("'THREE' global js.Value undefined/null")
			}
			three()
			continue
		}

		break
	}

	if THREE_BackSide.IsUndefined() || THREE_BackSide.IsNull() {
		return errors.New("'THREE_BackSide' js.Value undefined/null")
	}

	if THREE_Mesh.IsUndefined() || THREE_Mesh.IsNull() {
		return errors.New("'THREE_Mesh' js.Value undefined/null")
	}

	if THREE_BoxGeometry.IsUndefined() || THREE_BoxGeometry.IsNull() {
		return errors.New("'THREE_BoxGeometry' js.Value undefined/null")
	}

	if l := len(s.Images); l != 6 {
		return errors.New("skybox image map size does not equal 6")
	}

	materialArray := make([]interface{}, 6)

	for i, fn := range faces {
		v, ok := s.Images[fn]
		if !ok {
			return errors.New("face['" + fn + "'] not found in skybox image map")
		}

		// texture for the image src (v)
		texture, err := newTexture(v)
		if err != nil {
			return errors.New("failed to create texture: " + err.Error())
		}

		// material for the new texture
		material, err := newBasicMaterial(texture)
		if err != nil {
			return errors.New("failed to create material: " + err.Error())
		}
		material.Set("side", THREE_BackSide)

		materialArray[i] = material
	}

	skyboxGeo := THREE_BoxGeometry.New(s.Length, s.Height, s.Depth)
	if skyboxGeo.IsUndefined() {
		return errors.New("'skyboxGeometry' js.Value undefined")
	}
	if skyboxGeo.IsNull() {
		return errors.New("'skyboxGeometry' js.Value null")
	}

	// new THREE.Mesh(THREE.Geometry, []THREE.Material)
	skyboxMesh := THREE_Mesh.New(skyboxGeo, materialArray)
	if skyboxMesh.IsUndefined() {
		return errors.New("'skyboxMesh' js.Value undefined")
	}
	if skyboxMesh.IsNull() {
		return errors.New("'skyboxMesh' js.Value null")
	}

	id := skyboxMesh.Get("id")
	if id.IsUndefined() || id.IsNull() {
		return errors.New("skyboxMesh id js.Value undefined/null")
	}

	sceneObj := aframe.scene.Get("object3D")
	if sceneObj.IsUndefined() || sceneObj.IsNull() {
		return errors.New("'scene.object3D' js.Value undefined/null")
	}

	sceneObj.Call("add", skyboxMesh)

	s.uuid = id.Int()
	aframe.skyboxes[s.Name] = id.Int()

	return nil
}
