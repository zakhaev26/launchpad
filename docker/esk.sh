

sudo docker network create elastic
sudo docker pull docker.elastic.co/elasticsearch/elasticsearch:8.13.0
sudo docker run --name elasticsearch --net elastic -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -t docker.elastic.co/elasticsearch/elasticsearch:8.13.0

