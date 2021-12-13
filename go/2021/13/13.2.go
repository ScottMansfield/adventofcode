package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y uint16
}

type fold struct {
	axis byte
	loc  uint16
}

type paper map[point]struct{}

func foldPaper(pap paper, f fold) paper {
	if f.axis != 'x' && f.axis != 'y' {
		panic(f.axis)
	}

	ret := make(paper)

	// subtract the distance to the line from the line
	if f.axis == 'y' {
		// modify y values
		for p := range pap {
			if p.y > f.loc {
				// make sure to remove the existing point
				delete(pap, p)
				// fold!
				p.y = f.loc - (p.y - f.loc)
			}

			ret[p] = struct{}{}
		}
		return ret
	}

	// x axis line, modify x values
	for p := range pap {
		if p.x > f.loc {
			// remove existing point
			delete(pap, p)
			// fold!
			p.x = f.loc - (p.x - f.loc)
		}

		ret[p] = struct{}{}
	}

	return ret

}

func readInput(s *bufio.Scanner) (paper, []fold) {
	pap := make(paper)
	var folds []fold
	readingPoints := true

	for s.Scan() {
		if readingPoints {
			if s.Text() == "" {
				readingPoints = false
				continue
			}

			parts := strings.Split(s.Text(), ",")

			x, err := strconv.ParseUint(parts[0], 10, 16)
			if err != nil {
				panic(err)
			}

			y, err := strconv.ParseUint(parts[1], 10, 16)
			if err != nil {
				panic(err)
			}

			pap[point{uint16(x), uint16(y)}] = struct{}{}

		} else { // read fold instructions
			temp := strings.TrimPrefix(s.Text(), "fold along ")
			parts := strings.Split(temp, "=")

			loc, err := strconv.ParseUint(parts[1], 10, 16)
			if err != nil {
				panic(err)
			}

			folds = append(folds, fold{parts[0][0], uint16(loc)})
		}
	}

	return pap, folds
}

func maxes(pap paper) (uint16, uint16) {
	var maxx, maxy uint16

	for p := range pap {
		if p.x > maxx {
			maxx = p.x
		}
		if p.y > maxy {
			maxy = p.y
		}
	}

	return maxx, maxy
}

func main() {
	infile, err := os.Open("13.input")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(infile)

	pap, folds := readInput(s)

	// fmt.Println(pap)
	// fmt.Println(folds)

	fmt.Println(len(pap))

	for _, f := range folds {
		fmt.Println("fold", f)
		pap = foldPaper(pap, f)
		fmt.Println(len(pap))
		//fmt.Println(pap)
	}

	maxx, maxy := maxes(pap)

	fmt.Println(maxx, maxy)

	img := image.NewGray(
		image.Rectangle{
			Min: image.Point{},
			Max: image.Point{
				X: int(maxx) + 1,
				Y: int(maxy) + 1,
			},
		},
	)

	for p := range pap {
		img.Set(int(p.x), int(p.y), color.White)
	}

	outfile, err := os.Create("13.png")

	err = png.Encode(outfile, img)
	if err != nil {
		panic(err)
	}
}
