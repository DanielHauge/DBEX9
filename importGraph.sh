sudo docker exec neo4j sh -c 'neo4j stop'
sudo docker exec neo4j sh -c 'rm -rf /var/lib/neo4j/data/databases/graph.db'
sudo docker exec neo4j sh -c 'rm -rf data/databases/graph.db'
sudo docker exec neo4j sh -c 'neo4j-admin import --nodes:Person import/social_network_nodes.csv --relationships:ENDORSES import/social_network_edges.csv --ignore-missing-nodes=true --ignore-duplicate-nodes=true --id-type=INTEGER'
sudo docker restart neo4j