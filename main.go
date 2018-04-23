package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

// go get github.com/lib/pq
// go get github.com/johnnadratowski/golang-neo4j-bolt-driver   - Didn't work
// go get gopkg.in/cq.v1

var TestPeople []int
var TestJobs []string
var TestNames []string
var TestCounts []int
var TestBirthdays []string
var EndorsmentCount []int

func main(){

	SQLSetup("user=postgres dbname=postgres host=psql sslmode=disable")
	//SetupNeo4j()
	SetupTestVariables()
	PrintTestVariables()
	var dur []time.Duration


	fmt.Println()
	fmt.Println(" ###################################")
	fmt.Println(" ######## SQL Experiments ##########")
	fmt.Println(" ###################################")
	fmt.Println()

	fmt.Println("SQL - Experiement: All person that a person endorses (Depth one)")
	dur = ExperimentIntQuery(TestPeople, sqld1)
	PrintStats(dur)

	fmt.Println("SQL - Experiement: Endorsements of depth two.")
	dur = ExperimentIntQuery(TestPeople, sqld2)
	PrintStats(dur)

	fmt.Println("SQL - Experiement: Endorsements of depth three.")
	dur = ExperimentIntQuery(TestPeople, sqld3)
	PrintStats(dur)

	fmt.Println("SQL - Experiement: Endorsements of depth four.")
	dur = ExperimentIntQuery(TestPeople, sqld4)
	PrintStats(dur)

	fmt.Println("SQL - Experiement: Endorsements of depth five.")
	dur = ExperimentIntQuery(TestPeople, sqld5)
	PrintStats(dur)


	fmt.Println("SQL - Experiement: All persons that has a specific job")
	dur = ExperimentStringQuery(TestJobs, sqljob)
	PrintStats(dur)

	fmt.Println("SQL - Experiement: All persons that has a speicifc name")
	dur = ExperimentStringQuery(TestNames, sqlname)
	PrintStats(dur)

	fmt.Println("SQL - Experiement:  All persons that has a specific birthday.")
	dur = ExperimentStringQuery(TestBirthdays, sqlbirth)
	PrintStats(dur)

	fmt.Println("SQL - Experiement: All persons that has a specific ammount of people they endorse.")
	dur = ExperimentIntQuery(EndorsmentCount, sqlendorsing)
	PrintStats(dur)

	fmt.Println("SQL - Experiement:  All persons that has a specific amount of endorsments..")
	dur = ExperimentIntQuery(EndorsmentCount, sqlEndorsment)
	PrintStats(dur)

	////////////////// NEO4J
	fmt.Println()
	fmt.Println(" ###################################")
	fmt.Println(" ######## NEO4J Experiments ########")
	fmt.Println(" ###################################")
	fmt.Println()

	fmt.Println("Neo4j - Experiement: All person that a person endorses (Depth one)")
	dur = ExperimentIntQuery(TestPeople, Neod1)
	PrintStats(dur)

	fmt.Println("Neo4j - Experiement: Endorsements of depth two.")
	dur = ExperimentIntQuery(TestPeople, Neod2)
	PrintStats(dur)

	fmt.Println("Neo4j - Experiement: Endorsements of depth three.")
	dur = ExperimentIntQuery(TestPeople, Neod3)
	PrintStats(dur)

	fmt.Println("Neo4j - Experiement: Endorsements of depth four.")
	dur = ExperimentIntQuery(TestPeople, Neod4)
	PrintStats(dur)

	fmt.Println("Neo4j - Experiement: Endorsements of depth five.")
	dur = ExperimentIntQuery(TestPeople, Neod5)
	PrintStats(dur)


	fmt.Println("Neo4j - Experiement: All persons that has a specific job")
	dur = ExperimentStringQuery(TestJobs, NeoJob)
	PrintStats(dur)

	fmt.Println("Neo4j - Experiement: All persons that has a speicifc name")
	dur = ExperimentStringQuery(TestNames, NeoName)
	PrintStats(dur)

	fmt.Println("Neo4j - Experiement:  All persons that has a specific birthday.")
	dur = ExperimentStringQuery(TestBirthdays, NeoBirth)
	PrintStats(dur)

	fmt.Println("Neo4j - Experiement: All persons that has a specific ammount of people they endorse.")
	dur = ExperimentIntQuery(EndorsmentCount, NeoEndorsing)
	PrintStats(dur)

	fmt.Println("Neo4j - Experiement:  All persons that has a specific amount of endorsments..")
	dur = ExperimentIntQuery(EndorsmentCount, NeoEndorsment)
	PrintStats(dur)
	


}

func PrintStats(dur []time.Duration){
	for _,x := range dur {
		fmt.Print( "   "+x.String())
	}
	fmt.Print("\n")
	fmt.Print("Average: ")
	fmt.Println(GetAverage(dur))
	fmt.Print("Median: ")
	fmt.Println(GetMedian(dur))
}

func GetAverage(durs []time.Duration)time.Duration{
	var total int64 = 0
	for _, x := range durs{
		total += x.Nanoseconds()
	}
	return time.Duration(total/int64(len(durs)))
}

func GetMedian(durs []time.Duration)time.Duration{
	sort.Slice(durs, func(i, j int) bool { return durs[i].Nanoseconds() < durs[j].Nanoseconds()})
	return durs[len(durs)/2]
}



type person struct {
	id int
	name string
	job string
	birthday string
}


func SetupTestVariables(){
	TestNames = []string{"Hertha Bergdorf", "Fiona Flueck", "Maybell Lettieri", "Phillis Stoneham", "Gianna Papania", "Stormy Christion", "Rima Dipaola", "Clyde Kotter"}
	TestJobs = []string{"Wharf Attendant","ewing-Machine Operator","Fish-Bin Tender", "Boiler", "Production Clerk", "Cloth Grader", "Batch-Still Operator"}
	TestBirthdays = []string{"1940-01-16", "1943-02-23", "1941-06-13", "1966-10-16", "1980-01-16", "1994-02-17"}
	for i := 20; i>0; i--{
		TestPeople = append(TestPeople, rand.Intn(500000))
	}
	TestCounts = []int{10, 20, 30, 40, 50}
	EndorsmentCount = []int{ rand.Intn(50),rand.Intn(50)  }
}

func PrintTestVariables(){
	fmt.Print("TEST DATA PRINT: \nNames:\n\t")
	for _, x := range TestNames{
		fmt.Print(x)
		fmt.Print(" ")
	}
	fmt.Print("\nJobs:\n \t")
	for _, x := range TestJobs{
		fmt.Print(x)
		fmt.Print("  ")
	}
	fmt.Print("\nBirthdays:\n \t")
	for _, x := range TestBirthdays{
		fmt.Print(x)
		fmt.Print("  ")
	}
	fmt.Print("\nRandom 20 people's id:\n \t")
	for _, x := range TestPeople{
		fmt.Print(x)
		fmt.Print(" ")
	}
	fmt.Print("\nEndorsements Counts:\n \t")
	for _, x := range EndorsmentCount{
		fmt.Print(x)
		fmt.Print(" ")
	}


	fmt.Print("\n\nStarting Experiments: \n\n")
}