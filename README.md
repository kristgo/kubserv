# New cars in Russia in 2020 with Ingress in Kubernetes

docker build -t kubserv .

docker tag kubserv kristgo/kubserv:1.0

docker push kristgo/kubserv:1.0

minikube start

minikube addons enable ingress

kubectl apply -f deployment.yaml

kubectl patch svc ingress-nginx-controller -n ingress-nginx -p '{"spec": {"type": "LoadBalancer"}}'

kubectl patch service ingress-nginx-controller -n ingress-nginx --patch "$(cat ip-patch.yaml)"

kubectl apply -f ingress.yaml

kubectl get pods --all-namespaces

kubectl port-forward ingress-nginx-controller-77669ff58-jswx8 8080:80 --namespace ingress-nginx

http://localhost:8080

kubectl delete deployment kubserv

kubectl delete service kubserv

minikube stop

minikube delete
