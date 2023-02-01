# Launching the service in Kubernetes

git add .

git commit -m "tmp.yaml change"

git push

go run main.go

curl -i http://127.0.0.1:8080/home

dep init

docker build -t kubserv:0.0.5 .

make build

make container

make push

make run

minikube start

kubectl apply -f tmp.yaml

kubectl get pods

kubectl get deployment

kubectl get service

kubectl get ingress

traceroute -T kubserv.test -p 80

kubectl describe ingress minimal-ingress

kubectl get pods --all-namespaces

kubectl describe pod nginx-ingress-controller-6fc5bcc --namespace kube-system | grep Ports

kubectl port-forward nginx-ingress-controller-6fc5bcc 8080:80 --namespace kube-system

echo "$(minikube ip) kubserv.test" | sudo tee -a /etc/hosts

curl -i http://kubserv.test/home

minikube stop

minikube delete

yes| docker image prune



curl: (7) failed to connect to kubserv.test port 80: no route to host