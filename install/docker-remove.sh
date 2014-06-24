#/!bin/bash

# Remove all stopped containers.
docker rm $(docker ps -a -q)

# Remove all untagged images
# docker rmi $(docker images | grep "^<none>" | awk "{print $3}")


docker stop mongo          
docker stop postgres       
docker stop rabbitmq       
docker stop redis          
docker stop kontrol        
docker stop proxy          
docker stop kloud          
docker stop rerouting      
docker stop webserver      
docker stop sourceMapServer
docker stop authWorker     
docker stop social         
docker stop guestCleaner   
docker stop cronJobs       
docker stop broker         
docker stop emailSender   

docker rm mongo          
docker rm postgres       
docker rm rabbitmq       
docker rm redis          
docker rm kontrol        
docker rm proxy          
docker rm kloud          
docker rm rerouting      
docker rm webserver      
docker rm sourceMapServer
docker rm authWorker     
docker rm social         
docker rm guestCleaner   
docker rm cronJobs       
docker rm broker         
docker rm emailSender     