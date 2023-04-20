package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"flag"
	"log"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"

	ds "github.com/rsned/weather/datastructures"
	"github.com/rsned/weather/importers/regions/us/noaa/ghcnd"
)

var (
	input  = flag.String("input", "", "File(s) to read.")
	output = flag.String("output", "", "Output file (required).")
)

func init() {
	register.Function2x0(generateID)
	register.Function2x0(stationToCSV)
	register.Emitter1[*ds.Station]()
	register.Emitter1[string]()
}

func generateID(s *ds.Station, emit func(*ds.Station)) {
	s.ID = s.Identifiers.GhcnID
	emit(s)
}

// Convert the station to a form that is serializable.
func stationToCSV(s *ds.Station, emit func(string)) {
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	w.UseCRLF = false
	w.Write(s.ValueColumns())
	w.Flush()
	emit(buf.String())
}

func main() {
	flag.Parse()
	beam.Init()
	ctx := context.Background()

	if *output == "" {
		log.Fatal("No output provided")
	}

	pipeline := beam.NewPipeline()
	scope := pipeline.Root()

	// Reading station inputs in.

	// Start with the the source we feel is the "fullest" starting point.
	// For now this is NOAA GHCN-D, eventually NOAA MSHR.
	lines := textio.Read(scope, *input)

	// Create the initial partial station objects for the lines.
	initial := beam.ParDo(scope, &ghcnd.StationParserFn{}, lines)

	// For each additional source to try to merge in:
	//   Read in its lines and convert to partial station objects.
	//   Run the merge and straggler function on existing PCol and the new PCol.

	// Merge all records into one PCollection.

	// Now that all merges have completed, generate the final station ID.
	stations := beam.ParDo(scope, generateID, initial)

	// Convert the station to a form that is serializable.
	formatted := beam.ParDo(scope, stationToCSV, stations)

	// Save to disk.
	textio.Write(scope, *output, formatted)

	if err := beamx.Run(ctx, pipeline); err != nil {
		log.Fatalf("Failed to execute job: %v", err)
	}
}
