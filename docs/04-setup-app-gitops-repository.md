# Setup the app-gitops repository:

So the next logical step is to create a `patients-gitops` directory in Github, and populate it with the contents of local/corresponding directory `patients-gitops`. 

```
[kamran@kworkhorse patients-gitops]$ pwd
/home/kamran/Projects/Personal/github/patients-gitops

[kamran@kworkhorse patients-gitops]$ git commit -m "First commit - with Kustomize files"
[master (root-commit) 14f78a5] First commit - with Kustomize files
 14 files changed, 162 insertions(+)
 create mode 100644 README.md
 create mode 100644 kustomize/base/deployment.yaml
 create mode 100644 kustomize/base/kustomization.yaml
 create mode 100644 kustomize/base/service.yaml
 create mode 100644 kustomize/overlays/dev/kustomization.yaml
 create mode 100644 kustomize/overlays/dev/replicas.yaml
 create mode 100644 kustomize/overlays/dev/variables.configmap
 create mode 100644 kustomize/overlays/prod/kustomization.yaml
 create mode 100644 kustomize/overlays/prod/replicas.yaml
 create mode 100644 kustomize/overlays/prod/variables.configmap
 create mode 100644 plain-kubernetes/configmap.yaml
 create mode 100644 plain-kubernetes/deployment.yaml
 create mode 100644 plain-kubernetes/namespace.yaml
 create mode 100644 plain-kubernetes/service.yaml
[kamran@kworkhorse patients-gitops]$ 



[kamran@kworkhorse patients-gitops]$ git branch -M master

[kamran@kworkhorse patients-gitops]$  git remote add origin git@github.com:KamranAzeem/patients-gitops.git

[kamran@kworkhorse patients-gitops]$  git push -u origin master
Enumerating objects: 21, done.
Counting objects: 100% (21/21), done.
Delta compression using up to 4 threads
Compressing objects: 100% (19/19), done.
Writing objects: 100% (21/21), 2.56 KiB | 437.00 KiB/s, done.
Total 21 (delta 4), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (4/4), done.
To github.com:KamranAzeem/patients-gitops.git
 * [new branch]      master -> master
branch 'master' set up to track 'origin/master' by rebasing.
[kamran@kworkhorse patients-gitops]$ 
```


## Deploy the app manually:
AT this point, we an deploy the application directly/manually in the `dev` namespace, using the following commands:


First, create the namespaces:
```
[kamran@kworkhorse patients-gitops]$ kubectl  get namespaces
NAME              STATUS   AGE
argocd            Active   3d
default           Active   3d1h
kube-node-lease   Active   3d1h
kube-public       Active   3d1h
kube-system       Active   3d1h

[kamran@kworkhorse patients-gitops]$ kubectl create namespace dev
namespace/dev created

[kamran@kworkhorse patients-gitops]$ kubectl create namespace demo
namespace/demo created

[kamran@kworkhorse patients-gitops]$ kubectl create namespace prod
namespace/prod created
[kamran@kworkhorse patients-gitops]$ 
```


Deploy the app:

```
[kamran@kworkhorse patients-gitops]$ kubectl  -n dev apply -k kustomize/overlays/dev/
configmap/patients-khtkg49662 created
service/patients created
deployment.apps/patients created
[kamran@kworkhorse patients-gitops]$ 
```


```
[kamran@kworkhorse patients-gitops]$ kubectl  -n dev get pods,service,configmaps
NAME                            READY   STATUS    RESTARTS   AGE
pod/patients-757ddb786b-lw2ds   1/1     Running   0          51s

NAME               TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
service/patients   ClusterIP   10.106.16.152   <none>        80/TCP    51s

NAME                            DATA   AGE
configmap/kube-root-ca.crt      1      2m9s
configmap/patients-khtkg49662   2      51s
[kamran@kworkhorse patients-gitops]$ 
```

## Forward the port to local computer to test out:
```
[kamran@kworkhorse patients-gitops]$ kubectl -n dev port-forward svc/patients 8000:80 
Forwarding from 127.0.0.1:8000 -> 8080
Forwarding from [::1]:8000 -> 8080
Handling connection for 8000
```



```
[kamran@kworkhorse ~]$ curl localhost:8000
Hello Patients! We are located in: Oslo
```
