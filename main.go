package main

import (
	"database/sql"
	"fmt"
	"github.com/ajstarks/svgo"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

var (
	err        error
	width      = 10000
	height     = 10000
	startx     int
	starty     int
	grid       int
	span       int
	count      int
	children   = count
	father     = "王大明"
	mother     = "陳大美"
	WIDTH      = children*grid + (children-1)*span
	YPARENTORG = startx + grid/2
	YCHILDORG  = YPARENTORG + grid + span
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

	Db, err = sql.Open("postgres", "user=tree1 dbname=tree1 password=tree1")
	if err != nil {
		panic(err)
	}

	row, err := Db.Query("select startx, starty, grid, span from boundary")
	if err != nil {
		return
	}

	for row.Next() {
		err = row.Scan(&startx, &starty, &grid, &span)
		if err != nil {
			return
		}
		log.Println(startx, starty, grid, span)
	}

	// var count int
	rows := Db.QueryRow("SELECT COUNT(*)  FROM idcard WHERE father = '王大明'  AND  mother = '陳大美'")
	err = rows.Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count, YPARENTORG, YCHILDORG)
	/*

		fmt.Println(count)

		switch count {
		case 1:
			http.Handle("/onekidsvg", http.HandlerFunc(OneKidSVG))
		case 2:
			http.Handle("/twokidsvg", http.HandlerFunc(TwoKidSVG))
		case 3:
			http.Handle("/threekidsvg", http.HandlerFunc(ThreeKidSVG))
		case 4:
			http.Handle("/fourkidsvg", http.HandlerFunc(FourKidSVG))
		case 5:
			http.Handle("/fivekidsvg", http.HandlerFunc(FiveKidSVG))
			}

		err = http.ListenAndServe(":8080", nil)
		if err != nil {
				log.Fatal("ListenAndServe:", err)
		}

	*/

	http.Handle("/singlemansvg", http.HandlerFunc(SingleManSVG))
	http.Handle("/singlewomansvg", http.HandlerFunc(SingleWomanSVG))
	http.Handle("/couplesvg", http.HandlerFunc(CoupleSVG))
	http.Handle("/onekidsvg", http.HandlerFunc(OneKidSVG))
	http.Handle("/twokidsvg", http.HandlerFunc(TwoKidSVG))
	http.Handle("/threekidsvg", http.HandlerFunc(ThreeKidSVG))
	http.Handle("/fourkidsvg", http.HandlerFunc(FourKidSVG))
	http.Handle("/fivekidsvg", http.HandlerFunc(FiveKidSVG))
	err = http.ListenAndServe(":2003", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

func CoupleSVG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(WIDTH, height)

	/*
		for i := 0; i <= width; i = i + 50 {
			s.Line(0, i, width, i, "fill:black;stroke:black")
		}

		for j := 0; j <= height; j = j + 10 {
			s.Line(j, 0, j, height, "fill:black;stroke:gray;stroke-dasharray:2")
		}

	*/

	s.Square(125, 125, 250, "fill:blue")
	s.Line(375, 250, 1875, 250, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Circle(2000, 250, 125, "fill:red")
	s.End()
}

func SingleManSVG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	s := svg.New(w)
	s.Start(WIDTH, height)

	/*
		for i := 0; i <= width; i = i + 50 {
			s.Line(0, i, width, i, "fill:black;stroke:black")
		}

		for j := 0; j <= height; j = j + 10 {
			s.Line(j, 0, j, height, "fill:black;stroke:gray;stroke-dasharray:2")
		}

	*/

	s.Square(125, 125, 250, "fill:blue")
	s.End()
}

func SingleWomanSVG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(WIDTH, height)

	/*
		for i := 0; i <= width; i = i + 50 {
			s.Line(0, i, width, i, "fill:black;stroke:black")
		}

		for j := 0; j <= height; j = j + 10 {
			s.Line(j, 0, j, height, "fill:black;stroke:gray;stroke-dasharray:2")
		}

	*/

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
	s.Line(startx+grid, starty+grid/2, startx+grid+span, starty+grid/2, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(startx+grid+span/2, starty+grid/2, startx+grid+span/2, starty+grid/2+span, "stroke:black;stroke-width:20;stroke-linecap:butt")
	for row.Next() {
		err = row.Scan(&id.Id, &id.Name, &id.Age, &id.Gender,
			&id.IdNumber, &id.Father, &id.Mother, &id.Address)
		if err != nil {
			return
		}
		if id.Id == 1 {
			s.Text(startx+(grid/4)*3, starty+grid/2, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Square(startx+grid/2, starty+grid/4, grid/2, "fill:none;stroke:black")
		} else if id.Id == 2 {
			s.Text(startx+(grid/4)*5+span, starty+grid/2, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Circle(startx+(grid/4)*5+span, starty+grid/2, grid/4, "fill:none;stroke:black")
		} else if id.Id == 3 {
			s.Text(startx+grid+span/2, starty+(grid/4)*3+span, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			if id.Gender == "男" {
				s.Square(startx+grid+span/4, starty+grid/2+span, grid/2, "fill:none;stroke:black")
			} else {
				s.Circle(startx+grid+span/2, starty+grid/2+span+grid/4, grid/4, "fill:none;stroke:black")
			}
		}
	}
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
	s.Line(startx+grid, starty+grid/2, startx+grid+span, starty+grid/2, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(startx+grid+span/2, starty+grid/2, startx+grid+span/2, starty+grid/2+span/2, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(startx+(grid/4)*3, starty+grid/2+span/2, startx+grid+span+grid/4, starty+grid/2+span/2, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(startx+(grid/4)*3, starty+grid/2+span/2, startx+(grid/4)*3, starty+grid/2+span, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(startx+grid+span+grid/4, starty+grid/2+span/2, startx+grid+span+grid/4, starty+grid/2+span, "stroke:black;stroke-width:20;stroke-linecap:butt")
	for row.Next() {
		err = row.Scan(&id.Id, &id.Name, &id.Age, &id.Gender,
			&id.IdNumber, &id.Father, &id.Mother, &id.Address)
		if err != nil {
			return
		}
		if id.Id == 1 {
			s.Text(startx+(grid/4)*3, starty+grid/2, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Square(startx+grid/2, starty+grid/4, grid/2, "fill:none;stroke:black")
		} else if id.Id == 2 {
			s.Text(startx+(grid/4)*5+span, starty+grid/2, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Circle(startx+(grid/4)*5+span, starty+grid/2, grid/4, "fill:none;stroke:black")
		} else if id.Id == 3 {
			s.Text((grid/4)*3, starty+(grid/4)*3+span, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			if id.Gender == "男" {
				s.Square(startx+grid/2, starty+grid/2+span, grid/2, "fill:none;stroke:black")
			} else {
				s.Circle(startx+(grid/4)*5+span, starty+grid/2+span+grid/4, grid/4, "fill:none;stroke:black")
			}
		} else if id.Id == 4 {
			s.Text(startx+grid+span+grid/4, starty+(grid/4)*3+span, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			if id.Gender == "男" {
				s.Square(startx+grid+span, starty+grid/2+span, grid/2, "fill:none;stroke:black")
			} else {
				s.Circle(startx+(grid/4)*5+span, starty+grid/2+span+grid/4, grid/4, "fill:none;stroke:black")
			}
		}
	}
	s.End()
}

func ThreeKidSVG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(WIDTH, height)

	/*
		for i := 0; i <= width; i = i + 50 {
			s.Line(0, i, width, i, "fill:black;stroke:black")
		}

		for j := 0; j <= height; j = j + 10 {
			s.Line(j, 0, j, height, "fill:black;stroke:gray;stroke-dasharray:2")
		}
	*/

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

	// s.Square(125, 125, 250, "fill:blue")
	s.Line(375, 250, 1875, 250, "stroke:black;stroke-width:20;stroke-linecap:butt")
	// s.Circle(2000, 250, 125, "fill:red")

	s.Line(1125, 250, 1125, 500, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(250, 500, 2000, 500, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(250, 500, 250, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(1125, 500, 1125, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(2000, 500, 2000, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")

	// s.Circle(250, 875, 125, "fill:green")
	// s.Circle(1125, 875, 125, "fill:green")
	// s.Circle(2000, 875, 125, "fill:green")

	for row.Next() {
		err = row.Scan(&id.Id, &id.Name, &id.Age, &id.Gender,
			&id.IdNumber, &id.Father, &id.Mother, &id.Address)
		if err != nil {
			return
		}
		//log.Println(id.Id, id.Name, id.Age, id.Gender,
		//	id.IdNumber, id.Father, id.Mother, id.Address)
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
			s.Text(1125, 875, id.Name, "text-anchor:middle;font-size:30px")
			s.Text(1125, 905, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Text(1125, 935, id.IdNumber, "text-anchor:middle;font-size:30px")
			if id.Gender == "男" {
				s.Square(1000, 750, 250, "fill:none;stroke:black")
			} else {
				s.Circle(1125, 875, 125, "fill:none;stroke:black")
			}

		} else if id.Id == 5 {
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
	s.End()
}

func FourKidSVG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(WIDTH, height)
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
	s.Line(687, 500, 687, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(1562, 500, 1562, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(2000, 500, 2000, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	for row.Next() {
		err = row.Scan(&id.Id, &id.Name, &id.Age, &id.Gender,
			&id.IdNumber, &id.Father, &id.Mother, &id.Address)
		if err != nil {
			return
		}
		if id.Id == 1 {
			s.Text(225, 250, id.Name, "text-anchor:middle;font-size:30px")
			s.Text(225, 280, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Text(225, 310, id.IdNumber, "text-anchor:middle;font-size:30px")
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
			//s.Circle(1125, 875, 125, "fill:green")
			s.Text(687, 875, id.Name, "text-anchor:middle;font-size:30px")
			s.Text(687, 905, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Text(687, 935, id.IdNumber, "text-anchor:middle;font-size:30px")
			if id.Gender == "男" {
				s.Square(562, 750, 250, "fill:none;stroke:black")
			} else {
				s.Circle(687, 875, 125, "fill:none;stroke:black")
				// s.Circle(1125, 875, 125, "fill:none;stroke:black")
			}

		} else if id.Id == 5 {
			// log.Println(id.Age)
			// log.Println(string(id.Age))
			// log.Println(id.Gender)
			// s.Circle(2000, 875, 125, "fill:green")
			s.Text(1562, 875, id.Name, "text-anchor:middle;font-size:30px")
			s.Text(1562, 905, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Text(1562, 935, id.IdNumber, "text-anchor:middle;font-size:30px")
			if id.Gender == "男" {
				s.Square(1437, 750, 250, "fill:none;stroke:black")
			} else {
				//s.Circle(1562, 875, 125, "fill:green")
				s.Circle(1562, 875, 125, "fill:none;stroke:black")
			}

		} else if id.Id == 6 {
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
	// s.Circle(250, 875, 125, "fill:green")
	// s.Circle(1125, 875, 125, "fill:green")
	// s.Circle(687, 875, 125, "fill:green")
	// s.Circle(1562, 875, 125, "fill:green")
	// s.Circle(2000, 875, 125, "fill:green")

	//s.Circle(1125, 875, 125 ,"fill:green")

	s.End()
}

func FiveKidSVG(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	s := svg.New(w)
	s.Start(WIDTH, height)

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

	/*
		for i := 0; i <= width; i = i + 50 {
			s.Line(0, i, width, i, "fill:black;stroke:black")
		}

		for j := 0; j <= height; j = j + 10 {
			s.Line(j, 0, j, height, "fill:black;stroke:gray;stroke-dasharray:2")
		}

	*/

	// s.Square(125, 125, 250, "fill:blue")
	s.Line(375, 250, 1875, 250, "stroke:black;stroke-width:20;stroke-linecap:butt")
	// s.Circle(2000, 250, 125, "fill:red")

	s.Line(1125, 250, 1125, 500, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(250, 500, 2000, 500, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(250, 500, 250, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(1125, 500, 1125, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(687, 500, 687, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(1562, 500, 1562, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")
	s.Line(2000, 500, 2000, 750, "stroke:black;stroke-width:20;stroke-linecap:butt")

	// s.Circle(250, 875, 125, "fill:green")
	// s.Circle(1125, 875, 125, "fill:green")
	// s.Circle(687, 875, 125, "fill:green")
	// s.Circle(1562, 875, 125, "fill:green")
	// s.Circle(2000, 875, 125, "fill:green")

	// s.Circle(1125, 875, 125 ,"fill:green")

	for row.Next() {
		err = row.Scan(&id.Id, &id.Name, &id.Age, &id.Gender,
			&id.IdNumber, &id.Father, &id.Mother, &id.Address)
		if err != nil {
			return
		}
		//log.Println(id.Id, id.Name, id.Age, id.Gender,
		//	id.IdNumber, id.Father, id.Mother, id.Address)
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
			s.Text(687, 875, id.Name, "text-anchor:middle;font-size:30px")
			s.Text(687, 905, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Text(687, 935, id.IdNumber, "text-anchor:middle;font-size:30px")
			if id.Gender == "男" {
				s.Square(562, 750, 250, "fill:none;stroke:black")
			} else {
				s.Circle(687, 875, 125, "fill:none;stroke:black")
			}

		} else if id.Id == 5 {
			// log.Println(id.Age)
			// log.Println(string(id.Age))
			// log.Println(id.Gender)
			// s.Circle(2000, 875, 125, "fill:green")
			s.Text(1125, 875, id.Name, "text-anchor:middle;font-size:30px")
			s.Text(1125, 905, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Text(1125, 935, id.IdNumber, "text-anchor:middle;font-size:30px")
			if id.Gender == "男" {
				s.Square(1000, 750, 250, "fill:none;stroke:black")
			} else {
				s.Circle(1125, 875, 125, "fill:none;stroke:black")
			}

		} else if id.Id == 6 {
			// log.Println(id.Age)
			// log.Println(string(id.Age))
			// log.Println(id.Gender)
			// s.Circle(2000, 875, 125, "fill:green")
			s.Text(1562, 875, id.Name, "text-anchor:middle;font-size:30px")
			s.Text(1562, 905, strconv.Itoa(id.Age), "text-anchor:middle;font-size:30px")
			s.Text(1562, 935, id.IdNumber, "text-anchor:middle;font-size:30px")
			if id.Gender == "男" {
				s.Square(1875, 750, 250, "fill:none;stroke:black")
			} else {
				s.Circle(1562, 875, 125, "fill:none;stroke:black")
			}

		} else if id.Id == 7 {
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
	s.End()
}
