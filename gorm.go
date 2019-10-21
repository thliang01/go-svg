package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Idcard struct {
	gorm.Model
	Name    string `gorm:"size:32;not null"`
	Age     int
	Uid     string `gorm:"not null;size:10"`
	Father  string `gorm:"size:32"`
	Mother  string `gorm:"size:32"`
	Couple  string `gorm:"size:32"`
	Address string
	Alive   bool
}

func main() {
	db, err := gorm.Open("postgres", "user=Genograms dbname=genograms password=genograms")
	if err != nil {
		panic("connect failed")
	}
	defer db.Close()

	// 自动迁移模式
	db.AutoMigrate(&Idcard{})

	// 创建
	db.Create(&Idcard{Name: "王小明", Age: 45, Uid: "A123456789", Father: "王大明", Mother: "陳大美", Address: "台北市中山區大直街70號", Alive: true})
	// db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	// var product Product
	// db.First(&product, 1) // 查询id为1的product
	// db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
	var id Idcard
	db.First(&id, 1)

	// 更新 - 更新product的price为2000
	// db.Model(&id).Update("Price", 2500)

	// 删除 - 删除product
	db.Delete(&id)
}
