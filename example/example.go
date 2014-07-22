package main

import (
        "github.com/sensable/GoReporter";
        "time";
        "fmt"
    )

const requestbinUrl = "http://requestb.in/1jy9gkt1"

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

    api := GoReporter.Api {
        requestbinUrl,
    }

    fmt.Println("Check result at " + requestbinUrl + "?inspect")

    report := sensable.BuildReporter(settings, api)

    sample := GoReporter.Sample {
        Data: 32.5,
        Time: int64(time.Now().UnixNano() / 1e6),
        State: "it's getting warmer",
    }

    _, err := report(sample)

    if err != nil {
        fmt.Println("Something went wrong", err)
    }
}