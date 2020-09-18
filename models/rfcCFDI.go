package models
var (
	RFCCFDIList map[string]*RFCCFDI
)
func init() {
	RFCCFDIList = make(map[string]*RFCCFDI)
}
type CFDIDatos struct {
	XML string
	UUID string
	Total float64
	SubTotal float64
	TotalImpuestosRetenidos float64
	TotalImpuestosTrasladados float64
	IVARetenido float64
	ISRRetenido float64
	Regimen string
	Fecha string
}
type RFCCFDI struct {
	RFC string
	RazonSocial string
	Total float64
	SubTotal float64
	TotalImpuestosRetenidos float64
	TotalImpuestosTrasladados float64
	CFDIDatosList  map[string]*CFDIDatos
}
func DameRFCCFDI(index string) RFCCFDI {
	return *RFCCFDIList[index]
}
func AddRFCCFDI(u RFCCFDI, index string)  {
	RFCCFDIList[index] = &u	
}
func GetAllRFCCFDI() map[string]*RFCCFDI {
	return RFCCFDIList
}
func ClearRFCCFDI() {
	RFCCFDIList = make(map[string]*RFCCFDI)	
}