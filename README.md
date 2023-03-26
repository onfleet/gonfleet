# gonfleet

The official [Onfleet](https://onfleet.com/) Go client library.

## Installation

```bash
go get github.com/onfleet/gonfleet
```

## Documentation

Below are a few examples.

For comprehensive documentation / examples, visit [Onfleet API reference](https://docs.onfleet.com/)

### Initialize Client

```go
import (
    "log"
    "os"

    "github.com/onfleet/gonfleet/client"
)

apiKey := os.Getenv("onfleet_api_key")

client, err := client.New(apiKey, nil)
if err != nil {
    log.Fatal(err)
}

// do something with client ...
```

### Workers

```go
import (
    "fmt"

	onfleet "github.com/onfleet/gonfleet"
    "github.com/onfleet/gonfleet/client"
)

params := onfleet.WorkerCreateParams{
    Addresses: &onfleet.WorkerCreateParamsAddressRouting{
        Routing: "destination_id",
    },
    Capacity: 10,
    Teams:    []string{"team_id_a", "team_id_b"},
    Name:     "Janis Joplin",
    Phone:    "+13105550101",
    Vehicle: &onfleet.WorkerCreateParamsVehicle{
        Type: onfleet.WorkerVehicleTypeBicycle,
    },
}

worker, err := client.Workers.Create(params)
if err != nil {
    fmt.Println(err)
    return
}

// do something with worker
```
