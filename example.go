package timeTrack

import "time"

func ExampleTimeTrackImport() {
	defer TimeTrack(time.Now(), "main", nil)
}
