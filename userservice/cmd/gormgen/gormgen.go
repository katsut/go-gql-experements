package main

import (
	"log"

	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("DB_CONNECTION_STRING")
	if dsn == "" {
		log.Fatal("DB_CONNECTION_STRING environment variable is not set")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// インスタンス化
	g := gen.NewGenerator(gen.Config{
		OutPath: "/app/internal/dao/query",
		Mode:    gen.WithDefaultQuery,
	})

	// DB接続を設定
	g.UseDB(db)

	// テーブルからモデルを生成
	g.ApplyBasic(g.GenerateModel("users"))

	// ファイルを出力
	g.Execute()
}
