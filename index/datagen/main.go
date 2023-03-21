package main

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/jaswdr/faker"
)

type Person struct {
	Index     int
	Name      string
	Age       int
	CreatedAt time.Time
}

func generatePersons() <-chan Person {
	threads := 20

	ch := make(chan Person)

	persons := 20000000

	wg := sync.WaitGroup{}
	for i := 0; i < threads; i++ {
		wg.Add(1)

		go func(thread int) {
			fake := faker.New()

			for j := 0; j < persons/threads; j++ {
				ch <- Person{
					Index:     thread + (j * threads),
					Name:      fake.Person().Name(),
					Age:       rand.Intn(120),
					CreatedAt: time.Now(),
				}
			}

			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {
	ch := generatePersons()

	f, err := os.Create("persons.csv")
	if err != nil {
		panic(err)
	}

	w := csv.NewWriter(f)

	for p := range ch {
		w.Write([]string{strconv.Itoa(p.Index), p.Name, strconv.Itoa(p.Age), p.CreatedAt.Format(time.RFC3339)})
	}

	w.Flush()

	f.Close()
}
