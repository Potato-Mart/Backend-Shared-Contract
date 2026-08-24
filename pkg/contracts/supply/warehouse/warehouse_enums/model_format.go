package warehouse_enums

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
