package models
var (
	DimensionSencillaList map[string]*DimensionSencilla
)
func init() {
	DimensionSencillaList = make(map[string]*DimensionSencilla)
}
type DimensionSencilla struct {
	ANL_CODE string
	NAME string
}
func ClearDimensionSencillaList() {
	DimensionSencillaList = make(map[string]*DimensionSencilla)
}
func AddDimensionSencilla(u DimensionSencilla, index string)  {
	DimensionSencillaList[index] = &u	
}
func GetAllDimensionSencilla() map[string]*DimensionSencilla {
	return DimensionSencillaList
}
