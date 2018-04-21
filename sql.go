package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
)

var sqldb *sql.DB

func SQLSetup(con string){
	fmt.Println("Setup DB")
	// "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	db, err := sql.Open("postgres", con)
	if err != nil{
		log.Fatal(err)
	}
	sqldb = db
	fmt.Println("Setup Finished")
}

func sqlid(ids []int)[]person{
	fmt.Println("sql job")
	rows,err := sqldb.Query(`SELECT * FROM vertex WHERE id IN ($1);`,  ids)
	if err != nil{
		log.Fatal(err)
	}
	results := []person{}
	temp := person{}
	for rows.Next(){
		err = rows.Scan(&temp.id, &temp.name, &temp.job, &temp.birthday)
		results = append(results, temp)
	}
	if err != nil{
		log.Fatal(err)
	}
	return results
}

func sqljob(job string)[]person{
	fmt.Println("sql job")
	rows, err := sqldb.Query(`SELECT * FROM vertex WHERE job = $1;`,  job)
	if err != nil{
		log.Fatal(err)
	}
	results := []person{}
	temp := person{}
	for rows.Next(){
		err = rows.Scan(&temp.id, &temp.name, &temp.job, &temp.birthday)
		results = append(results, temp)
	}
	if err != nil{
		log.Fatal(err)
	}
	return results
}

func sqlname(name string)[]person{
	fmt.Println("sql job")
	rows, err := sqldb.Query(`SELECT * FROM vertex WHERE name = $1;`,  name)
	if err != nil{
		log.Fatal(err)
	}
	results := []person{}
	temp := person{}
	for rows.Next(){
		err = rows.Scan(&temp.id, &temp.name, &temp.job, &temp.birthday)
		results = append(results, temp)
	}
	if err != nil{
		log.Fatal(err)
	}
	return results

}

func sqlendorsing(count int)[]person{
	rows, err := sqldb.Query(`SELECT * FROM (SELECT source, COUNT(dest) as count FROM edges GROUP BY source) as q WHERE count = $1;`,  count)
	if err != nil{
		log.Fatal(err)
	}
	results := []person{}
	temp := person{}
	for rows.Next(){
		err = rows.Scan(&temp.id, &temp.name, &temp.job, &temp.birthday)
		results = append(results, temp)
	}
	if err != nil{
		log.Fatal(err)
	}
	return results
}

func sqlEndorsment(id int)[]person{
	fmt.Println("sql job")
	rows, err := sqldb.Query(`SELECT source FROM edges WHERE dest = $1;`,  id)
	if err != nil{
		log.Fatal(err)
	}
	dests := []int{}
	temp := 0
	for rows.Next(){
		err = rows.Scan(&temp)
		dests = append(dests, temp)
	}
	if err != nil{
		log.Fatal(err)
	}

	return sqlid(dests)
}

func sqld2(count int)[]person{
	rows, err := sqldb.Query(`SELECT * FROM (SELECT dest, COUNT(source) as count FROM edges GROUP BY dest) as q WHERE count = $1;`,  count)
	if err != nil{
		log.Fatal(err)
	}
	results := []person{}
	temp := person{}
	for rows.Next(){
		err = rows.Scan(&temp.id, &temp.name, &temp.job, &temp.birthday)
		results = append(results, temp)
	}
	if err != nil{
		log.Fatal(err)
	}
	return results
}

func sqld3(id int)[]person{
	results := []person{}
	rows, err := sqldb.Query(`SELECT * FROM vertex WHERE id IN (SELECT dest FROM edges WHERE source IN (SELECT dest FROM edges WHERE source IN (SELECT dest FROM edges WHERE source = $1)));`, id)
	if err != nil{
		log.Fatal(err)
	}
	temp := person{}
	for rows.Next(){
		err = rows.Scan(&temp.id, &temp.name, &temp.job, &temp.birthday)
		results = append(results, temp)
	}
	if err != nil{
		log.Fatal(err)
	}


	return results
}

func sqld4(id int)[]person{
	results := []person{}
	rows, err := sqldb.Query(`SELECT * FROM vertex WHERE id IN (SELECT dest FROM edges WHERE source IN (SELECT dest FROM edges WHERE source IN (SELECT dest FROM edges WHERE source IN (SELECT dest FROM edges WHERE source = $1))));`, id)
	if err != nil{
		log.Fatal(err)
	}
	temp := person{}
	for rows.Next(){
		err = rows.Scan(&temp.id, &temp.name, &temp.job, &temp.birthday)
		results = append(results, temp)
	}
	if err != nil{
		log.Fatal(err)
	}


	return results
}

func sqld5(id int)[]person{
	results := []person{}
	rows, err := sqldb.Query(`SELECT * FROM vertex WHERE id IN (SELECT dest FROM edges WHERE source IN (SELECT dest FROM edges WHERE source IN (SELECT dest FROM edges WHERE source IN (SELECT dest FROM edges WHERE source IN (SELECT dest FROM edges WHERE source = $1)))));`, id)
	if err != nil{
		log.Fatal(err)
	}
	temp := person{}
	for rows.Next(){
		err = rows.Scan(&temp.id, &temp.name, &temp.job, &temp.birthday)
		results = append(results, temp)
	}
	if err != nil{
		log.Fatal(err)
	}

	return results
}
