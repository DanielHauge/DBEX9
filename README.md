# Database Exercise 9 (Technical Comparrison: SQL vs Graph)
This repository is a exercise project for Software development (PBA) Database course. Daniel (cph-dh136).

This exercise is based on this ressource: [Here](https://github.com/datsoftlyngby/soft2018spring-databases-teaching-material/blob/master/assignments/Neo4J%20Exercise.ipynb)

## Introduction
### Initial Problem statement
SQL based databases vs Graph based databases, which is best?.

Does Graph based databases have an advantage over SQL based databases in regards to searching or is it reversed, and if so, in under which circumstances?.

### Hypothesis
Considering the question at hand. We need to consider the circumstances and what is searched for. The hypothesis is that, Graph based databases are alot faster at searching for nested relations in deeper layers, where as SQL based databases might pull slightly ahead on straight value searches and less deep layers. Reasoning for this hypothesis, is that graph databases seems very effective at traversing data from a known location but might not be as effective to search for elements with a specific value, where as in contrast SQL based databases might be faster to run through lots of relations(rows) quickly, but might be very ineffective when it comes to searching for relations in deeper layers(nested), because it does not traverse the data the same way a graph database does.

### Tests / Experiments
To test the hypothesis, we can setup an experiment.

The experiment will be as follows: From a pre-existing dataset, which consists of people with jobs, names and connections(endorsments) to other people). Meassure average and median query runtime on predetermined queries for both database systems. The queries are as follows:
- All person that a person endorses (Depth one)
- All person that are endorsed by endorsed persons of a person (Depth two)
- Endorsements of depth three
- Endorsements of depth four.
- Endorsements of depth five.
- All persons that has a specific job
- All persons that has a specific name
- All persons that has a specific birthday.
- All persons that has a specific ammount of people they endorse.
- All persons that has a specific amount of endorsments.

Both database queries will use the same parameters. ie. The same persons, job, name to search for et cetera.

