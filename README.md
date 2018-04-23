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
Note: These results are after indexing of both databases. See Execution sub-results for the unindexed results. (It will affect all except depth searches.)

#### Graph database
Query | Average | Median
-----:|:-------:|:---------
Depth 1 |
Depth 2 |
Depth 3 |
Depth 4 |
Depth 5 |
Job |
Name |
Birthday |
Endorsing count |
Endorsment count |

#### SQL database
Query | Average | Median
-----:|:-------:|:---------
Depth 1 |
Depth 2 |
Depth 3 |
Depth 4 |
Depth 5 |
Job |
Name |
Birthday |
Endorsing count |
Endorsment count |

### Conclusion

### Execution
#### Enviroment
Vagrant VM and docker

##### Specs
Localmachine has the following specs:
Relevant specs: 
 - Processor: Intel(R) Core(TM) i5-3570k CPU @ 3.40GHz 3.4 GHz
 - RAM: 8 GB
 - OS: Windows 10 Home edition (With all available updates)
 
Virtual machine: [vagrantfile](#DBEX9/vagrantfile) 

##### Reproductions instructions.
(How to simulate the same results)

#### Executions
(What i did to run the tests)

##### 1: Initial Experiments
note: In this experiment, the go program has been run on the local machine up against the databases which is on the virtual machine.

I ran the first tests. Not all tests completed because neo4j chrashed because of memory issues. My guess is that it was including everything and not having only distinct users. Also i suspect that Neo4j uses absurdly huge amounts of ram. So i will try to add more ram to the next execution. 

See [IntitialTest output gist](https://gist.github.com/DanielHauge/6d8d007ecfe3d76e26898126225589ab)

Query | Average SQL | Median SQL | Average Neo4j | Median Neo4j
-----:|:-------:|:---------|:-------:|:---------
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

#### Result set 2
(what second tests show)






