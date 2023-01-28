export PATH :=$(PATH):$(GOPATH)/bin

APP=$(shell basename "$(PWD)")
POD=$(shell kubectl get pod -l project=${APP} -o jsonpath="{.items[0].metadata.name}")

help:   ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

start: ## start one-node local cluster
	minikube start
stop: ## stop one-node local cluster
	minikube stop

delete: ## delete all dev cluster instances
	kubectl delete -f deployments/dev --ignore-not-found=true

build-image: ## build image and restart app
	minikube image build -t ${APP}:latest -f build/Dockerfile .
	kubectl apply -f deployments/dev
	kubectl rollout restart deployment/${APP}

run: ##  run instance
	minikube service ${APP}

serve: start build-image run ## build and run application

exec: ## enter to app container
	kubectl exec -it ${POD} -- /bin/sh

app-logs: ## show app logs in console
	kubectl logs po/${POD} -f