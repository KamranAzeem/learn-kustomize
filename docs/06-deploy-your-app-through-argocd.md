# Deploy your app through Argo CD:

First, delete the deployment (and all other objects) that were created during the manual application deployment in the previous steps.


```
[kamran@kworkhorse patients-gitops]$ kubectl -n dev delete -k kustomize/overlays/dev/
configmap "patients-khtkg49662" deleted
service "patients" deleted
deployment.apps "patients" deleted
[kamran@kworkhorse patients-gitops]$ 
```


## Create a YAML deployment file for your application required by ArgoCD:

This file will tell ArgoCD what to deploy and where and what to watch for, etc.

It looks like this:

```
[kamran@kworkhorse argocd-apps-dev]$ pwd
/home/kamran/Projects/Personal/github/argocd-apps-dev


[kamran@kworkhorse argocd-apps-dev]$ cat patients-dev.yaml

# This file deploys the application and also watches:
#   a) the specified path in the related "-gitops" directory
#      for configuration changes e.g. "dev", or "prod"
#   b) the container registry for changes in the image tag
# ----------------------------------------------------------

apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  #name: image-updater-patients
  name: patients-dev
  namespace: argocd
  #~ finalizers:
    #~ - resources-finalizer.argocd.argoproj.io
  annotations:
    argocd-image-updater.argoproj.io/image-list: patients=docker.io/kamranazeem/patients
    argocd-image-updater.argoproj.io/patients.update-strategy: latest
    argocd-image-updater.argoproj.io/patients.allow-tags: regexp:^[0-9a-f]{7}$
    #  the regular expression matches a 7-digit hexadecimal string 
    #    that could represent the short version of a Git commit SHA, 
    #    so it would match tags like a5fb3d3 or f7bb2e3, but not latest or master.
    argocd-image-updater.argoproj.io/patients.force-update: "true"
    argocd-image-updater.argoproj.io/write-back-method: git:secret:argocd/github-credentials
    argocd-image-updater.argoproj.io/branch: master
spec:
  project: default
  source:
    repoURL: https://github.com/KamranAzeem/patients-gitops.git
    targetRevision: HEAD
    path: kustomize/overlays/dev 
  destination:
    server: https://kubernetes.default.svc
    namespace: dev
  syncPolicy:
    automated:
      selfHeal: true
      prune: true
      allowEmpty: true

```


```
[kamran@kworkhorse argocd-apps-dev]$ kubectl apply -f patients-dev.yaml 
application.argoproj.io/patients created
[kamran@kworkhorse argocd-apps-dev]$ 
```


This should immediately show something happening in the Argo CD web UI.

(screenshot here)


Verify:

```
[kamran@kworkhorse argocd-apps-dev]$ kubectl  -n dev get pods
NAME                        READY   STATUS    RESTARTS   AGE
patients-757ddb786b-fg5t4   1/1     Running   0          44s
[kamran@kworkhorse argocd-apps-dev]$ 
```


## Change code in main application repository to force a new build of container image:
This will cause a new container image to be built , and ideally ArgoCD should detect this, and will update the "patients-gitops" directory.


To help understand the example, notice the application code says our office is in Oslo.

```
[kamran@kworkhorse patients-gitops]$ kubectl -n dev port-forward svc/patients 8000:80 
Forwarding from 127.0.0.1:8000 -> 8080
Forwarding from [::1]:8000 -> 8080
Handling connection for 8000
```


(in a separate terminal)
```
[kamran@kworkhorse argocd-apps-dev]$ curl localhost:8000
Hello Patients! We are located in: Oslo
```

Now, change application code, (change name of city), and commit / push your changes.


```
[kamran@kworkhorse patients]$ pwd
/home/kamran/Projects/Personal/github/patients
[kamran@kworkhorse patients]$ 
```


```
[kamran@kworkhorse patients]$ grep located src/main.go 
    fmt.Fprintf(w, os.Getenv("GREETING") + " Patients! We are located in: Oslo")


[kamran@kworkhorse patients]$ sed -i 's/Oslo/Asker/g' src/main.go 


[kamran@kworkhorse patients]$ grep located src/main.go 
    fmt.Fprintf(w, os.Getenv("GREETING") + " Patients! We are located in: Asker")
[kamran@kworkhorse patients]$ 
```



Push changes to github:

