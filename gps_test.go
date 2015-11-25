package binmsg

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"testing"
)

func TestGpsAlt(t *testing.T) {
	var (
		g    Gps
		want float32
		got  float32
	)

	log.SetOutput(ioutil.Discard)

	want = math.Nextafter32(123.45, 124.0)
	g.Altitude = want
	got = float32(g.Alt())
	if got != want {
		s := fmt.Sprintf("Got: %f, want: %f", got, want)
		t.Fatal(s)
	}
}

func TestGpsLat(t *testing.T) {
	var (
		g    Gps
		want float32
		got  float32
	)

	log.SetOutput(ioutil.Discard)

	want = math.Nextafter32(123.45, 124.0)
	g.Latitude = want
	got = float32(g.Lat())
	if got != want {
		s := fmt.Sprintf("Got: %f, want: %f", got, want)
		t.Fatal(s)
	}
}

func TestGpsLon(t *testing.T) {
	var (
		g    Gps
		want float32
		got  float32
	)

	log.SetOutput(ioutil.Discard)

	want = math.Nextafter32(123.45, 124.0)
	g.Longitude = want
	got = float32(g.Lon())
	if got != want {
		s := fmt.Sprintf("Got: %f, want: %f", got, want)
		t.Fatal(s)
	}
}

func TestGpsSetAlt(t *testing.T) {
	var (
		g    Gps
		want float32
		got  float32
	)

	log.SetOutput(ioutil.Discard)

	want = math.Nextafter32(123.45, 124.0)
	g.SetAlt((float64(want)))
	got = g.Altitude
	if got != want {
		s := fmt.Sprintf("Got: %f, want: %f", got, want)
		t.Fatal(s)
	}
}

func TestGpsSetLat(t *testing.T) {
	var (
		g    Gps
		want float32
		got  float32
	)

	log.SetOutput(ioutil.Discard)

	want = math.Nextafter32(123.45, 124.0)
	g.SetLat((float64(want)))
	got = g.Latitude
	if got != want {
		s := fmt.Sprintf("Got: %f, want: %f", got, want)
		t.Fatal(s)
	}
}

func TestGpsSetLon(t *testing.T) {
	var (
		g    Gps
		want float32
		got  float32
	)

	log.SetOutput(ioutil.Discard)

	want = math.Nextafter32(123.45, 124.0)
	g.SetLon((float64(want)))
	got = g.Longitude
	if got != want {
		s := fmt.Sprintf("Got: %f, want: %f", got, want)
		t.Fatal(s)
	}
}
