package main

import (
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// マイグレーション用のディレクトリパス
	migrationPath := "file:///Users/koshimaruryoma/my-dev/go/migration"

	// マイグレーション用のDBのDSN
	dsn := "mysql://root:password@tcp(127.0.0.1:3306)/task_app?charset=utf8mb4&parseTime=True&loc=Local"

	// migrateのインスタンスを作成
	m, err := migrate.New(migrationPath, dsn)
	if err != nil {
		panic(err)
	}

	// マイグレーションを実行
	err = m.Up()
	if err != nil {
		panic(err)
	}
}
