package datastructures

import (
	"fmt"
	"strings"
)

var (
	geographyFields []string
)

func init() {
	geographyFields = fields(&Geography{})
}

// Geography represents info about a specific entities location such as an
// ISO 3166-1 region, an administrative subdivision (such as Counties, Parish,
// Province, etc.), Localities, Postal areas, etc.
//
// This is as opposed to any type of more general objects or places that may exist
// in the world.
type Geography struct {
	// The following is a list of entities that contain this entity.
	// Values that are not known or don't exist are left blank.

	// String label of the continent enums.
	Continent string `beam:"continent" json:"continent"`
	// MetaRegion groups the world into a handful or related areas.
	// NA, SA, EMEA, APAC, etc.
	// Higher level bucketing into things like EMEA, NA, etc.
	MetaRegion string `beam:"meta_region" json:"meta_region"`
	// ISO 3166-1 English Display name.
	RegionName string `beam:"region_name" json:"region_name"`
	// ISO 3166-1 Alpha-2 Region Code.
	RegionCode string `beam:"region_code" json:"region_code"`
	// Top administrative subdivision in a region.
	Subdivision1Name string `beam:"subdivision_1_name" json:"subdivision_1_name"`
	// ISO 3166-2 Code.
	Subdivision1Code string `beam:"subdivision_1_code" json:"subdivision_1_code"`
	// Second level administrative subdivision in a region (if any).
	Subdivision2Name string `beam:"subdivision_2_name" json:"subdivision_2_name"`
	// Third level administrative subdivision in a region (if any).
	Subdivision3Name string `beam:"subdivision_3_name" json:"subdivision_3_name"`
	// Locality / city.
	Locality string `beam:"locality" json:"locality"`
	// Postal code.
	PostalCode string `beam:"postal_code" json:"postal_code"`

	// StreetAddress is the street address, if known, for the entity.
	StreetAddress string `beam:"street_address" json:"street_address"`

	Lat float32 `beam:"lat" json:"lat"`
	Lng float32 `beam:"lng" json:"lng"`

	// E7 forms are integer encoding of the lat/lng scaled by 1e7.
	LatE7 int32 `beam:"lat_e7" json:"lat_e7"`
	LngE7 int32 `beam:"lng_e7" json:"lng_e7"`

	// Datum is generally going to be WSG84.
	Datum string `beam:"datum" json:"datum"`

	ElevationMeters int32 `beam:"elevation_meters" json:"elevation_meters"`

	// Locations using other geographic systems for working with earth locations.

	// S2CellID is the s2geometry.io CellID for the entity.
	S2CellID uint64 `beam:"s2_cell_id" json:"s2_cell_id"`
	// TODO(rsned): Add H3 Geo, Plus Codes, GeoHash, and others.

	// Timezone is a TZData string like "PST8PDT", "AEST", "Etc/GMT-13",
	// "US/Pacific", or "America/Los_Angeles". No explicit offsets are stored
	// only strings which can handle STD/DST variations.
	Timezone string `beam:"time_zone" json:"time_zone"`

	// TODO(rsned): Expand this with additional common standard geographic ID systems.

}

func (g *Geography) String() string {
	return g.CSV(",")
}

func (g *Geography) CSV(delim string) string {
	// TODO(rsned): Move this to use full CSV encoding.
	return strings.Join(g.ValueColumns(), delim)
}

func (g *Geography) HeaderColumns(prefix string) []string {
	// TODO(rsned): Cache the list by prefix to save on redundant work.
	return prefixLabels(prefix, geographyFields)
}

func (g *Geography) ValueColumns() []string {
	return []string{
		g.Continent,
		g.MetaRegion,
		g.RegionName,
		g.RegionCode,
		g.Subdivision1Name,
		g.Subdivision1Code,
		g.Subdivision2Name,
		g.Subdivision3Name,
		g.Locality,
		g.PostalCode,
		g.StreetAddress,
		fmt.Sprintf("%f", g.Lat),
		fmt.Sprintf("%f", g.Lng),
		fmt.Sprintf("%d", g.LatE7),
		fmt.Sprintf("%d", g.LngE7),
		g.Datum,
		fmt.Sprintf("%d", g.ElevationMeters),
		fmt.Sprintf("0x%x", g.S2CellID),
		g.Timezone,
	}
}
