package binmsg

const gpsOctets = 13

type Gps struct {
	FixMode   FixMode `json:"fixmode"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Altitude  float32 `json:"altitude"`
}

func (g *Gps) Alt() float64 {
	return float64(g.Altitude)
}

func (g *Gps) Lat() float64 {
	return float64(g.Latitude)
}

func (g *Gps) Lon() float64 {
	return float64(g.Longitude)
}

func (g *Gps) SetAlt(v float64) {
	g.Altitude = float32(v)
}

func (g *Gps) SetLat(v float64) {
	g.Latitude = float32(v)
}

func (g *Gps) SetLon(v float64) {
	g.Longitude = float32(v)
}
