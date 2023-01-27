# Launching the service in Kubernetes

make push

minikube start

minikube addons enable ingress

kubectl config use-context minikube

kubectl get namespaces

kubectl get pods --all-namespaces

kubectl get logs -f <pod_name>

kubectl logs -f <pod_name> -n namespace

kubectl get pods -o wide

kubectl get ingress

echo "$(minikube ip) krist.test" | sudo tee -a /etc/hosts

curl -i http://krist.test/home