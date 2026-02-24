package models
type Simulation struct{
	IDSim      string     `json:"id_sim"`      // maps to id_sim column
	IDCell     int     `json:"id_cell"`     // maps to id_cell column
    NitroKgHa  int `json:"nitro_kg_ha"` // maps to nitro_kg_ha
    YieldKgHa  float64 `json:"yield_kg_ha"` // maps to yield_kg_ha
}
