package timeTrack

import "log"

func ExampleTimeTrackImport() {
	log.Println(`import with _ to call init() on package load. this will start the timer.\nimport _ \"github.com/sullirobert/timeTrack\"`)
	// Output: import with _ to call init() on package load. this will start the timer.
}
