package models
type Simulation struct{
	IDSim      string     `json:"id_sim"`      // maps to id_sim column
	IDCell     int     `json:"id_cell"`     // maps to id_cell column
    NitroKgHa  float64 `json:"nitro_kg_ha"` // Nitrogen in kg/ha
	NitroLbAc  float64 `json:"nitro_lb_ac"` // Nitrogen in lb/ac
    YieldKgHa  float64 `json:"yield_kg_ha"` // Yield in Kg/ha
	YieldBsAc  float64 `json:"yield_bs_ac"` // Yield in Bushes/ac
	NitroPrice float64 `json:"nitro_price"` 
	GrainPrice float64 `json:"grain_price"`	
	Profit_dol float64 `json:"profit_dol"`
}
