# Use/setup a Kubernetes cluster:

If you have access to a Kubernetes cluster that you can do experiments with, then simply connect to it (or make sure that you can connect to it).

If you don't have one, you have several options:
* Install a small/test Kubrnetes cluser at one of the cloud providers. Such as GKE, AKS, DigitialOcean, etc.
* Install a local single-node Kubernetes cluster on your local computer using MiniKube, Kind, K3d, etc.

## Local Kubernetes cluster - MiniKube:

Install minikube.

Reference: [https://minikube.sigs.k8s.io/docs/start/](https://minikube.sigs.k8s.io/docs/start/)


```
[root@kworkhorse ~]# curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-latest.x86_64.rpm
sudo rpm -Uvh minikube-latest.x86_64.rpm
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 19.3M  100 19.3M    0     0  6671k      0  0:00:02  0:00:02 --:--:-- 6671k
Verifying...                          ################################# [100%]
Preparing...                          ################################# [100%]
Updating / installing...
   1:minikube-1.32.0-0                ################################# [100%]
[root@kworkhorse ~]# 

```

Start minikube:

```
[kamran@kworkhorse ~]$ minikube start
😄  minikube v1.32.0 on Fedora 38
    ▪ KUBECONFIG=/home/kamran/.kube/config:/home/kamran/.kube/kubeadm-cluster.conf
✨  Automatically selected the docker driver. Other choices: kvm2, qemu2, ssh
📌  Using Docker driver with root privileges
👍  Starting control plane node minikube in cluster minikube
🚜  Pulling base image ...
💾  Downloading Kubernetes v1.28.3 preload ...
    > preloaded-images-k8s-v18-v1...:  403.35 MiB / 403.35 MiB  100.00% 5.31 Mi
    > gcr.io/k8s-minikube/kicbase...:  453.90 MiB / 453.90 MiB  100.00% 4.62 Mi
🔥  Creating docker container (CPUs=2, Memory=3900MB) ...
🐳  Preparing Kubernetes v1.28.3 on Docker 24.0.7 ...
    ▪ Generating certificates and keys ...
    ▪ Booting up control plane ...
    ▪ Configuring RBAC rules ...
🔗  Configuring bridge CNI (Container Networking Interface) ...
    ▪ Using image gcr.io/k8s-minikube/storage-provisioner:v5
🔎  Verifying Kubernetes components...
🌟  Enabled addons: storage-provisioner, default-storageclass
🏄  Done! kubectl is now configured to use "minikube" cluster and "default" namespace by default
[kamran@kworkhorse ~]$ 
```

Check that it works:
```
[kamran@kworkhorse ~]$ kubectl  get nodes
NAME       STATUS   ROLES           AGE   VERSION
minikube   Ready    control-plane   92s   v1.28.3
[kamran@kworkhorse ~]$ 
```
