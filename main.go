package main

import (
	"github.com/ajstarks/svgo"
	"log"
	"net/http"
)

var (
	height = 10000
	width  = 10000
)

func main() {
	http.Handle("/singlemansvg", http.HandlerFunc(SingleManSVG))
	http.Handle("/singlewomansvg", http.HandlerFunc(SingleWomanSVG))
	http.Handle("/couplesvg", http.HandlerFunc(CoupleSVG))
	http.Handle("/onekidsvg", http.HandlerFunc(OneKidSVG))
	http.Handle("/twokidsvg", http.HandlerFunc(TwoKidSVG))
	http.Handle("/threekidsvg", http.HandlerFunc(ThreeKidSVG))
	http.Handle("/fourkidsvg", http.HandlerFunc(FourKidSVG))
	http.Handle("/fivekidsvg", http.HandlerFunc(FiveKidSVG))
	err := http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func CoupleSVG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(width, height)

	for i := 0; i <= width; i = i + 50 {
		s.Line(0, i, width, i, "fill:black;stroke:black")
	}

	for j := 0; j <= height; j = j + 10 {
		s.Line(j, 0, j, height, "fill:black;stroke:gray;stroke-dasharray:2")
	}

	/*
		echo '<rect x="'.($myfather*$mygrid-($mygrid/2)).'" y="'.($mygrid*2-($mygrid/2)).'" width="'.$mygrid.'" height="'.$mygrid.'"
		style="fill:blue;stroke:pink;stroke-width:5;fill-opacity:0.1;stroke-opacity:1" />';
	*/

	s.Square(125, 125, 250, "fill:blue")
	s.Line(375, 250, 1875, 250, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Circle(2000, 250, 125, "fill:red")
	s.End()
}

func SingleManSVG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(width, height)

	for i := 0; i <= width; i = i + 50 {
		s.Line(0, i, width, i, "fill:black;stroke:black")
	}

	for j := 0; j <= height; j = j + 10 {
		s.Line(j, 0, j, height, "fill:black;stroke:gray;stroke-dasharray:2")
	}

	s.Square(125, 125, 250, "fill:blue")
	s.End()
}

func SingleWomanSVG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(width, height)

	for i := 0; i <= width; i = i + 50 {
		s.Line(0, i, width, i, "fill:black;stroke:black")
	}

	for j := 0; j <= height; j = j + 10 {
		s.Line(j, 0, j, height, "fill:black;stroke:gray;stroke-dasharray:2")
	}

	s.Circle(250, 250, 125, "fill:red")
	s.End()
}

func OneKidSVG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(width, height)

	for i := 0; i <= width; i = i + 50 {
		s.Line(0, i, width, i, "fill:black;stroke:black")
	}

	for j := 0; j <= height; j = j + 10 {
		s.Line(j, 0, j, height, "fill:black;stroke:gray;stroke-dasharray:2")
	}

	/*
		echo '<rect x="'.($myfather*$mygrid-($mygrid/2)).'" y="'.($mygrid*2-($mygrid/2)).'" width="'.$mygrid.'" height="'.$mygrid.'"
		style="fill:blue;stroke:pink;stroke-width:5;fill-opacity:0.1;stroke-opacity:1" />';
	*/

	s.Square(125, 125, 250, "fill:blue")
	s.Line(375, 250, 1875, 250, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Circle(2000, 250, 125, "fill:red")

	s.Line(1125, 250, 1125, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Circle(1125, 875, 125, "fill:green")

	s.End()
}

func TwoKidSVG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(width, height)

	for i := 0; i <= width; i = i + 50 {
		s.Line(0, i, width, i, "fill:black;stroke:black")
	}

	for j := 0; j <= height; j = j + 10 {
		s.Line(j, 0, j, height, "fill:black;stroke:gray;stroke-dasharray:2")
	}

	s.Square(125, 125, 250, "fill:blue")
	s.Line(375, 250, 1875, 250, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Circle(2000, 250, 125, "fill:red")

	s.Line(1125, 250, 1125, 500, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(250, 500, 2000, 500, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(250, 500, 250, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(2000, 500, 2000, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")

	s.Circle(250, 875, 125, "fill:green")
	s.Circle(2000, 875, 125, "fill:green")

	//s.Circle(1125, 875, 125 ,"fill:green")

	s.End()
}

func ThreeKidSVG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(width, height)

	for i := 0; i <= width; i = i + 50 {
		s.Line(0, i, width, i, "fill:black;stroke:black")
	}

	for j := 0; j <= height; j = j + 10 {
		s.Line(j, 0, j, height, "fill:black;stroke:gray;stroke-dasharray:2")
	}

	s.Square(125, 125, 250, "fill:blue")
	s.Line(375, 250, 1875, 250, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Circle(2000, 250, 125, "fill:red")

	s.Line(1125, 250, 1125, 500, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(250, 500, 2000, 500, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(250, 500, 250, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(1125, 500, 1125, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(2000, 500, 2000, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")

	s.Circle(250, 875, 125, "fill:green")
	s.Circle(1125, 875, 125, "fill:green")
	s.Circle(2000, 875, 125, "fill:green")

	//s.Circle(1125, 875, 125 ,"fill:green")

	s.End()
}

func FourKidSVG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(width, height)

	for i := 0; i <= width; i = i + 50 {
		s.Line(0, i, width, i, "fill:black;stroke:black")
	}

	for j := 0; j <= height; j = j + 10 {
		s.Line(j, 0, j, height, "fill:black;stroke:gray;stroke-dasharray:2")
	}

	s.Square(125, 125, 250, "fill:blue")
	s.Line(375, 250, 1875, 250, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Circle(2000, 250, 125, "fill:red")

	s.Line(1125, 250, 1125, 500, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(250, 500, 2000, 500, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(250, 500, 250, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	// s.Line(1125, 500, 1125, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(687, 500, 687, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(1562, 500, 1562, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(2000, 500, 2000, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")

	s.Circle(250, 875, 125, "fill:green")
	// s.Circle(1125, 875, 125, "fill:green")
	s.Circle(687, 875, 125, "fill:green")
	s.Circle(1562, 875, 125, "fill:green")
	s.Circle(2000, 875, 125, "fill:green")

	//s.Circle(1125, 875, 125 ,"fill:green")

	s.End()
}

func FiveKidSVG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(width, height)

	for i := 0; i <= width; i = i + 50 {
		s.Line(0, i, width, i, "fill:black;stroke:black")
	}

	for j := 0; j <= height; j = j + 10 {
		s.Line(j, 0, j, height, "fill:black;stroke:gray;stroke-dasharray:2")
	}

	s.Square(125, 125, 250, "fill:blue")
	s.Line(375, 250, 1875, 250, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Circle(2000, 250, 125, "fill:red")

	s.Line(1125, 250, 1125, 500, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(250, 500, 2000, 500, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(250, 500, 250, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(1125, 500, 1125, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(687, 500, 687, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(1562, 500, 1562, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(2000, 500, 2000, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")

	s.Circle(250, 875, 125, "fill:green")
	s.Circle(1125, 875, 125, "fill:green")
	s.Circle(687, 875, 125, "fill:green")
	s.Circle(1562, 875, 125, "fill:green")
	s.Circle(2000, 875, 125, "fill:green")

	//s.Circle(1125, 875, 125 ,"fill:green")

	s.End()
}
