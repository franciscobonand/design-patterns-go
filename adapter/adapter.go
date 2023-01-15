package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
)

func minmax (a, b int) (int, int) {
    if a < b {
        return a, b
    } else {
        return b, a
    }
}

// The interface given by some API:
type Line struct {
    X1, Y1, X2, Y2 int
}

type VectorImage struct {
    Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
    width -= 1
    height -= 1
    return &VectorImage{
        []Line{
            {0, 0, width, 0},
            {0, 0, 0, height},
            {width, 0, width, height},
            {0, height, width, height},
        },
    }
}

// Interface we have:
type Point struct {
    X, Y int
}

type RasterImage interface {
    GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
    maxX, maxY := 0, 0
    points := owner.GetPoints()
    for _, pixel := range points {
        if pixel.X > maxX { maxX = pixel.X }
        if pixel.Y > maxY { maxY = pixel.Y }
    }
    maxX += 1
    maxY += 1

    data := make([][]rune, maxY)
    for i := 0; i < maxY; i++ {
        data[i] = make([]rune, maxX)
        for j := range data[i] { data[i][j] = ' ' }
    }

    for _, point := range points {
        data[point.Y][point.X] = '*'
    }

    b := strings.Builder{}
    for _, line := range data {
        b.WriteString(string(line))
        b.WriteRune('\n')
    }

    return b.String()
}

// Problem: I want to print a RasterImage but I can only make a VectorImage
// Solution: Create an Adapter!

type vectorToRasterAdapter struct {
    points []Point
}

func (va vectorToRasterAdapter) GetPoints() []Point {
    return va.points
}

func (va *vectorToRasterAdapter) addLine(line Line) {
    left, right := minmax(line.X1, line.X2)
    top, bottom := minmax(line.Y1, line.Y2)
    dx := right - left
    dy := line.Y2 - line.Y1

    if dx == 0 {
        for y := top; y <= bottom; y++ {
            va.points = append(va.points, Point{left, y})
        }
    } else if dy == 0 {
        for x := left; x <= right; x++ {
            va.points = append(va.points, Point{x, top})
        }
    }

    fmt.Println("generated", len(va.points), "points so far")
}

func VectorToRaster(vi *VectorImage) RasterImage {
    adapter := vectorToRasterAdapter{}

    for _, line := range vi.Lines {
        adapter.addLine(line)
    }

    return adapter // as RasterImage
}

// In case the adapter is immutable (internal state can't be altered)
// a cache mechanism can be created for it:

// Using a hashmap:
var pointCache = map[[16]byte] []Point{}

func (va *vectorToRasterAdapter) addLineCached(line Line) {
    hash := func (obj interface{}) [16]byte {
        bytes, _ := json.Marshal(obj)
        return md5.Sum(bytes)
    }

    h := hash(line)
    if pts, ok := pointCache[h]; ok {
        va.points = pts
        return
    }

    left, right := minmax(line.X1, line.X2)
    top, bottom := minmax(line.Y1, line.Y2)
    dx := right - left
    dy := line.Y2 - line.Y1

    if dx == 0 {
        for y := top; y <= bottom; y++ {
            va.points = append(va.points, Point{left, y})
        }
    } else if dy == 0 {
        for x := left; x <= right; x++ {
            va.points = append(va.points, Point{x, top})
        }
    }

    pointCache[h] = va.points
    fmt.Println("generated", len(va.points), "points so far")
}


func VectorToRasterCached(vi *VectorImage) RasterImage {
    adapter := vectorToRasterAdapter{}

    for _, line := range vi.Lines {
        adapter.addLineCached(line)
    }

    return adapter // as RasterImage
}

func main() {
    rc := NewRectangle(6, 4)

    fmt.Println("Without cache:")
    a := VectorToRaster(rc)
    _ = VectorToRaster(rc)
    fmt.Println(DrawPoints(a))


    fmt.Println("With cache:")
    ac := VectorToRasterCached(rc)
    _ = VectorToRasterCached(rc)
    fmt.Println(DrawPoints(ac))
}
