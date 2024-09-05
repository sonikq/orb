package mvt

import (
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/maptile"
	"github.com/paulmach/orb/project"
)

const (
	// DefaultExtent for mapbox vector tiles. (https://www.mapbox.com/vector-tiles/specification/)
	DefaultExtent = 4096
)

// Layer is intermediate MVT layer to be encoded/decoded or projected.
type Layer struct {
	Name     string
	Version  uint32
	Extent   uint32
	Features []geojson.Feature
}

// ProjectToTile will project all the geometries in the layer
// to tile coordinates based on the extent and the mercator projection.
func (l *Layer) ProjectToTile(tile maptile.Tile) {
	p := newProjection(tile, l.Extent)
	for _, f := range l.Features {
		f.Geometry = project.Geometry(f.Geometry, p.ToTile)
	}
}

// ProjectToWGS84 will project all the geometries backed to WGS84 from
// the extent and mercator projection.
func (l *Layer) ProjectToWGS84(tile maptile.Tile) {
	p := newProjection(tile, l.Extent)
	for _, f := range l.Features {
		f.Geometry = project.Geometry(f.Geometry, p.ToWGS84)
	}
}

// Layers is a set of layers.
type Layers []*Layer

// ProjectToTile will project all the geometries in all layers
// to tile coordinates based on the extent and the mercator projection.
func (ls Layers) ProjectToTile(tile maptile.Tile) {
	for _, l := range ls {
		l.ProjectToTile(tile)
	}
}

// ProjectToWGS84 will project all the geometries in all the layers backed
// to WGS84 from the extent and mercator projection.
func (ls Layers) ProjectToWGS84(tile maptile.Tile) {
	for _, l := range ls {
		l.ProjectToWGS84(tile)
	}
}
