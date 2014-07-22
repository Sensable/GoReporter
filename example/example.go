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
        Latitude: 52.5,
        Longitude: 13.3,
        Name: "some name here",
    }

    settings := GoReporter.Settings {
        AccessToken: "some-access-token",
        Private: false,
    }

    api := GoReporter.Api {
        "http://requestb.in/1jy9gkt1",
    }

    report := sensable.BuildReporter(settings, api);

    sample := GoReporter.Sample {
        Data: 32.5,
        Time: int64(time.Now().UnixNano() / 1e6),
    }

    fmt.Println(report(sample));
}