# simple-webapp
Simple web app to read and write data to MySQL

### Assumptions

- Working kubernetes cluster
- Installed and configured kubectl

## V1

For exact instructions, please visit below URL 

```
https://medium.com/devopslinks/delivering-secrets-management-for-database-connectivity-with-vault-part-1-1f1ca2dbfd8c
```

Simple Web Application V1 creates below Kubernetes configuration

| Configuration | Default |
| ------------- | ------- |
| Namespace | webapp |
| Secret | dbsecret |
| Replicaset | selector - webapp |
| Service | Nodeport - 30327 |

Installation
```
kubectl apply -f simple-webapp.yaml
```

## V2

For exact instructions, please visit below URL 

```
https://medium.com/devopslinks/delivering-secrets-management-for-database-connectivity-with-vault-part-2-b8dfd782b9b1
```

Simple Web Application V2 creates below Kubernetes configuration 

| Configuration | Default |
| ------------- | ------- |
| Namespace | webapp-v2 |
| Replicaset | simple-webapp-dbbackend |
| Replicaset | simple-webapp-frontend  |
| Service | simple-webapp - Nodeport - 30328 |
| Service | dbbackend - ClusterIP |
