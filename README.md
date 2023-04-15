# resilience-testing-example
This repository contains an example for resilience testing in Go service.

## Launch cluster with Cilium network plugin
1. `minikube start --network-plugin=cni --cni=false`
2. Install Cilium from [Docs](https://docs.cilium.io/en/stable/gettingstarted/k8s-install-default/)
3. `cilium install`
4. `cilium status --wait`


## Deploy PoC application
1. `eval $(minikube docker-env)`
2. `make dockerbuild`
3. `kubectl create -f deployments/server.yaml`
4. `kubectl create -f deployments/worker.yaml`
5. Inspect pod logs in `k9s`

