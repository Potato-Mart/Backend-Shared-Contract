package warehouse_enums

// LayoutNodeType is the level of a node in the warehouse 3D hierarchy.
// The tree is conceptually Zone > Aisle > Rack > Shelf > Bin, but a
// concrete depot may skip levels (e.g., go straight from Aisle to Bin).
type LayoutNodeType string

const (
	LayoutNodeZone  LayoutNodeType = "ZONE"
	LayoutNodeAisle LayoutNodeType = "AISLE"
	LayoutNodeRack  LayoutNodeType = "RACK"
	LayoutNodeShelf LayoutNodeType = "SHELF"
	LayoutNodeBin   LayoutNodeType = "BIN"
)

// IsValid reports whether t is a known LayoutNodeType.
func (t LayoutNodeType) IsValid() bool {
	switch t {
	case LayoutNodeZone, LayoutNodeAisle, LayoutNodeRack,
		LayoutNodeShelf, LayoutNodeBin:
		return true
	}
	return false
}

func (t LayoutNodeType) String() string { return string(t) }

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

// ModelFormat is the file format of a 3D asset referenced by URL. GLB
// (binary glTF) is the recommended default for web viewers because it
// loads in one request.
type ModelFormat string

const (
	ModelFormatGLB  ModelFormat = "GLB"
	ModelFormatGLTF ModelFormat = "GLTF"
	ModelFormatOBJ  ModelFormat = "OBJ"
	ModelFormatFBX  ModelFormat = "FBX"
	ModelFormatUSDZ ModelFormat = "USDZ"
)

// IsValid reports whether m is a known ModelFormat.
func (m ModelFormat) IsValid() bool {
	switch m {
	case ModelFormatGLB, ModelFormatGLTF, ModelFormatOBJ,
		ModelFormatFBX, ModelFormatUSDZ:
		return true
	}
	return false
}

func (m ModelFormat) String() string { return string(m) }

// CameraProjection selects perspective or orthographic rendering for a
// saved camera preset.
type CameraProjection string

const (
	CameraPerspective  CameraProjection = "perspective"
	CameraOrthographic CameraProjection = "orthographic"
)

// IsValid reports whether p is a known CameraProjection.
func (p CameraProjection) IsValid() bool {
	switch p {
	case CameraPerspective, CameraOrthographic:
		return true
	}
	return false
}

func (p CameraProjection) String() string { return string(p) }
