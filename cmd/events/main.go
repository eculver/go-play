package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
)

// input: uid, lat, lng, ts
// data:
//   uid: (lat,lng,ts)

const maxVelocity = 2

type event struct {
	uid int
	lat float64
	lng float64
	ts  int
}

func (e *event) Encode() string {
	// fmt.Printf("encoded: %d,%.2f,%.2f,%d\n", e.uid, e.lat, e.lng, e.ts)
	return fmt.Sprintf("%d,%.2f,%.2f,%d", e.uid, e.lat, e.lng, e.ts)
}

func (e *event) Decode(raw string) {
	pcs := strings.Split(raw, ",")
	if len(pcs) != 4 {
		return
	}

	uid, _ := strconv.Atoi(pcs[0])
	lat, _ := strconv.ParseFloat(pcs[1], 64)
	lng, _ := strconv.ParseFloat(pcs[2], 64)
	ts, _ := strconv.Atoi(pcs[3])

	// fmt.Printf("decoded ts: %d\n", ts)

	e.uid = uid
	e.lat = lat
	e.lng = lng
	e.ts = ts
}

func main() {
	// get input params
	enow := event{}
	enow.Decode(os.Args[1])

	// get redis client
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   1,
	})

	// find all matching events for the given user
	eventRaw, err := client.Get(string(enow.uid)).Result()
	// fmt.Printf("event raw: %s\n", eventRaw)
	elast := event{}
	elast.Decode(eventRaw)

	if err == nil {
		// is the last user location within the velocity parameters?
		d := getDistance(elast, enow)
		tdelta := enow.ts - elast.ts

		if tdelta < 0 {
			fmt.Println("illegal move: time delta negative")
			os.Exit(0)
		}

		velocity := d / float64(tdelta)

		fmt.Printf("distance: %.2f\n", d)
		fmt.Printf("time delta: %d\n", tdelta)
		fmt.Printf("velocity: %.2f\n", velocity)
		if velocity > maxVelocity {
			fmt.Println("illegal move")
			os.Exit(0)
		}
	}

	fmt.Println("legal move")

	// store last event
	client.Set(string(enow.uid), enow.Encode(), 0)
}

// sqrt( (x0 - x1)^2 + (y0 - y1)^2 )
func getDistance(e0, e1 event) float64 {
	xdist := math.Pow((e0.lng - e1.lng), 2)
	ydist := math.Pow((e0.lat - e1.lat), 2)
	return math.Sqrt(xdist + ydist)
}
