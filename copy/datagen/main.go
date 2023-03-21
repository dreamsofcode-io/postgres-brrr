package main

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"math/rand"
	"net"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

type Event struct {
	ID        uuid.UUID `faker:"uuid"`
	Source    net.IP    `faker:"ip"`
	Payload   []byte    `faker:"payload"`
	Timestamp string    `faker:"timestamper"`
}

func setup() {
	rand.Seed(time.Now().UnixNano())

	_ = faker.AddProvider("uuid", func(v reflect.Value) (interface{}, error) {
		return uuid.New(), nil
	})

	_ = faker.AddProvider("ip", func(v reflect.Value) (interface{}, error) {
		return random_ip(), nil
	})

	_ = faker.AddProvider("timestamper", func(v reflect.Value) (interface{}, error) {
		days := 1 //rand.Intn(3) + 1
		secs := rand.Intn(60*60*24*days - 1)
		dt := time.Date(2023, 3, 18, 23, 59, 59, 0, time.UTC).Add(time.Duration(secs) * time.Second * -1)
		return dt.Format(time.RFC3339), nil
	})

	_ = faker.AddProvider("payload", func(v reflect.Value) (interface{}, error) {
		methods := strings.Split("GET,POST,PUT,DELETE,HEAD,OPTIONS,TRACE,CONNECT", ",")
		components := []string{
			"admin",
			"api",
			"auth",
			"blog",
			"cdn",
			"cdn2",
			"cdn3",
			"cdn4",
			"web",
			"www",
			"files",
			"images",
			"img",
			"static",
			"static1",
			"static2",
			"ip",
			"users",
			"user",
			"login",
			"logout",
			"register",
			"signup",
			"signin",
			"signout",
			"account",
			"accounts",
			"profile",
			"profiles",
			"dashboard",
			"dashboards",
			"admin",
			"admins",
			"administrator",
			"administrators",
			"home",
			"index",
			"about",
			"contact",
			"contacts",
			"contact-us",
			"contact_us",
		}

		extensions := []string{
			"html",
			"htm",
			"php",
			"asp",
			"aspx",
			"jsp",
			"js",
			"css",
			"png",
			"jpg",
			"jpeg",
			"gif",
			"svg",
			"ico",
			"json",
			"xml",
			"txt",
			"pdf",
			"doc",
		}

		isFile := rand.Intn(2) == 0

		numComponents := rand.Intn(5)

		pathComps := make([]string, numComponents)

		for i := 0; i < numComponents; i++ {
			pathComps[i] = components[rand.Intn(len(components))]
		}

		path := "/" + strings.Join(pathComps, "/")

		if isFile {
			path += "." + extensions[rand.Intn(len(extensions))]
		}

		method := methods[rand.Intn(len(methods))]

		return []byte(fmt.Sprintf("%s %s", method, path)), nil
	})
}

func random_ip() net.IP {
	buf := make([]byte, 4)
	ip := rand.Uint32()

	binary.LittleEndian.PutUint32(buf, ip)

	return net.IP(buf)
}

func generateEvents() chan Event {
	threads := 40
	ch := make(chan Event)

	events := 100000

	wg := sync.WaitGroup{}
	for th := 0; th < threads; th++ {
		wg.Add(1)

		go func() {
			for i := 0; i < events/threads; i++ {
				var event Event
				_ = faker.FakeData(&event)
				ch <- event
			}

			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		fmt.Println("Closing chan")
		close(ch)
	}()

	return ch
}

func main() {
	setup()

	ch := generateEvents()

	f, err := os.Create("events_new.csv")
	if err != nil {
		panic(err)
	}

	w := csv.NewWriter(f)

	for ev := range ch {
		w.Write([]string{
			ev.ID.String(),
			ev.Source.String(),
			base64.RawStdEncoding.EncodeToString(ev.Payload),
			ev.Timestamp,
		})
	}

	w.Flush()

	f.Close()
}
