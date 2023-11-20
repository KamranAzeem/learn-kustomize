# Prepare your app for deployment:

As mentioned earlier (in the main README file), your application will have:
* a main source-code git repository
* a configuration or "-gitops" git repository
* a docker/container image for your application with correct tags, such as `21b9c1d`, `latest`, `rc-1.0.1`, `v1.1.3`, `1.1.13` etc. This will be based on the hash or tag of the git commit in your main soruce-code git repository.


```
$ tree patients

patients
├── Dockerfile
├── README.md
└── src
    └── main.go
```

```
$ tree patients-gitops/

patients-gitops/
├── kustomize
│   ├── base
│   │   ├── deployment.yaml
│   │   ├── kustomization.yaml
│   │   └── service.yaml
│   └── overlays
│       ├── dev
│       │   ├── kustomization.yaml
│       │   ├── replicas.yaml
│       │   └── variables.configmap
│       └── prod
│           ├── kustomization.yaml
│           ├── replicas.yaml
│           └── variables.configmap
├── plain-kubernetes
│   ├── configmap.yaml
│   ├── deployment.yaml
│   ├── namespace.yaml
│   └── service.yaml
└── README.md
```

## Create a token for your container registry:

The `patients` git repository contains instructions for CI system to build a docker/container image and push it to a container registry. This means, you need:
* a container registry URL
* a container registry username
* a PAT (Personal Access Token) to be used as password when connecting to the container registry

In this example, I am using docker.io as container registry, so I will create a PAT under my username and save it on my disk at a safe location. (`/home/kamran/Keys-and-Tokens/github/dockerhub-token-for-github-actions.pat`)

You will need this PAT in the next step.

## Create git repositories at your git provider:
It's time to create two repositories at your git provider, under your personal or organization account. 


* patients (https://github.com/KamranAzeem/patients/settings)
* patients-gitops (https://github.com/KamranAzeem/patients-gitops)


Once you create the repositories, create the following variables under  `Settings -> Security -> Secrets and Variables -> Actions -> Secrets`, and set it to the value of PAT you obtained in the previous step.

```
CONTAINER_REGISTRY_URL=docker.io
CONTAINER_REGISTRY_USERNAME=your-username
CONTAINER_REGISTRY_TOKEN=token-from-docker-hub
```



This is important, because as soon as you push code to the remote repository, the CI instructions will kick-in and a docker image build process will kick off. These steps expect certain variables to be in place otherwise the step to push image to container registry will fail.

Once this is in place, you are ready to push your code from local computer to the remote repository .

## Setup your local code directories as git repositories:

Next, take out the `patients` and `patients-gitops` directories outside of this (learn-kustomize) directory tree and create/set them up as two individual/independent git repositories somewhere else on your file system. 

```
$ mkdir /home/kamran/Projects/Personal/github/patients

$ mkdir /home/kamran/Projects/Personal/github/patients-gitops
```

```
$ cp -a patients/. /home/kamran/Projects/Personal/github/patients/.

$ cp -a patients-gitops/. /home/kamran/Projects/Personal/github/patients-gitops/.
```


Turn these directories into git repositories and setup remote:

```

```




## Prepare the docker/container image - manually:

First, build a container image locally to ensure that your app has no problems building.
```

[kamran@kworkhorse patients]$ docker build -t local/patients .
 => => extracting sha256:c185020e1367a154f8bf1d26293b7c6fac06249fb9c7662388e1566fff8098da                                          2.1s
 => => extracting sha256:2302088a13ce9bf346485380cdb2c1540483f344115926186083e761f2a0f48e                                          3.3s
 => => extracting sha256:2755303bf1ef081c6ea63474816ac42791402933bd19caa54cd520e9c4bd88e7                                          0.0s
 => [stage-1 1/2] WORKDIR /app                                                                                                     0.0s
 => [internal] load build context                                                                                                  0.0s
 => => transferring context: 460B                                                                                                  0.0s
 => [builder 2/4] WORKDIR /app                                                                                                     5.6s
 => [builder 3/4] COPY src/ /app                                                                                                   0.1s
 => [builder 4/4] RUN CGO_ENABLED=0 go build main.go                                                                              11.4s
 => [stage-1 2/2] COPY --from=BUILDER    /app    /app                                                                              0.1s
 => exporting to image                                                                                                             0.1s
 => => exporting layers                                                                                                            0.1s
 => => writing image sha256:4d524e16da866897582b3fd61ed1f9707d771a67ee16365e02fd54fba53145f9                                       0.0s
 => => naming to docker.io/local/patients                                                                                          0.0s
[kamran@kworkhorse patients]$
```

## Setup the image to be built automatically using CI/CD:
Now we need a GitHub Ations file with steps to build the container image and push to container registry.

The steps to do this are listed in the `.github/workflows/build-and-push-docker-image.yml` file.

You need to test three scenarios:
* If a simple commit is pushed to git repo, a corresponding container image with the git hash of the latest commit as it's image tag. This container image will also be tagged with `latest`.
* If you tagged a commit with `rc-#-#-#` (e.g. rc-1.0.0), then a container image should be build for it, and it should be visible/accessible in docker hub with this tag.
* If you tagged a commit with `v#-#-#` (e.g. v1.0.1), then a container image should be build for it, and it should be visible/accessible in docker hub with this tag.

The above (git tagging) scheme will help us deploy different versions of application in different environments. i.e: 
* the latest and greatest in the `dev` environment
* `rc-#-#-#` in the `demo` (aka `staging`)  environment
* `v#.#.#` in the `prod` environment

Try adding few commits to the main (patients) application, add different tags and ensure that you see those variants in your container registry.


