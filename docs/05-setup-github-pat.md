# Setup a Github PAT to help ArgoCD:

ArgoCD image updater will need a writeback method to be able to update the image in the "app-gitops" repository. To be able to do this, it needs Git username and git password/token. You create these in the github web-ui and then make a note of it in a file on your local computer ,e.g. in `/home/kamran/Keys-and-Tokens/github/github-ARGOCD_TEST.pat` .

Then create a kubernetes secret as follows.


Setup two ENV variables on your shell:

```
export GITHUB_USERNAME=kamranazeem

export GITHUB_TOKEN=$(cat /home/kamran/Keys-and-Tokens/github/github-ARGOCD_TEST.pat)
```

Create the secret:
```
kubectl --namespace argocd \
  create secret generic github-credentials \
  --from-literal=username=${GITHUB_USERNAME} \
  --from-literal=password=${GITHUB_TOKEN}
```

Verify:
```
[kamran@kworkhorse github]$ kubectl -n argocd get secret
NAME                          TYPE     DATA   AGE
argocd-initial-admin-secret   Opaque   1      3d
argocd-notifications-secret   Opaque   0      3d1h
argocd-secret                 Opaque   5      3d1h
github-credentials            Opaque   2      16s
[kamran@kworkhorse github]$
```
