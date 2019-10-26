# Prometheus
Repo to learn and play with Prometheus

### Prerequisities:
Install minikube with docker (ubuntu machine):

	apt update && apt -y install docker.io
	curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl
	chmod +x kubectl && cp kubectl /usr/local/bin
	kubectl version
	curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
	chmod +x minikube
	cp minikube /usr/local/bin        
    
Run minikube with command:

	minikube start --kubernetes-version=v1.16.0 --memory=6g --bootstrapper=kubeadm --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook --extra-config=scheduler.address=0.0.0.0 --extra-config=controller-manager.address=0.0.0.0 --vm-driver=none
	minikube status
	
Verify installation:

	kubectl get po -n kube-system

Should look like this:

	NAME                               READY   STATUS    RESTARTS   AGE
	coredns-5644d7b6d9-2lvff           1/1     Running   0          94s
	coredns-5644d7b6d9-b9xkp           1/1     Running   0          94s
	etcd-minikube                      1/1     Running   0          42s
	kube-addon-manager-minikube        1/1     Running   0          42s
	kube-apiserver-minikube            1/1     Running   0          40s
	kube-controller-manager-minikube   1/1     Running   0          47s
	kube-proxy-h85pc                   1/1     Running   0          94s
	kube-scheduler-minikube            1/1     Running   0          30s
	storage-provisioner                1/1     Running   0          93s

## How to run demo:
### 1. create prometheus config:
	kubectl create configmap prometheus-example-cm --from-file prometheus-config.yml

### 2. create prometheus server with created configuration:
	kubectl create -f prometheus-deployment.yaml	

This should start prometheus server and expose prometheus ports

### 3. create test deployment which expose metrics
	kubectl create -f deploy-app.yaml
