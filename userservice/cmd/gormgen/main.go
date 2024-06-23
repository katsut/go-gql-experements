package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:password@tcp(db:3306)/root?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// インスタンス化
	g := gen.NewGenerator(gen.Config{
		OutPath: "dao/query",
		Mode:    gen.WithDefaultQuery,
	})

	// DB接続を設定
	g.UseDB(db)

	// テーブルからモデルを生成
	g.ApplyBasic(g.GenerateModel("users"))

	// ファイルを出力
	g.Execute()
}
