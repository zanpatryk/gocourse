package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/zanpatryk/gocourse/shape"
)

func run(args []string, out io.Writer) int {
	fs := flag.NewFlagSet("shapes", flag.ContinueOnError)
	fs.SetOutput(out)

	shapePtr := fs.String("shape", "", "type of shape(required)")
	widthPtr := fs.Float64("width", 1, "rectangle width")
	heightPtr := fs.Float64("height", 1, "rectangle height")
	radiusPtr := fs.Float64("radius", 1, "circle height")

	if err := fs.Parse(args); err != nil {
		fmt.Fprintln(out, err)
		return 2
	}

	if *shapePtr == "" {
		fmt.Fprintln(out, "[ERROR] -shape flag is required")
		fs.Usage()
		return 1
	}

	shapeType := strings.ToLower(*shapePtr)

	switch shapeType {
	case "rectangle":
		if *widthPtr < 0 || *heightPtr < 0 {
			fmt.Fprintln(out, "[ERROR] -width and -height cannot be less than 0")
			return 1
		}
		r := shape.Rectangle{Width: *widthPtr, Height: *heightPtr}
		fmt.Fprintf(out, "Rectangle Area: %.4f", r.Area())
		return 0

	case "circle":
		if *radiusPtr < 0 {
			fmt.Fprintln(out, "[ERROR] -radius cannot be less than 0")
			return 1
		}
		c := shape.Circle{Radius: *radiusPtr}
		fmt.Fprintf(out, "Circle Area: %.4f", c.Area())
		return 0

	default:
		fmt.Fprintln(out, "[ERROR] Invalid shape!")
		return 1
	}
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout))
}
