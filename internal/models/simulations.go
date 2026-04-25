package models

import "time"
type Simulation struct {
	NitroLbAc float64 `json:"nitro_lb_ac"`
	ProfitDol float64 `json:"profit_dol"`
}

type Eonr struct {
    IDTrial string     `json:"id_trial"`
    Region  string  `json:"region"`
    EONR    float64 `json:"eonr"`
    Profit  float64 `json:"profit"`
}

type NitroPrice struct {
	Date         time.Time `json:"date"`
	NitroSource  string    `json:"nitro_source"`
	NitroPriceLb float64   `json:"nitro_price_lb"`
}

type CornPrice struct {
	Date         time.Time `json:"date"`
	CornPriceLb float64   `json:"corn_price_lb"`
}