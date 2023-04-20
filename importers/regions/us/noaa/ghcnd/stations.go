package ghcnd

import (
	"strings"

	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/rsned/weather/importers/utils"

	ds "github.com/rsned/weather/datastructures"
)

// StationParserFn is an Apache Beam structural DoFn to process rows from a GHCN-D staton file.
type StationParserFn struct {
}

func init() {
	register.DoFn2x0[string, func(*ds.Station)](&StationParserFn{})
	register.Emitter1[*ds.Station]()

}

// ProcessElement reads one row in and attempts to convert it into a Station.
func (s *StationParserFn) ProcessElement(line string, emit func(*ds.Station)) {
	if len(line) != 85 {
		return
	}

	station := ds.EmptyStation()

	// https://www.ncei.noaa.gov/pub/data/ghcn/daily/readme.txt
	//
	// IV. FORMAT OF "ghcnd-stations.txt"
	//
	// ------------------------------
	// Variable   Columns   Type
	// ------------------------------
	// ID            1-11   Character
	// LATITUDE     13-20   Real
	// LONGITUDE    22-30   Real
	// ELEVATION    32-37   Real
	// STATE        39-40   Character
	// NAME         42-71   Character
	// GSN FLAG     73-75   Character
	// HCN/CRN FLAG 77-79   Character
	// WMO ID       81-85   Character
	// ------------------------------

	// ID         is the station identification code.  Note that the first two
	//            characters denote the FIPS  country code, the third character
	//            is a network code that identifies the station numbering system
	//            used, and the remaining eight characters contain the actual
	//            station ID.
	//
	// The first character of the ID relates the ID to other ID systems as well.
	//
	station.Identifiers.GhcnID = strings.TrimSpace(line[0:11])

	// TODO(rsned): Build tool to parse and convert these into a standardized form.
	// LATITUDE   is latitude of the station (in decimal degrees).
	// LONGITUDE  is the longitude of the station (in decimal degrees).
	station.Geography.Lat = float32(utils.ParseFloat(line[12:20], 0))
	station.Geography.Lng = float32(utils.ParseFloat(line[21:30], 0))

	// ELEVATION  is the elevation of the station (in meters, missing = -999.9).
	station.Geography.ElevationMeters = int32(utils.ParseFloat(line[31:37], -9999))

	// TODO(rsned): Update these to be dynamic based on the actual rows values.
	station.Geography.Continent = "North America"
	station.Geography.MetaRegion = "NA"
	station.Geography.RegionCode = "US"
	station.Geography.RegionName = "United States"

	// TODO(rsned): Convert this to ISO 3166-2 form.
	// STATE      is the U.S. postal code for the state (for U.S. stations only).
	station.Geography.Subdivision1Code = strings.TrimSpace(line[38:40])

	station.Name = strings.TrimSpace(line[41:71])

	// TODO(rsned): Are these flags of interest to keep?

	// GSN FLAG  is a flag that indicates whether the station is part of the GCOS
	//           Surface Network (GSN). The flag is assigned by cross-referencing
	//           the number in the WMOID field with the official list of GSN
	//           stations. There are two possible values:
	//
	//           Blank = non-GSN station or WMO Station number not available
	//           GSN   = GSN station
	//
	// HCN/      is a flag that indicates whether the station is part of the U.S.
	// CRN FLAG  Historical Climatology Network (HCN) or U.S. Climate Refererence
	//           Network (CRN).  There are three possible values:
	//
	//           Blank = Not a member of the U.S. Historical Climatology
	//	           or U.S. Climate Reference Networks
	//           HCN   = U.S. Historical Climatology Network station
	//           CRN   = U.S. Climate Reference Network or U.S. Regional Climate
	//	           Network Station
	// gsnFlag := strings.TrimSpace(line[72:75])
	// hcnFlag := strings.TrimSpace(line[76:79])

	// WMO ID    is the World Meteorological Organization (WMO) number for the
	// station. If the station has no WMO number (or one has not yet
	// been matched to this station), then the field is blank.
	station.Identifiers.WmoID = strings.TrimSpace(line[80:85])

	// TODO(rsned): Update these to be dynamic.
	station.StartDate = "0000-01-01"
	station.EndDate = "9999-12-31"
	station.LastUpdated = "2023-04-15"

	emit(station)
}
