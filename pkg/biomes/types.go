package biomes

type Biome string

type GeoLocation struct {
	Longitude float64
	Latitude  float64
}

const (
	Forest   Biome = "forest"
	Mountain Biome = "mountain"
)
