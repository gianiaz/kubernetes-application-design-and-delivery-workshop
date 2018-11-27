build:
	docker pull mongo:4.1.5-xenial
	docker tag mongo:4.1.5-xenial sighup/kubeprimer-mongo

run:
	docker rm -f kubeprimer-mongo || true
	docker run -d --name kubeprimer-mongo --network=kubeprimer -p 27017:27017 sighup/kubeprimer-mongo
