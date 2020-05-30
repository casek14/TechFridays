# Deploy and init vault
```
kubectl apply -f vault.yaml

vault operator init -format=json > cluster-keys.json

# unseal the vault with keys from cluter-keys.json file
vault operator unseal
```


## Enable key-value 
