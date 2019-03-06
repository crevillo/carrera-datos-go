package main

import (
    "fmt"
    "os"
    "log"
    "encoding/csv"
    "io"
)

func main() {
    home := os.Getenv("HOME")

	f, err := os.Open(home + "/data-race/natalidad000000000000.csv")

        if err != nil {
            log.Fatal(err)
        }

        defer f.Close()

        r := csv.NewReader(f)

        for {
            record, err := r.Read()

            if err == io.EOF {
                break
            }

            if err != nil {
                log.Fatal(err)
            }

            readline(record)
        }
}

func readline(line []string) {
    state := line[5]
    bornDecade := line[1][:len(line[1]) - 1] + '0'
    fmt.Println(state + bornDecade)
}