```
[kamran@kworkhorse patients]$ git add src/main.go 


[kamran@kworkhorse patients]$ git commit -m "Updated name of city to Asker"
[master 5c16893] Updated name of city to Asker
 1 file changed, 1 insertion(+), 1 deletion(-)


[kamran@kworkhorse patients]$ git push
Enumerating objects: 7, done.
Counting objects: 100% (7/7), done.
Delta compression using up to 4 threads
Compressing objects: 100% (3/3), done.
Writing objects: 100% (4/4), 352 bytes | 352.00 KiB/s, done.
Total 4 (delta 2), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (2/2), completed with 2 local objects.
To github.com:KamranAzeem/patients.git
   8d04857..5c16893  master -> master
[kamran@kworkhorse patients]$ 
```


This will kickoff the CI/CD pipeline. Check GitHub Actions in the 'patients' repository, to verify the pipeline ran correctly without any errors.

Use docker hub's web UI to verify that a new image was created.

(Screenshot here)


## Access the actual application:

(port-forward again if necessary)

```
[kamran@kworkhorse patients-gitops]$ kubectl -n dev port-forward svc/patients 8000:80 
Forwarding from 127.0.0.1:8000 -> 8080
Forwarding from [::1]:8000 -> 8080
Handling connection for 8000
```


Check application using curl:

```
[kamran@kworkhorse argocd-apps-dev]$ curl localhost:8000
Hello Patients! We are located in: Asker
```
(success!!)

Check the Argo CD interface, and it should show the application has synced, and a new image was detected.

(screenshot here)



Last check would be to pull latest changes of the `patients-gitops` repository. Then examine the file that argocd created under the `kustomize/overlays/dev/` path.


```
[kamran@kworkhorse patients-gitops]$ git pull
remote: Enumerating objects: 10, done.
remote: Counting objects: 100% (10/10), done.
remote: Compressing objects: 100% (6/6), done.
remote: Total 6 (delta 0), reused 6 (delta 0), pack-reused 0
Unpacking objects: 100% (6/6), 737 bytes | 368.00 KiB/s, done.
From github.com:KamranAzeem/patients-gitops
   14f78a5..3dcf5bf  master     -> origin/master
Updating 14f78a5..3dcf5bf
Fast-forward
 kustomize/overlays/dev/.argocd-source-patients.yaml | 3 +++
 1 file changed, 3 insertions(+)
 create mode 100644 kustomize/overlays/dev/.argocd-source-patients.yaml
```

```
[kamran@kworkhorse patients-gitops]$ cat kustomize/overlays/dev/.argocd-source-patients.yaml
kustomize:
  images:
  - docker.io/kamranazeem/patients:5c16893
[kamran@kworkhorse patients-gitops]$ 
```


------
# Troubleshooting:

Check logs of the pods running inside ArgoCD namespace.

```
[kamran@kworkhorse patients]$ kubectl -n argocd logs -f argocd-image-updater-88454679d-zzjb9 

. . . 

time="2023-11-20T21:04:50Z" level=info msg="git add /tmp/git-patients1001139719/kustomize/overlays/dev/.argocd-source-patients.yaml" dir=/tmp/git-patients1001139719 execID=46ffb
time="2023-11-20T21:04:50Z" level=info msg=Trace args="[git add /tmp/git-patients1001139719/kustomize/overlays/dev/.argocd-source-patients.yaml]" dir=/tmp/git-patients1001139719 operation_name="exec git" time_ms=6.99281
time="2023-11-20T21:04:50Z" level=info msg="git commit -a -F /tmp/image-updater-commit-msg2538615557" dir=/tmp/git-patients1001139719 execID=49377
time="2023-11-20T21:04:50Z" level=info msg=Trace args="[git commit -a -F /tmp/image-updater-commit-msg2538615557]" dir=/tmp/git-patients1001139719 operation_name="exec git" time_ms=13.469263999999999
time="2023-11-20T21:04:50Z" level=info msg="git push origin master" dir=/tmp/git-patients1001139719 execID=d3ce4
time="2023-11-20T21:04:51Z" level=info msg=Trace args="[git push origin master]" dir=/tmp/git-patients1001139719 operation_name="exec git" time_ms=1238.800447
time="2023-11-20T21:04:51Z" level=info msg="Successfully updated the live application spec" application=patients
time="2023-11-20T21:04:51Z" level=info msg="Processing results: applications=1 images_considered=1 images_skipped=0 images_updated=1 errors=0"
```

