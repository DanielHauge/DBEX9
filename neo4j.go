package main

import (
	"database/sql"
	"log"
	_ "gopkg.in/cq.v1"
	"fmt"

)

var neo4j *sql.DB

func SetupNeo4j(){
	db, err := sql.Open("neo4j-cypher", "http://192.168.33.10:7474")
	if err != nil {
		panic(err)
	}
	neo4j = db

	fmt.Println("Setup Finished")
}



func NeoEndorsing(count int)[]person{
	result := []person{}
	stmt, err := neo4j.Prepare(`MATCH ()-[r]->(n) WITH n, count(r) as rel_cnt WHERE rel_cnt = {0} return ID(n), n.name, n.job, n.birthday;`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(count)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var ids int
	var name string
	var job string
	var birthday string
	for rows.Next() {
		err := rows.Scan(&ids, &name, &job, &birthday)
		result = append(result, person{ids, name, job, birthday})
		if err != nil {
			log.Fatal(err)
		}

	}
	return result
}

func NeoEndorsment(count int)[]person{
	result := []person{}
	stmt, err := neo4j.Prepare(`MATCH (n)-[r]->() WITH n, count(r) as rel_cnt WHERE rel_cnt = {0} return ID(n), n.name, n.job, n.birthday;`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(count)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var ids int
	var name string
	var job string
	var birthday string
	for rows.Next() {
		err := rows.Scan(&ids, &name, &job, &birthday)
		result = append(result, person{ids, name, job, birthday})
		if err != nil {
			log.Fatal(err)
		}

	}
	return result
}

func NeoJob(jobin string)[]person{
	result := []person{}
	stmt, err := neo4j.Prepare(`MATCH (x:Person) WHERE x.job={0} return ID(x), x.name, x.job, x.birthday`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(jobin)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var ids int
	var name string
	var job string
	var birthday string
	for rows.Next() {
		err := rows.Scan(&ids, &name, &job, &birthday)
		result = append(result, person{ids, name, job, birthday})
		if err != nil {
			log.Fatal(err)
		}

	}
	return result
}

func NeoName(namein string)[]person{
	result := []person{}
	stmt, err := neo4j.Prepare(`MATCH (x:Person) WHERE x.name={0} return ID(x), x.name, x.job, x.birthday`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(namein)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var ids int
	var name string
	var job string
	var birthday string
	for rows.Next() {
		err := rows.Scan(&ids, &name, &job, &birthday)
		result = append(result, person{ids, name, job, birthday})
		if err != nil {
			log.Fatal(err)
		}

	}
	return result
}

func NeoBirth(birth string)[]person{
	result := []person{}
	stmt, err := neo4j.Prepare(`MATCH (x:Person) WHERE x.birthday={0} return ID(x), x.name, x.job, x.birthday`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(birth)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var ids int
	var name string
	var job string
	var birthday string
	for rows.Next() {
		err := rows.Scan(&ids, &name, &job, &birthday)
		result = append(result, person{ids, name, job, birthday})
		if err != nil {
			log.Fatal(err)
		}

	}
	return result
}


func Neod1(id int)[]person{
	result := []person{}
	stmt, err := neo4j.Prepare(`MATCH (x:Person)-[:ENDORSES]->(other) WHERE ID(x)= {0} return ID(other), other.name, other.job, other.birthday`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var ids int
	var name string
	var job string
	var birthday string
	for rows.Next() {
		err := rows.Scan(&ids, &name, &job, &birthday)
		result = append(result, person{ids, name, job, birthday})
		if err != nil {
			log.Fatal(err)
		}

	}
	return result
}

func Neod2(id int)[]person{
	result := []person{}
	stmt, err := neo4j.Prepare(`MATCH (x:Person)-[:ENDORSES]->()-[:ENDORSES]->(other) WHERE ID(x)= {0} return ID(other), other.name, other.job, other.birthday`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var ids int
	var name string
	var job string
	var birthday string
	for rows.Next() {
		err := rows.Scan(&ids, &name, &job, &birthday)
		result = append(result, person{ids, name, job, birthday})
		if err != nil {
			log.Fatal(err)
		}

	}
	return result
}

func Neod3(id int)[]person{
	result := []person{}
	stmt, err := neo4j.Prepare(`MATCH (x:Person)-[:ENDORSES]->()-[:ENDORSES]->()-[:ENDORSES]->(other) WHERE ID(x)= {0} return ID(other), other.name, other.job, other.birthday`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var ids int
	var name string
	var job string
	var birthday string
	for rows.Next() {
		err := rows.Scan(&ids, &name, &job, &birthday)
		result = append(result, person{ids, name, job, birthday})
		if err != nil {
			log.Fatal(err)
		}

	}
	return result
}

func Neod4(id int)[]person{
	result := []person{}
	stmt, err := neo4j.Prepare(`MATCH (x:Person)-[:ENDORSES]->()-[:ENDORSES]->()-[:ENDORSES]->()-[:ENDORSES]->(other) WHERE ID(x)= {0} return ID(other), other.name, other.job, other.birthday`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var ids int
	var name string
	var job string
	var birthday string
	for rows.Next() {
		err := rows.Scan(&ids, &name, &job, &birthday)
		result = append(result, person{ids, name, job, birthday})
		if err != nil {
			log.Fatal(err)
		}

	}
	return result
}

func Neod5(id int)[]person{
	result := []person{}
	stmt, err := neo4j.Prepare(`MATCH (x:Person)-[:ENDORSES]->()-[:ENDORSES]->()-[:ENDORSES]->()-[:ENDORSES]->()-[:ENDORSES]->(other) WHERE ID(x)= {0} return ID(other), other.name, other.job, other.birthday`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var ids int
	var name string
	var job string
	var birthday string
	for rows.Next() {
		err := rows.Scan(&ids, &name, &job, &birthday)
		result = append(result, person{ids, name, job, birthday})
		if err != nil {
			log.Fatal(err)
		}

	}
	return result
}
