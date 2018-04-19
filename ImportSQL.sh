docker exec psql sh -c "sed -i 1d /root/social_network_nodes.csv"
docker exec psql sh -c "sed -i 1d /root/social_network_edges.csv"
echo 'header lines removes'
docker exec psql sh -c "psql -U postgres -c 'CREATE TABLE vertex(id INTEGER PRIMARY KEY, name VARCHAR, job VARCHAR, birthday VARCHAR);'"
docker exec psql sh -c "psql -U postgres -c 'CREATE TABLE edges(source INTEGER references vertex(id), dest INTEGER references vertex(id));'"
echo 'tables created. IMPORTING IS STARTING!'
docker exec psql sh -c "psql -U postgres -c '\copy vertex FROM /root/social_network_nodes.csv CSV'"
echo 'Depending on your specs, this step varies in time. Can take alot of time, be patient. 11 milion * 2 edges ids that needs to be hashed :D'
docker exec psql sh -c "psql -U postgres -c '\copy edges FROM /root/social_network_edges.csv CSV'"
echo 'data imported'