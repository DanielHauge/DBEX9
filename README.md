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
- All persons that has a speicifc name
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
Id |
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
Id |
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


##### Reproductions instructions.
(How to simulate the same results)

#### Executions
(What i did to run the tests)

#### Result set 1
(What initial tests show)

#### Result set 2
(what second tests show)