### Results
Note: This is after indexing. See [Executions](#executions)

Query | Average SQL | Median SQL | Average Neo4j | Median Neo4j
-----:|:-------:|:---------|:-------:|:---------
Depth 1 | 655ms | 656ms | 7.8ms | 4.8ms
Depth 2 | 1.5s | 1.5s | 14.5ms | 9ms 
Depth 3 | 2.5s | 2.5s | 185.4ms | 139.6ms 
Depth 4 | 5.6s | 5.4s | 3.6s | 2.8s 
Depth 5 | 10s | 10s | 45.4s | 30s 
Job | 637 µs | 495 µs | 108ms | 15ms
Name | 253 µs | 190 µs | 19.6ms | 12.2ms
Birthday | 407 µs | 391 µs | 23ms | 7.6ms
Endorsing count | 9.2s | 9.4s | 17.9s | 19.9s
Endorsment count | 11.6s | 12s | 19.2s | 19.7s

### Conclusion
as the hypothesis states, graphbased databases are clearly faster at searching data from relationships than sql is. Allthough the results seem to favor sql when it comes to deeper levels of relationships, it is thought to be that neo4j just runs out of ram or something along those lines. However taking all the experiements into consideration, it is clear to say that graphbased databases is faster at searching for data by relationships. In addition SQL is faster at searching for data by flat values on properties, with or without indexes. Allthough Neo4j becomes fast to a degree where you would say it's fine or adequate. But the fast SQL might be neccesary in very time sensitive applications.

These experiments are influenced by the choice of language. The given language go is a young language and hence does not have an official neo4j driver, this will affect the results. Further investigation across different languages can be done, to better answer the question at hand. However these results does still indicate largely that graphbased databases is a good choice when it comes to searching for data by nested relationships. Where as SQL based databases is a good choice when it comes to searching for data by flat values. 

### Experiment Execution

##### Specs
Localmachine has the following specs:
Relevant specs: 
 - Processor: Intel(R) Core(TM) i7-4720HQ CPU @ 2.60GHz 2.60 GHz
 - RAM: 16 GB
 
Virtual machine: [vagrantfile](https://github.com/DanielHauge/DBEX9/blob/master/vagrantfile) 

##### Reproductions instructions.
To setup the same experiments. Do the following: (It is assumed linux is used) Vagrantfile can be used to create virtual machine. (ubuntu 16 distribution)

- Step 1. Install docker
```
wget -O - http://get.docker.com | bash
```
- step 2. Setup SQL database (This script will setup docker container and also import relevant data)
```
wget -O - https://raw.githubusercontent.com/DanielHauge/DBEX9/master/RunSQLOnly.sh | bash
```
- Step 3. Setup Neo4J database ((This script will setup docker container)
```
wget -O - https://raw.githubusercontent.com/DanielHauge/DBEX9/master/RunNeo4JWithDatas.sh | bash
```
- Step 4. Import data to Neo4j
```
./importGraph.sh
```
- Step 5. Run the benchmarking application
```
sudo docker run -it --link psql --link neo4j danielhauge/dbex9
```
---------------
To setup, experiments where both databases are not running database container can be stopped and removed and this image tag of the program container can be run
```
docker run -it --link psql danielhauge/dbex9:sql
```
```
docker run -it --link neo4j danielhauge/dbex9:neo4j
```

#### Executions
To experiment, i've developed a simple benchmarking program in golang. With functions to integrate with both neo4j and sql with similar inputs and outputs. Note, that results will also affect which given language is used. It should be noted that Neo4j does not have a official driver for golang, therefor the results does not apply 100% to all circumstances in integration of every language. The code for the golang application can be found in this repository.

##### 1: Initial Experiments
note: In this experiment, the go program has been run on the local machine up against the databases which is on the virtual machine, both databases is running in different docker containers on the same virtual machine.

I ran the first experiments. Not all experiments completed because neo4j chrashed because of memory issues. My guess is that it was including everything and not having only distinct users. Also i suspect that Neo4j uses absurdly huge amounts of ram. So i will try to add more ram to the next execution. 

See [IntitialTest output gist](https://gist.github.com/DanielHauge/6d8d007ecfe3d76e26898126225589ab)

Query | Average SQL | Median SQL | Average Neo4j | Median Neo4j
-----:|:-------:|:---------:|:-------:|:---------
Depth 1 | 777ms | 772ms | 125ms | 24ms |
Depth 2 | 1.8s | 1.8s | 261ms | 156ms|
Depth 3 | 3s | 3s | 518ms | 210ms |
Depth 4 | 6.5s | 6.2s | 4.2s | 2.8s |
Depth 5 | 12.5s | 11.9s | N/A | N/A
Job | 58ms | 57ms | N/A | N/A
Name | 53ms | 54ms | N/A | N/A
Birthday | 64ms | 63ms | N/A | N/A
Endorsing count | 9.8s | 9.8s | N/A | N/A
Endorsment count | 13.5s | 13.5s | N/A | N/A

Disregarding missing data. We can stil see very promising results when it comes to finding nested relations. As hypothesized, the sql database seems slower when it comes to finding the relations by relationships. Neo4j is many times faster for 1,2 and 3 depth, but is starting to halt on the 4th, and crashed on the 5h level. I suspect that the level 4 query on neo4j is alot faster, but because of ram and memory shortage it is starting to halt, but also because it wasn't distincting persons.

#### 2: Neo4j query upgade + Ram upgrade
I shifted to a stronger machine with more ram. Since my computer only has 8gb ram. This also means the processor will get slower, so the numbers might look different. but the propotions should stay intact.

From previusly, it looks like it is true that neo4j uses lots and lots of ram.

See [Secondtest output gist](https://gist.github.com/DanielHauge/f7ad843a53fce4f0e5126cc8db77d521)

Query | Average SQL | Median SQL | Average Neo4j | Median Neo4j
-----:|:-------:|:---------:|:-------:|:---------
Depth 1 | 2.2s | 2.2s | 64ms | 14ms
Depth 2 | 3s | 2.5s | 101ms | 54ms 
Depth 3 | 4.5s | 4.2s | 288ms | 223ms 
Depth 4 | 14.8s | 15.4s | 7.1s | 5.7s 
Depth 5 | 14.8s | 11.6s | 45.2s | 35.5s 
Job | 46ms | 46ms | 738ms | 475ms
Name | 44ms | 44ms | 415ms | 380ms
Birthday | 52ms | 52ms | 456ms | 442ms
Endorsing count | 9s | 9s | 18s | 18s
Endorsment count | 11.6s | 12.1s | 20s | 20.6s

These results is a little weird in my opinion. The change is that i shifted to a computer with more ram, and gave the virtualmachine more ram, and also ran the go program in a docker container. The weird thing here is, neo4j got faster except level 4 but sql got slower. It is theorized that neo4j is using lots and lots of ram, and got faster by hogging all the ram. To best simulate a real integration and get better results, another experiment could be where only 1 database is on at a time.

#### 3. Only 1 DBMS at a time.

See [SQL-Only gist](https://gist.github.com/DanielHauge/16edb5a3175e49c3eb50e47fa39fb93b)
See [Neo4j-Only gist](https://gist.github.com/DanielHauge/0a368b303760ac2fdfda9fd28f244354)

Query | Average SQL | Median SQL | Average Neo4j | Median Neo4j
-----:|:-------:|:---------:|:-------:|:---------
Depth 1 | 651ms | 650ms | 9.2ms | 6.5ms
Depth 2 | 1.5s | 1.5s | 2.6s | 1.3s
Depth 3 | 2.5s | 2.5s | 6.7s | 447ms
Depth 4 | 5.6s | 5.2s | 3.3s | 2.3s
Depth 5 | 11s | 10.7s | 39s | 28s
Job | 73ms | 47ms | 546ms | 397ms
Name | 45ms | 45ms | 265ms | 249ms
Birthday | 52ms | 52ms | 320ms | 297ms
Endorsing count | 9.3s | 9.3s | 16s | 16s
Endorsment count | 11.8s | 12.2s | 17s | 17s

These results is representing a better picture i think. Allthough the results are still very inconsistent. As seen in individual query runtimes, Neo4j depth 2 has alot of very fast queries where as it has a single outlier with 26 seconds instead of 50 ms. It is theorized that neo4j or it's go driver needs to warm up or has some delay sometimes. The usual results is as: average is alot higher than median, which means it has outliers. Usualy in the start.

#### 4. Without restarting neo4j

See [SQL-Only2 gist](https://gist.github.com/DanielHauge/65df066d05ded8c90070e0ffb063ce14)
See [Neo4j-Only2 gist](https://gist.github.com/DanielHauge/3bd63ee460ec291ab9eea40fe9b5519a)

Query | Average SQL | Median SQL | Average Neo4j | Median Neo4j
-----:|:-------:|:---------|:-------:|:---------
Depth 1 | 668ms | 649ms | 8.8ms | 3.1ms 
Depth 2 | 1.5s | 1.5s | 38ms | 10ms 
Depth 3 | 2.5s | 2.5s | 207ms | 97ms 
Depth 4 | 5.6s | 5.5s | 3.3s | 2.3s 
Depth 5 | 10.5s | 10.3s | 39.8s | 27.1s
Job | 48ms | 47ms | 357ms | 324ms 
Name | 45ms | 44.9ms | 294ms | 250ms 
Birthday | 53ms | 53ms | 313ms | 298ms 
Endorsing count | 9s | 9s | 15.3s | 15.8s 
Endorsment count | 11.7s | 12.2s | 16s | 16s

as can be seen, neo4j is very fast the second time around. Allthough it is not clearly the reason. It might also be because of the go driver, that the there is some inconsistencies.





