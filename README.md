# simple-webapp
Simple web app to read and write data to MySQL

### Assumptions

- Working kubernetes cluster
- Installed and configured kubectl

## V1

Simple Web Application V1 creates below Kubernetes configuration


| Namespace | webapp |
| --------- | -------|
| Secret | dbsecret |
| Replicaset | selector - webapp |
| Service | Nodeport - 30327 |

