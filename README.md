# Launching the service in Kubernetes

git add .

git commit -m "tmp.yaml change"

git push

go run main.go

curl -i http://127.0.0.1:8080/home

dep init

docker build -t kubserv .

make build

make container

make push

docker container run -it --name kubserv --rm -p 8080:8080 kubserv

curl -i http://127.0.0.1:8080/home

minikube start

minikube addons enable ingress

kubectl apply -f tmp.yaml

kubectl get pods

kubectl get deployment

kubectl get service

kubectl get ingress

sudo traceroute -T kubserv.test -p 80

kubectl describe ingress minimal-ingress

kubectl get pods --all-namespaces

kubectl -n ingress-nginx logs ingress-nginx-controller-77669ff58-grhzh

kubectl describe pod ingress-nginx-controller-77669ff58-ck8ws --namespace ingress-nginx | grep Ports

kubectl port-forward ingress-nginx-controller-77669ff58-ck8ws 8080:80 --namespace ingress-nginx

http://localhost:8080

echo "$(minikube ip) kubserv.test" | sudo tee -a /etc/hosts

curl -i http://kubserv.test/home

minikube stop

minikube delete

yes| docker image prune

