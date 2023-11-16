# Learn Kustomize - with ArgoCD

## Motivation behind it:
* Templating systems are extremely difficult to implement and maintain, e.g. Helm. You can't template everything. If upstream changes something, you will need to create a new copy/fork of the HElm chart and adjust your changes again, etc.
* Solving this through programming (Tanka/Jsonnet) is even more difficult, because everything must be programmed in a specific programming language (Jsonnet). It is painful to learn yet another programming language, painful to code, almost impossible to understand, and painful to debug. It does not follow any standard, and depends on how the programmer designed it / wrote it. 
* Eventually the above generates YAML. So why not have YAML files in the first place, and adjust them when and where needed using some patching system. Afterall, we are using patching anyway in our code when we use any version control system, such as `git`. So why this should be any different.
* Also it's not that we have 4 KB of RAM and 360 KB of floppy disk to store our stuff. We have pretty powerful computers with lots of memory and lots of storage. Storing and processing long YAML files should really not be a problem. 
* We are favoring simplicity and uniformity by using YAML which is the most important win here. Even if there is somewhat repetition of code, it is much better than completely obscure and incomprehensible code.


## How the apps would be deployed:

* Application git repo with actual source code , which generates Docker image using some CD steps
* Application configuration (gitops) repo, which contains Kubernetes YAML manifest files for various environments (dev, demo, prod) (under separate directory paths), with different configuration for each environment (as required)
* `ArgoCD-applications-dev` repo with application definitions to watch the `/dev` path in the above mentioned gitops repo
* `ArgoCD-applications-demo` repo with application definitions to watch the `/demo` path in the above mentioned gitops repo
* `ArgoCD-applications-prod` repo with application definitions to watch the `/prod` path in the above mentioned gitops repo


We also need:
* A kubernetes cluster - to host/run these applications
* ArgoCD - running inside the kubernetes cluster


# (Basic) Structure of repositories:

```

(Your Git provider - GitHub/GitLab/etc)

|
|
|------(repo) patients
|               |
|               |-src
|               |   |-main.go 
|               |
|               |-Dockerfile   
|
|
|-------(repo) patients-gitops
|               |
|               |-base
|               |  |-deployment.yaml
|               |  |-kustomize.yaml
|               |
|               |-overlays
|                  |-dev
|                  |  |-replicas.yaml
|                  |  |-kustomize.yaml
|                  |
|                  |-demo
|                  |  |-replicas.yaml
|                  |  |-kustomize.yaml
|                  |
|                  |-prod
|                     |-replicas.yaml
|                     |-kustomize.yaml
|
|
|-------(repo) argocd-apps-dev
|               |
|               |-patients-dev.yaml
|               |-doctors-dev.yaml
|
|-------(repo) argocd-apps-demo
|               |
|               |-patients-demo.yaml
|               |-doctors-demo.yaml
|
|-------(repo) argocd-apps-prod
|               |
|               |-patients-prod.yaml
|               |-doctors-prod.yaml
|
```

This repository contains all of the above mentioned directories to help understand this example. In real-world situation, the above will all be separate git repositories.




------

# Resources:

## Howtos:
* [Kustomize official website](https://kustomize.io/)

## Videos:
* [Simplify configuration management](https://youtu.be/Twtbg6LFnAg?si=flyaac2RyyHuXkjN)
* [Organize YAML mess with Kustomize](https://youtu.be/1fCAwFGX38U)
* [Declarative configuration for Kubernetes](https://youtu.be/WWJDbHo-OeY)
* [Stop forking Helm charts](https://youtu.be/pRG47EQ5OAg)
* [Deploy your apps with template-free YAML](https://youtu.be/ahMIBxufNR0)
* [TGI Kubernetes with Joe Beda](https://youtu.be/NFnpUlt0IuM)
* [Template-Free Configuration Customization for Kubernetes - Jeffrey Regan/Google](https://youtu.be/EZ7kxa2GKYQ?si=jbXAikcZsVGsQrsu)
* [kustomize our manifests with style](https://youtu.be/KvXcc7lXiXc?si=8AxXyEDT64512vu9)


