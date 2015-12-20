package binmsg

import "github.com/larsth/go-gpsfix"

const gpsOctets = 13

//Gps is type that contain a FixMode and the 32-bit floating point values:
//Latitude, Longitude, and Altitude.
type Gps struct {
	FixMode   gpsfix.FixMode `json:"fixmode"`
	Latitude  float32        `json:"latitude"`
	Longitude float32        `json:"longitude"`
	Altitude  float32        `json:"altitude"`
}

//Alt return a float64 representation of the Altitude.
func (g *Gps) Alt() float64 {
	return float64(g.Altitude)
}

//Lat return a float64 representation of the Latitude.
func (g *Gps) Lat() float64 {
	return float64(g.Latitude)
}

//Lon return a float64 representation of the Longitude.
func (g *Gps) Lon() float64 {
	return float64(g.Longitude)
}

//SetAlt sets the 32-bit floating point altitude value via a 64-bit floating
//point value.
func (g *Gps) SetAlt(v float64) {
	g.Altitude = float32(v)
}

//SetLat sets the 32-bit floating point latitude value via a 64-bit floating
//point value.
func (g *Gps) SetLat(v float64) {
	g.Latitude = float32(v)
}

//SetLon sets the 32-bit floating point longitude value via a 64-bit floating
//point value.
func (g *Gps) SetLon(v float64) {
	g.Longitude = float32(v)
}
