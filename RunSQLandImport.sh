wget https://github.com/datsoftlyngby/soft2018spring-databases-teaching-material/raw/master/data/archive_graph.tar.gz
tar -xvzf archive_graph.tar.gz
docker run -p 5432:5432 --rm -d -v $(pwd)/:/root/ --name psql postgres:alpine
docker exec -i psql bash -c "sed -i 1d /root/social_network_nodes.csv"
docker exec -i psql bash -c "sed -i 1d /root/social_network_edges.csv"
docker exec -i psql bash -c "psql -U postgres -c 'CREATE TABLE vertex(id INTEGER PRIMARY KEY, name VARCHAR, job VARCHAR, birthday VARCHAR);'"
docker exec -i psql bash -c "psql -U postgres -c 'CREATE TABLE edges(source INTEGER references vertex(id), dest INTEGER references vertex(id));'"
docker exec -i psql bash -c "psql -U postgres -c '\copy vertex FROM /root/social_network_nodes.csv DELIMITER ',' CSV'"
docker exec -i psql bash -c "psql -U postgres -c '\copy edges FROM /root/social_network_edges.csv DELIMITER ',' CSV'"

