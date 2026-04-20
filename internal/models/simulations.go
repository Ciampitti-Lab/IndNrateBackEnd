package models
type Simulation struct{
	IDCell       int     `json:"id_cell"`     // maps to id_cell column
	IDWithinCell int     `json:"id_within_cell"`     // maps to id_within_cell column
	Year         int     `json:"year"`     // maps to years column
    NitroKgHa  float64 `json:"nitro_kg_ha"` // Nitrogen in kg/ha
	NitroLbAc  float64 `json:"nitro_lb_ac"` // Nitrogen in lb/ac
    YieldKgHa  float64 `json:"yield_kg_ha"` // Yield in Kg/ha
	YieldBsAc  float64 `json:"yield_bs_ac"` // Yield in Bushes/ac
	NitroPrice float64 `json:"nitro_price"` 
	GrainPrice float64 `json:"grain_price"`	
	Profit_dol float64 `json:"profit_dol"`
}

type Eonr struct {
    IDTrial string     `json:"id_trial"`
    Region  string  `json:"region"`
    EONR    float64 `json:"eonr"`
    Profit  float64 `json:"profit"`
}