package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "sync"
)

type singletonDatabase struct {
    capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
    return db.capitals[name]
}

// sync.Once or init() to make it thread safe
// laziness is not guaranteed using 'init()'
var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
    once.Do(func() {
        capitals, err := readData("capitals.txt")
        if err != nil {
            panic(err.Error())
        }
        db := singletonDatabase{capitals}
        instance = &db
    }) 

    return instance
}

func readData(path string) (map[string]int, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    result := map[string]int{}

    for scanner.Scan() {
        k := scanner.Text()
        scanner.Scan()
        v, _ := strconv.Atoi(scanner.Text())
        result[k] = v
    }

    return result, nil
}

func GetTotalPopulation(cities []string) int {
    result := 0
    for _, city := range cities {
        result += GetSingletonDatabase().GetPopulation(city)
    }
    return result
}

// Solving the DIP issue:
type Database interface {
    GetPopulation(name string) int
}

func GetTotalPopulationEx(db Database, cities []string) int {
    result := 0
    for _, city := range cities {
        result += db.GetPopulation(city)
    }
    return result
}

type DummyDatabase struct {
    dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
    if len(d.dummyData) == 0 {
        d.dummyData = map[string]int {
            "alpha": 1,
            "beta": 2,
            "gamma": 3,
        }
    }

    return d.dummyData[name]
}

func main() {
    // Using the singleton:
     db := GetSingletonDatabase()
     pop := db.GetPopulation("Seoul")
     fmt.Println("Population of Seoul =", pop)

    // Problems with singleton:
    // cities := []string{"Seoul", "Manila"}
    // tp := GetTotalPopulation(cities)
    // ok := tp == (17500000 + 14750000)
    // fmt.Println(ok)
    /*
    The test above depends on data from a real live database, and in most of
    the times tests aren't run with real live databases (so these are magic values)

    This test is also testing the database itself, so it's not a unit test

    The 'GetTotalPopulation' method is also violating the dependency inversion
    principle, as it's using the database method direcly instead of an abstraction
    */
    
    // Testing using the 'Database' abstraction so it's not relying on a live db:
    names := []string{"alpha", "gamma"}
    tp := GetTotalPopulationEx(&DummyDatabase{}, names)
    fmt.Println(tp == 4)
}
