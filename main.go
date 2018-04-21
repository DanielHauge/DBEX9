package main


// go get github.com/lib/pq

func main(){

	SQLSetup("user=postgres dbname=postgres host=192.168.33.10 sslmode=disable")

	sqljob("Driver")

}



type person struct {
	id int
	name string
	job string
	birthday string
}