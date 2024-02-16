#!/bin/bash

#################################################################################################
## create Network
echo "-----------------[[ The system is installing Network Container ]]----------------

"
docker network create \
  --driver=bridge \
  --subnet=10.5.0.0/27 \
  --ip-range=10.5.0.0/27 \
  --gateway=10.5.0.1 \
  network_service

echo "

-----------------------------------------------------------------------


"
#################################################################################################
## create file webservice api

echo "-----------------[[ The system is Create file webservice API ]]----------------

"
cd ..
cd webservice_api 
# go build web
# echo "file : webservice_api is up-to-date"
echo "

-----------------------------------------------------------------------

"

#################################################################################################
## install docker container

echo "--------------------[[ The system is installing Container ]]------------------------

"
cd ..

echo "--------------------------------------"

docker-compose -f docker-compose.database.yml -p "database" up -d

echo "--------------------------------------"

docker-compose -f docker-compose.zabbix.yml -p "zabbix" up -d

echo "--------------------------------------"

docker-compose -f docker-compose.grafana.yml -p "grafana" up -d


echo "--------------------------------------"

docker-compose -f docker-compose.portainer.yml -p "portainer" up -d

echo "--------------------------------------"

docker-compose -f docker-compose.logrotate.yml -p "logrotate" up -d

echo "--------------------------------------"

docker-compose -f docker-compose.webservice_api.yml -p "webservice_api" up -d

echo "--------------------------------------"

docker-compose -f docker-compose.nginx.yml -p "nginx" up -d

echo "

-----------------------------------------------------------------------

"  
##################################################################################################





