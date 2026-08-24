package warehouse_enums

// ShapeType is the primitive shape used to render a LayoutNode when no
// 3D model file is supplied. CUSTOM means a model URL must be provided.
type ShapeType string

const (
	ShapeBox      ShapeType = "BOX"
	ShapeCylinder ShapeType = "CYLINDER"
	ShapeSphere   ShapeType = "SPHERE"
	ShapePlane    ShapeType = "PLANE"
	ShapeCustom   ShapeType = "CUSTOM"
)

// IsValid reports whether s is a known ShapeType.
func (s ShapeType) IsValid() bool {
	switch s {
	case ShapeBox, ShapeCylinder, ShapeSphere, ShapePlane, ShapeCustom:
		return true
	}
	return false
}

func (s ShapeType) String() string { return string(s) }
