package main

import (
	"database/sql"
	"fmt"
	"log"

	//プログラムでは使用しないが、一緒にコンパイルする必要があるため記述
	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	//第一引数に使うDBのドライバ名
	//第二引数のsqlファイルがなければ新しく作成
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	/*
		テーブル作成
	*/
	cmd := `CREATE TABLE IF NOT EXISTS persons(
                name STRING,
                age INT)`
	_, err := Db.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	/*
		データ追加
	*/
	cmd = "INSERT INTO persons (name, age) VALUES (?, ?)"
	_, err = Db.Exec(cmd, "tarou", 20)
	if err != nil {
		log.Fatalln(err)
	}

	/*
		データ更新
	*/
	cmd = "UPDATE persons SET age = ? WHERE name = ?"
	_, err = Db.Exec(cmd, 25, "tarou")
	if err != nil {
		log.Fatalln(err)
	}

	/*
		単体データ取得
	*/
	cmd = "SELECT * FROM persons WHERE age = ?"
	row := Db.QueryRow(cmd, 25)
	var p Person
	//pにデータを代入
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Fatalln(err)
		}
	}
	fmt.Println("単体", p.Name, p.Age)

	/*
		複数データ取得
	*/
	cmd = "SELECT * FROM persons WHERE age = ?"
	rows, errs := Db.Query(cmd, 25)
	if errs != nil {
		log.Fatalln(errs)
	}
	//Queryを使用した場合はClose()関数で閉じる
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	var pp []Person
	for rows.Next() {
		var p2 Person
		err = rows.Scan(&p2.Name, &p2.Age)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p2)
	}

	for _, p := range pp {
		fmt.Println("複数", p.Name, p.Age)
	}

	/*
		データ削除
	*/
	cmd = "DELETE FROM persons WHERE name = ?"
	_, err = Db.Exec(cmd, "tarou")
	if err != nil {
		log.Fatalln(err)
	}

	/*
		sqliteに入る
		sqlite3 example.sql

		テーブルを確認
		.table

		テーブル内のデータ確認
		SELECT * FROM persons;

		終了
		.exit
	*/
}
