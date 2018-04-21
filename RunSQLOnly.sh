wget https://github.com/datsoftlyngby/soft2018spring-databases-teaching-material/raw/master/data/archive_graph.tar.gz
tar -xvzf archive_graph.tar.gz
docker run -p 5432:5432 --rm -d -v $(pwd)/:/root/ --name psql postgres:alpine
echo 'Docker running'
sleep 1s
wget -O - https://raw.githubusercontent.com/DanielHauge/DBEX9/master/ImportSQL.sh | bash

