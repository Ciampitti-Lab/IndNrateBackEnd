package models
type SimulationAgg struct {
	IDCell     int     `json:"id_cell"`
	NitroLbAc  float64 `json:"nitro_lb_ac"`
	ProfitDol  float64 `json:"profit_dol"`
}

type Eonr struct {
    IDTrial string     `json:"id_trial"`
    Region  string  `json:"region"`
    EONR    float64 `json:"eonr"`
    Profit  float64 `json:"profit"`
}