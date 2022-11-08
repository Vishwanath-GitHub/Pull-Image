# PULL IMAGE CODE

Steps to run this code:

1. kubectl label node <node-name> value=pullimage
2. kubectl create configmap sample-configmap --from-literal=image="nginx,redis"
3. kubectl apply -f deploy.yaml
4. SSH into node
5. docker images


Helpful links:
* https://stackoverflow.com/questions/58622015/using-docker-socket-in-kubernetes-pod
* https://stackoverflow.com/questions/45805563/pull-a-file-from-a-docker-image-in-golang-to-local-file-system
* https://stackoverflow.com/questions/38125212/how-to-connect-to-a-remote-socket-in-docker-engine-api
* https://kubernetes.io/docs/tasks/configure-pod-container/assign-pods-nodes/
