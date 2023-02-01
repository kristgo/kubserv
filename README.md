# Launching the service in Kubernetes

dep init

make push

minikube start

kubectl apply -f tmp.yaml

minikube addons enable metrics-server

minikube addons enable ingress

kubectl config use-context minikube

make minikube

kubectl get deployment

kubectl get service

kubectl get ingress

echo "$(minikube ip) kubserv.test" | sudo tee -a /etc/hosts

curl -i http://kubserv.test/home

docker image prune