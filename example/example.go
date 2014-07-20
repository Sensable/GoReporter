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

    options := GoReporter.Options {
        AccessToken: "some-access-token",
        Private: false,
    }

    report := sensable.BuildReporter(options);

    sample := GoReporter.Sample {
        Data: 32.5,
        Time: int64(time.Now().Unix()),
    }

    fmt.Println(report(sample));
}