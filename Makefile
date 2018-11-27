#!/bin/bash
dev-machine:
	minikube start --cpus=2 --memory=4096 --kubernetes-version=v1.11.4

prepare:
	scp -ri `minikube ssh-key` app docker@`minikube ip`:~
	ssh -i `minikube ssh-key` docker@`minikube ip` 'curl -L "https://github.com/docker/compose/releases/download/1.23.1/docker-compose-`uname -s`-`uname -m`" -o docker-compose && chmod +x ./docker-compose'
	ssh -i `minikube ssh-key` docker@`minikube ip` './docker-compose -f app/docker-compose.yml pull'

run-compose:
	ssh -i `minikube ssh-key` docker@`minikube ip` './docker-compose -f app/docker-compose.yml up -d'

stop-compose:
	ssh -i `minikube ssh-key` docker@`minikube ip` './docker-compose -f app/docker-compose.yml down'
