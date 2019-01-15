# docker setup es 9200映射到宿主机 9300没有映射 只用9200即可
systemctl start docker

# docker run -d -p 9200:9200 elasticsearch
docker run -d -v /data/docker/elasticsearch/etc:/etc/elasticsearch -v /data/docker/elasticsearch/data:/usr/share/elasticsearch/data -p 9200:9200  elasticsearch:5.6
docker ps | grep elasticsearch

# docker run -d -e ELASTICSEARCH_URL=http://x.x.x.x:9200 -p 9000:5601 kibana:5.6
