# Go Reporter for Sensable.io

## Disclaimer

This is the first Go code I have ever written.

## Example

```go
package main

import (
        "github.com/sensable/GoReporter";
        "time";
        "fmt"
    )

func main() {
    sensable := GoReporter.Sensable {
        SensorId: "some-id",
        Unit: "°c",
        SensorType: "temperature",
        Location: GoReporter.Location {
            Latitude: 52.5,
            Longitude: 13.3,
        },
        Name: "some name here",
    }

    settings := GoReporter.Settings {
        AccessToken: "some-access-token", //Get yours on sensable.io
        Private: false,
    }

    reporter := sensable.BuildReporter(settings)

    sample := GoReporter.Sample {
        Value: 32.5,
        Timestamp: int64(time.Now().UnixNano() / 1e6),
        State: "it's getting warmer",
    }

    _, err := reporter.Report(sample)

    if err != nil {
        fmt.Println("Something went wrong", err)
    }
}
```