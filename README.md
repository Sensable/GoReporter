# Go Reporter for Sensable.io

## Disclaimer

This is the first Go code I have ever written.

## Documentation

The documentation for this package is available at: [godoc.org/github.com/Sensable/GoReporter](http://godoc.org/github.com/Sensable/GoReporter)

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

## Run the tests

Unit tests use [testify](https://github.com/stretchr/testify) for assertions. You can install it with

    go get github.com/stretchr/testify

To run the tests simply do

    go test

To run the tests and obtain the coverage report you can use

    go test -coverprofile=coverage.out

The coverage report can then be visualized in HTML format with

    go tool cover -html=coverage.out