package datastructures

import "testing"

func TestObservationCSV(t *testing.T) {
	tests := []struct {
		have *Observation
		want string
	}{
		//
		{
			have: &Observation{
				StationID: "",
				TempC:     -9999,
			},
			want: ",,,-9999",
		},
	}

	for _, test := range tests {

		if got := test.have.CSV(","); got != test.want {
			t.Errorf("CSV(%q) = %q, want %q", test.have, got, test.want)
		}
	}
}
