package main

import (
    "fmt"
    "net/url"
    "os"
    "strings"
)

// Single responsability principle

// Journal responsability is to handle entries
type Journal struct {
    entries []string
    entryCount int
}

func (j *Journal) AddEntry(text string) int {
    j.entryCount++
    entry := fmt.Sprintf("%d: %s", j.entryCount, text)
    j.entries = append(j.entries, entry)
    return j.entryCount
}

func (j *Journal) RemoveEntry(index int) {
    // Removed an entry
}

func (j *Journal) String() string {
    return strings.Join(j.entries, "/n")
}

/*
One could implement the following methods to save and load a Journal
into/from a file, but i'd violate the single responsability principle.

It's important to have separation of concerns: a Journal should handle 
entries, and the storage (save/load) should be done by another structure,
as there might exist structures other than Journal which also need to be 
saved/loaded from a file and the methods to do so are pretty similar.
*/

func (j *Journal) Save(filename string) {
    _ =  os.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
    // Load from a file
}

func (j *Journal) LoadFromWeb(url *url.URL) {
    // Load from web
}

// Separating concerns examples:

// Using a method
func SaveToFile(j *Journal, filename string) {
    _ =  os.WriteFile(filename, []byte(j.String()), 0644)
}

// Using a structure
type Persistence struct {
    lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
    _ =  os.WriteFile(
        filename,
        []byte(strings.Join(j.entries, p.lineSeparator)),
        0644,
    )
}

// Usage example
func mainSRP() {
    j := &Journal{}
    j.AddEntry("Oh hello there")
    j.AddEntry("A very cool entry")

    // Instead of using the following:
    j.Save("journal.txt")

    // Use one of:
    SaveToFile(j, "journal.txt")
    // or
    p := Persistence{lineSeparator: "-"}
    p.SaveToFile(j, "journal.txt")
}
