package main

import (
        "github.com/sensable/GoReporter";
        "time";
        "fmt"
    )

const requestbinUrl = "http://requestb.in/1696gmd1"

func main() {
    sensable := GoReporter.Sensable {
        SensorId: "some-id",
        Unit: "°c",
        SensorType: "temperature",
        Latitude: 52.5,
        Longitude: 13.3,
        Name: "some name here",
    }

    settings := GoReporter.Settings {
        AccessToken: "some-access-token",
        Private: false,
    }

    fmt.Println("Check result at " + requestbinUrl + "?inspect")

    reporter := sensable.BuildReporter(settings, requestbinUrl)

    sample := GoReporter.Sample {
        Data: 32.5,
        Time: int64(time.Now().UnixNano() / 1e6),
        State: "it's getting warmer",
    }

    _, err := reporter.Report(sample)

    if err != nil {
        fmt.Println("Something went wrong", err)
    }
}