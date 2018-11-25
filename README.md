# Kubernetes Application Design Workshop by [SIGHUP](https://sighup.io)

Welcome to the Kubernetes Application Design and Delivery workshop. This workshop is organized by [SIGHUP](https://sighup.io) and aims to provide a comprehensive deep dive on the patterns and best practices that could be adopted when developing distributed applications on Kubernetes.

## Workshop info

The first edition of this workshop has been develivered at [Codemotion Milan 2018](https://milan2018.codemotionworld.com/workshop/kubernetes-application-design-and-delivery/).

In this repo you will find all the code used during the workshop. In case you are interested in the workshop, [email us](mailto:training@sighup.io). We have scheduled workshops and deliver on-site sessions for interested organisations.

### Abstract and Agenda

In this hands-on workshop we are going to dive deep into the principles, development patterns and best practices to adopt while developing and deploying your services to Kubernetes. For this workshop you must have a good understanding of the core domain concepts of Kubernetes as we will focus on more advanced aspects for both stateless and stateful applications. We will see what best practices and design patterns to adopt while developing applications, how to make them scalable, how to integrate and monitor your applications with Prometheus and what logging strategies to ensure when developing container based applications. We will also focus on security with a clear overview of the Kubernetes RBAC model and how to secure and isolate correctly your applications.

## Program

1. A quick review of the Kubernetes core domain concepts  
2. Troubleshooting and debugging Applications on Kubernetes  
3. Cloud Native application design: architectural patterns for distributed systems  
4. Container security, effective RBAC and secrets management  
5. Distribution and Delivery of cloud native applications  
6. Kubernetes monitoring principles, exporting application metrics for Prometheus  
7. Rock solid scalable logging strategies  
8. Running stateful applications in Kubernetes: Operators and the current state of databases in Kubernetes  

You can contact the authors of this workshop by email at [training@sighup.io](mailto:training@sighup.io)

### Setup & Environment

During the workshop, we encourage participants to work in pair. Therefore at least one computer per couple should be available and correctly configured.

#### Kubernetes

You should have access to a Kubernetes cluster. You can very easily provision a single-node cluster on your machine with `minikube` as well can use any other Kubernetes cluster from public clouds or self-provisioning.

##### Minikube
Here you can find [installation instructions](https://github.com/kubernetes/minikube#installation) for your system to easily get `minikube` up and running. Note that you will need to have a hypervisor installed (something like Virtualbox).

For macOS:  

```shell
brew cask install minikube
```

For Linux:  
```shell
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 \
  && sudo install minikube-linux-amd64 /usr/local/bin/minikube
```

#### Kubectl
It's necessary to have `kubectl` installed. `Kubectl` is the CLI tool to interact with Kubernetes.

On macOS:  
```shell
brew install kubernetes-cli
```

On Linux Ubuntu/Debian:  
```shell
sudo apt-get update && sudo apt-get install -y apt-transport-https
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee -a /etc/apt/sources.list.d/kubernetes.list
sudo apt-get update
sudo apt-get install -y kubectl
```

For other OS, refer to the [official install docs](https://kubernetes.io/docs/tasks/tools/install-kubectl/#install-kubectl)

As a bonus, you should install `kubectx` and `kubens` [(installation instructions here)](https://github.com/ahmetb/kubectx#installation), they are not strictly required but are actually priceless, so... why not.

#### Helm and Kustomize

Install `kustomize` and `helm`. To install Kustomize, [here you can find the installation instructions](https://github.com/kubernetes-sigs/kustomize/blob/master/docs/INSTALL.md). Same goes for [helm](https://github.com/helm/helm#install)

### Using Minikube

On macOS:  
```shell
```

On Linux:  
```shell
```
