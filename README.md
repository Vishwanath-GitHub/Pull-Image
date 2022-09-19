# PULL IMAGE CODE

Steps to run this code:

1. kubectl label node <node-name> value=pullimage
2. kubectl create configmap sample-configmap --from-literal=image="nginx,redis"
3. kubectl apply -f deploy.yaml
4. SSH into node
5. docker images