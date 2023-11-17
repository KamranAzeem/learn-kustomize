# Use/setup a Kubernetes cluster:

If you have access to a Kubernetes cluster that you can do experiments with, then simply connect to it (or make sure that you can connect to it).

If you don't have one, you have several options:
* Install a small/test Kubrnetes cluser at one of the cloud providers. Such as GKE, AKS, DigitialOcean, etc.
* Install a local single-node Kubernetes cluster on your local computer using MiniKube, Kind, K3d, etc.

## Local Kubernetes cluster - MiniKube:

Install minikube.

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
