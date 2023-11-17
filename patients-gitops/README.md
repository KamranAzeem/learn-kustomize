# hello-world-golang (configuration)

This repository does not contain actual application code for the hello-world application. It only contains some kubernetes manifests to deploy the application, through argocd.

# Create github credentials:

ArgoCD image updater will need a writeback method to be able to update the image in the configuration repository. To be able to do this, it needs Git username and git password/token. You create these in the github web-ui and then make a note of it, and then create a kubernetes secret as follows:


Assume the token is created and is saved in the local computer as the file `/home/kamran/Keys-and-Tokens/github/github-ARGOCD_TEST.pat` .

Setup two ENV variables on your shell:

```
export GITHUB_USERNAME=kamranazeem

export GITHUB_TOKEN=$(cat /home/kamran/Keys-and-Tokens/github/github-ARGOCD_TEST.pat)

kubectl --namespace argocd \
  create secret generic github-credentials \
  --from-literal=username=${GITHUB_USERNAME} \
  --from-literal=password=${GITHUB_TOKEN}
```

