package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "time"

    ex "github.com/markus-wa/demoinfocs-golang/v4/examples"
    demoinfocs "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
    common "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
    events "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
)

// Run like this: go run parse_demo.go -demo /path/to/demo.dem
func main() {
    f, err := os.Open(ex.DemoPathFromArgs())
    checkError(err)

    defer f.Close()

    csvFile, err := os.Create("data.csv")
    checkError(err)
    csvFile2, err := os.Create("data2.csv")
    checkError(err)

    writer := csv.NewWriter(csvFile)

    results := [][]string{{"source", "sourceTeam", "target", "targetTeam", "weapon", "headshot", "wallbang", "Time", "StrTime"}}
    err = writer.Write(results[0])
    checkError(err)
    writer.Flush()

    p := demoinfocs.NewParser(f)
    defer p.Close()


    // Parse header
    header, err := p.ParseHeader()
    checkError(err)
    fmt.Println("Map:", header.MapName)

    // Register handler on kill events
    p.RegisterEventHandler(func(e events.Kill) {
        if e.Killer != nil {
            var hs string
            var hsCSV string
            if e.IsHeadshot {
                hs = " (HS)"
                hsCSV = "Yes"
            } else {
                hsCSV = "No"
            }
            var wallBang string
            var wallBangCSV string
            if e.PenetratedObjects > 0 {
                wallBang = " (WB)"
                wallBangCSV = "Yes"
            } else {
                wallBangCSV = "No"
            }
            source := playerName(e.Killer)
            sourceTeam := returnTeam(e.Killer)
            target := playerName(e.Victim)
            targetTeam := returnTeam(e.Victim)
            tempSlice := make([]string, 0)
            tempSlice = append(tempSlice, source, sourceTeam, target, targetTeam, fmt.Sprintf("%v", e.Weapon), hsCSV, wallBangCSV, fmt.Sprintf("%d", p.CurrentTime()), fmt.Sprintf("%s", p.CurrentTime().Truncate(time.Second)))
            err := writer.Write(tempSlice)
            checkError(err)
            writer.Flush()
            results = append(results, tempSlice)

            fmt.Printf("%s <%v%s%s> %s, %d\n", formatPlayer(e.Killer), e.Weapon, hs, wallBang, formatPlayer(e.Victim), p.CurrentTime())
        }
    })

    p.RegisterEventHandler(func(e events.TeamSideSwitch) {
        fmt.Printf("Switched sides\n")
        csvFile.Close()
        writer = csv.NewWriter(csvFile2)
        results := [][]string{{"source", "sourceTeam", "target", "targetTeam", "weapon", "headshot", "wallbang", "Time", "StrTime"}}
        err = writer.Write(results[0])
        checkError(err)
        writer.Flush()
    })

    // Parse to end
    err = p.ParseToEnd()
    checkError(err)
    csvFile2.Close()
}

func playerName(p *common.Player) string {
    if p == nil {
        return "?"
    }
    return p.Name
}

func formatPlayer(p *common.Player) string {
    if p == nil {
        return "?"
    }

    switch p.Team {
    case common.TeamTerrorists:
        return "[T]" + p.Name
    case common.TeamCounterTerrorists:
        return "[CT]" + p.Name
    }

    return p.Name
}

func returnTeam(p *common.Player) string {
    if p == nil {
        return "?"
    }

    switch p.Team {
    case common.TeamTerrorists:
        return "T"
    case common.TeamCounterTerrorists:
        return "CT"
    }

    return "?"
}


func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
