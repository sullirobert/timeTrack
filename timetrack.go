package timeTrack

import (
	"fmt"
	"io"
	"time"
)

//TimeTrack will write out the time duration to the specifed writer
func TimeTrack(start time.Time, name string, writer io.Writer) {
	elapsed := time.Since(start)
	message := fmt.Sprintf("\n%s took %s", name, elapsed)
	if writer == nil {
		fmt.Println(message)
	} else {
		writer.Write([]byte(message))
	}
}
