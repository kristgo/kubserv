# Launching the service in Kubernetes

make push

minikube start

minikube addons enable ingress

kubectl config use-context minikube

kubectl get deployment

kubectl get service

kubectl get ingress

echo "$(minikube ip) krist.test" | sudo tee -a /etc/hosts

curl -i http://krist.test/home