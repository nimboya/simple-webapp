# Install helm depenedencies
kubectl -n kube-system create serviceaccount tiller
kubectl create clusterrolebinding tiller \\n  --clusterrole cluster-admin \\n  --serviceaccount=kube-system:tiller
# Assuming that helm is already installed and is in PATH
helm init --service-account tiller
kubectl -n kube-system  rollout status deploy/tiller-deploy
helm repo add incubator http://storage.googleapis.com/kubernetes-charts-incubator
helm repo update