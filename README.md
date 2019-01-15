# kodemo
A multi-part demo of ko

## Script


```
# First show the app source.
cat ./cmd/myapp/main.go

# Show help for kubectl run
kubectl help run | less

# Then "run" the app on Kubernetes.
time ko run myapp --image=./cmd/myapp

# Show the running deployment
kubectl get deploy

# Expose the service so that we can see it.
# Pro-tip: this is long, have a "chicken" ready
kubectl expose deployment myapp --port=80 --target-port=8080 --type=LoadBalancer

# Wait for public IP
watch kubectl get svc

# Show the service's output
curl http://{THE_IP_ADDRESS}

```
