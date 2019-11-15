package main

import (
	"database/sql"
	"github.com/ajstarks/svgo"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

var (
	height = 10000
	width  = 10000
	/*
		Id       int
		Name     string
		Age      int
		Gender   string
		IdNumber string
		Father   string
		Mother   string
		Couple   string
		Address  string

	*/
)

type idcard struct {
	Id       int
	Name     string
	Age      int
	Gender   string
	IdNumber string
	Father   string
	Mother   string
	Couple   string
	Address  string
}

var Db *sql.DB

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

	var err error
	Db, err = sql.Open("postgres", "user=tree1 dbname=tree1 password=tree1")
	if err != nil {
		panic(err)
	}

	row, err := Db.Query("select id, name, age , gender, id_number, father, mother, address from idcard")
	if err != nil {
		return
	}

	var id idcard

	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(width, height)
	s.Line(375, 250, 1875, 250, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(1125, 250, 1125, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	for row.Next() {
		err = row.Scan(&id.Id, &id.Name, &id.Age, &id.Gender,
			&id.IdNumber, &id.Father, &id.Mother, &id.Address)
		if err != nil {
			return
		}
		log.Println(id.Id, id.Name, id.Age, id.Gender,
			id.IdNumber, id.Father, id.Mother, id.Address)
		if id.Id == 1 {
			s.Text(225, 250, id.Name, "text-anchor:middle;font-size:30px")
			// log.Println(id.Age)
			s.Text(225, 280, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Text(225, 310, id.IdNumber, "text-anchor:middle;font-size:30px")
			// log.Println(id.Gender)
			s.Square(125, 125, 250, "fill:none;stroke:black")
		} else if id.Id == 2 {
			s.Text(1975, 250, id.Name, "text-anchor:middle;font-size:30px")
			// log.Println(id.Age)
			s.Text(1975, 280, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Text(1975, 310, id.IdNumber, "text-anchor:middle;font-size:30px")
			// log.Println(id.Gender)
			s.Circle(2000, 250, 125, "fill:none;stroke:black")
			// log.Println(id.Id)
		} else if id.Id == 3 {
			// log.Println(id.Age)
			// log.Println(string(id.Age))
			// log.Println(id.Gender)
			s.Text(1100, 875, id.Name, "text-anchor:middle;font-size:30px")
			s.Text(1100, 905, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Text(1100, 935, id.IdNumber, "text-anchor:middle;font-size:30px")
			if id.Gender == "男" {
				s.Square(1000, 750, 250, "fill:none;stroke:black")
			} else {
				s.Circle(1125, 875, 125, "fill:none;stroke:black")
			}

		}
	}

	//s.Text(250, 250, id.Name)

	/*
		for i := 0; i <= width; i = i + 50 {
			s.Line(0, i, width, i, "fill:black;stroke:black")
		}

		for j := 0; j <= height; j = j + 10 {
			s.Line(j, 0, j, height, "fill:black;stroke:gray;stroke-dasharray:2")
		}

	*/

	/*
		echo '<rect x="'.($myfather*$mygrid-($mygrid/2)).'" y="'.($mygrid*2-($mygrid/2)).'" width="'.$mygrid.'" height="'.$mygrid.'"
		style="fill:blue;stroke:pink;stroke-width:5;fill-opacity:0.1;stroke-opacity:1" />';
	*/

	// s.Square(125, 125, 250, "fill:none;stroke:black")
	// s.Line(375, 250, 1875, 250, "stroke:black;stroke-width:20;stroke-linecap:butt")
	// s.Circle(2000, 250, 125, "fill:none;stroke:black")

	// s.Line(1125, 250, 1125, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	// s.Circle(1125, 875, 125, "fill:none;stroke:black")

	s.End()
}

func TwoKidSVG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(width, height)

	var err error
	Db, err = sql.Open("postgres", "user=tree1 dbname=tree1 password=tree1")
	if err != nil {
		panic(err)
	}

	row, err := Db.Query("select id, name, age , gender, id_number, father, mother, address from idcard")
	if err != nil {
		return
	}

	var id idcard

	s.Line(375, 250, 1875, 250, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(1125, 250, 1125, 500, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(250, 500, 2000, 500, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(250, 500, 250, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(2000, 500, 2000, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	for row.Next() {
		err = row.Scan(&id.Id, &id.Name, &id.Age, &id.Gender,
			&id.IdNumber, &id.Father, &id.Mother, &id.Address)
		if err != nil {
			return
		}
		log.Println(id.Id, id.Name, id.Age, id.Gender,
			id.IdNumber, id.Father, id.Mother, id.Address)
		if id.Id == 1 {
			s.Text(225, 250, id.Name, "text-anchor:middle;font-size:30px")
			// log.Println(id.Age)
			s.Text(225, 280, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Text(225, 310, id.IdNumber, "text-anchor:middle;font-size:30px")
			// log.Println(id.Gender)
			s.Square(125, 125, 250, "fill:none;stroke:black")
		} else if id.Id == 2 {
			s.Text(1975, 250, id.Name, "text-anchor:middle;font-size:30px")
			// log.Println(id.Age)
			s.Text(1975, 280, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Text(1975, 310, id.IdNumber, "text-anchor:middle;font-size:30px")
			// log.Println(id.Gender)
			s.Circle(2000, 250, 125, "fill:none;stroke:black")
			// log.Println(id.Id)
		} else if id.Id == 3 {
			// log.Println(id.Age)
			// log.Println(string(id.Age))
			// log.Println(id.Gender)
			// s.Circle(250, 875, 125, "fill:green")
			s.Text(250, 875, id.Name, "text-anchor:middle;font-size:30px")
			s.Text(250, 905, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Text(250, 935, id.IdNumber, "text-anchor:middle;font-size:30px")
			if id.Gender == "男" {
				s.Square(125, 750, 250, "fill:none;stroke:black")
			} else {
				s.Circle(250, 875, 125, "fill:none;stroke:black")
			}
		} else if id.Id == 4 {
			// log.Println(id.Age)
			// log.Println(string(id.Age))
			// log.Println(id.Gender)
			// s.Circle(2000, 875, 125, "fill:green")
			s.Text(2000, 875, id.Name, "text-anchor:middle;font-size:30px")
			s.Text(2000, 905, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Text(2000, 935, id.IdNumber, "text-anchor:middle;font-size:30px")
			if id.Gender == "男" {
				s.Square(1875, 750, 250, "fill:none;stroke:black")
			} else {
				s.Circle(2000, 875, 125, "fill:none;stroke:black")
			}

		}
	}

	/*
		for i := 0; i <= width; i = i + 50 {
			s.Line(0, i, width, i, "fill:black;stroke:black")
		}

		for j := 0; j <= height; j = j + 10 {
			s.Line(j, 0, j, height, "fill:black;stroke:gray;stroke-dasharray:2")
		}

	*/

	// s.Square(125, 125, 250, "fill:blue")
	// s.Line(375, 250, 1875, 250, "stroke:black;stroke-width:20;stroke-linecap:butt")
	// s.Circle(2000, 250, 125, "fill:red")

	// s.Line(1125, 250, 1125, 500, "stroke:black;stroke-width:20;stroke-linecap:butt")
	// s.Line(250, 500, 2000, 500, "stroke:black;stroke-width:20;stroke-linecap:butt")
	// s.Line(250, 500, 250, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	// s.Line(2000, 500, 2000, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")

	// s.Circle(250, 875, 125, "fill:green")
	// s.Circle(2000, 875, 125, "fill:green")

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
