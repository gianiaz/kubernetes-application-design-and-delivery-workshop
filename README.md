# Kubernetes Application Design

Welcome to the Kubernetes Application Design and Delivery workshop. This workshop is organized by [SIGHUP](https://sighup.io) and aims to provide a comprehensive deep dive on the patterns and best practices that could be adopted when developing distributed applications on Kubernetes.

## Workshop info

The first edition of this workshop has been develivered at [Codemotion Milan 2018](https://milan2018.codemotionworld.com/workshop/kubernetes-application-design-and-delivery/).

In this repo you will find all the code used during the workshop. In case you are interested in the workshop, [email us](mailto:training@sighup.io). We have scheduled workshops and deliver on-site sessions for interested organisations.

### Abstract and Agenda

In this hands-on workshop we are going to dive deep into the principles, development patterns and best practices to adopt while developing and deploying your services to Kubernetes. For this workshop you must have a good understanding of the core domain concepts of Kubernetes as we will focus on more advanced aspects for both stateless and stateful applications. We will see what best practices and design patterns to adopt while developing applications, how to make them scalable, how to integrate and monitor your applications with Prometheus and what logging strategies to ensure when developing container based applications. We will also focus on security with a clear overview of the Kubernetes RBAC model and how to secure and isolate correctly your applications.

#### Program

– A quick review of the Kubernetes core domain concepts  
- Troubleshooting and debugging Applications on Kubernetes  
– Cloud Native application design: architectural patterns for distributed systems  
– Container security, effective RBAC and secrets management  
- Distribution and Delivery of cloud native applications  
– Kubernetes monitoring principles, exporting application metrics for Prometheus  
– Rock solid scalable logging strategies  
– Running stateful applications in Kubernetes: Operators and the current state of databases in Kubernetes  

You can contact the authors of this workshop by email at [training@sighup.io](mailto:training@sighup.io)

### Setup & Environment

During the workshop, we encourage participants to work in pair. Therefore at least one computer per couple should be available and correctly configured.

#### Kubernetes

1. Each couple should have access to a Kubernetes cluster. This could be [minikube](https://github.com/kubernetes/minikube#installation) or [Kubernetes on docker](https://www.docker.com/get-started) to have kubernetes running on your machine. It could also be a public cloud cluster or some cluster you decided to setup yourself. It's important that you have full unrestricted access to the test cluster you will use. For a simple and hassle-free setup we recommend using [minikube](https://github.com/kubernetes/minikube#installation).

2. Have [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/#install-kubectl) installed

3. Install `kubectx` and `kubens` [(installation instructions here)](https://github.com/ahmetb/kubectx#installation), they are not strictly required but are actually priceless, so... why not.

4. Install `kustomize` and `helm`. To install Kustomize, [here you can find the installation instructions](https://github.com/kubernetes-sigs/kustomize/blob/master/docs/INSTALL.md). Same goes for [helm](https://github.com/helm/helm#install)

