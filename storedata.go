package main

import (
	//"fmt"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Id struct {
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
  db, err := gorm.Open("postgres", "user=Genograms dbname=genograms password=genograms sslmode=disable")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()

  // Migrate the schema
  db.AutoMigrate(&Id{})

  // 创建
   db.Create(&Id{Name: "王小明", Age: 45, Uid: "A123456789", Father: "王大明", Mother: "陳大美", Address: "台北市中山區大直街70號", Alive: true})
   db.Create(&Id{Name: "王大明", Age: 70, Uid: "A123456780", Address: "台北市中山區大直街70號", Alive: true})
   db.Create(&Id{Name: "陳大美", Age: 65, Uid: "D23456789",  Address: "台北市中山區大直街70號", Alive: true})


  // 读取


  // var id Id

  // db.Find(&id)
  // fmt.Println(id)
  //db.First(&id, 1) //查询id为1的id
  //db.First(&id, "Name = ?", "A123456789")



  // 更新 - 更新product的price为2000
  // db.Model(&product).Update("Price", 2000)

  // 删除 - 删除product
  /*
  db.Delete(&product)

   */
}