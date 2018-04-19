docker exec psql bash -c "sed -i 1d /root/social_network_nodes.csv"
docker exec psql bash -c "sed -i 1d /root/social_network_edges.csv"
echo 'header lines removes'
docker exec psql bash -c "psql -U postgres -c 'CREATE TABLE vertex(id INTEGER PRIMARY KEY, name VARCHAR, job VARCHAR, birthday VARCHAR);'"
docker exec psql bash -c "psql -U postgres -c 'CREATE TABLE edges(source INTEGER references vertex(id), dest INTEGER references vertex(id));'"
echo 'tables created'
docker exec psql bash -c "psql -U postgres -c '\copy vertex FROM /root/social_network_nodes.csv DELIMITER ',' CSV'"
docker exec psql bash -c "psql -U postgres -c '\copy edges FROM /root/social_network_edges.csv DELIMITER ',' CSV'"
echo 'data imported'