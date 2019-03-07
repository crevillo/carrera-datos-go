package main

import (
    "fmt"
    "os"
    "log"
    "encoding/csv"
    "io"
    "time"
)

func main() {
    now := time.Now()
    unixNano := now.UnixNano()
    umillisec := unixNano / 1000000

	f, err := os.Open("n.csv")

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

    now2 := time.Now()
    unixNano2 := now2.UnixNano()
    umillisec_end := unixNano2 / 1000000
    seconds := float64(umillisec_end - umillisec) / 1000

    fmt.Println("Go takes:", seconds)
}

func readline(line []string) {
    //state := line[5]
    //bornDecade := line[1][:len(line[1]) - 1] + "0"
    //childRace := line[7]
    //isMale := line[6]
    //fmt.Println(state + bornDecade)
}
