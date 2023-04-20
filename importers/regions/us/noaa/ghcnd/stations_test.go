package ghcnd

import (
	"testing"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/testing/passert"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/testing/ptest"
	ds "github.com/rsned/weather/datastructures"
)

func TestFoo(t *testing.T) {
	tests := []struct {
		have      string
		want      *ds.Station
		wantEmpty bool
	}{
		{
			have:      "",
			want:      nil,
			wantEmpty: true,
		},
		{
			have:      "pizza hamburgers",
			want:      nil,
			wantEmpty: true,
		},
		{
			// Input too short.
			have:      `USW00023234  37.6197 -122.3656    3.0 CA SAN FRANCISCO INT`,
			want:      nil,
			wantEmpty: true,
		},
		{
			have: `USW00023234  37.6197 -122.3656    3.0 CA SAN FRANCISCO INTL AP                  72494`,
			want: &ds.Station{
				ID:   "",
				Name: "SAN FRANCISCO INTL AP",
				Identifiers: &ds.Identifiers{
					WmoID:  "72494",
					GhcnID: "USW00023234",
				},
				Geography: &ds.Geography{
					Continent:        "North America",
					MetaRegion:       "NA",
					RegionCode:       "US",
					RegionName:       "United States",
					Subdivision1Code: "CA",
					ElevationMeters:  3,
					Lat:              37.619701,
					Lng:              -122.365601,
				},
				Attributions: &ds.Attributions{},
				StartDate:    "0000-01-01",
				EndDate:      "9999-12-31",
				LastUpdated:  "2023-04-15",
			},
			wantEmpty: false,
		},
	}

	beam.Init()
	for _, test := range tests {
		pipeline, scope := beam.NewPipelineWithRoot()
		// Turn the input string into a PCollection<string>
		inputs := beam.Create(scope, test.have)
		stations := beam.ParDo(scope, &StationParserFn{}, inputs)

		if test.wantEmpty {
			passert.Empty(scope, stations)
		} else {
			want := beam.Create(scope, test.want)
			passert.Equals(scope, stations, want)
		}

		if err := ptest.Run(pipeline); err != nil {
			t.Errorf("Failed to execute job: %v", err)
		}
	}
}
