#!/bin/bash
docker build -t timvw/demo .
docker push timvw/demo:latest
sleep 1
kubectl rollout restart deployment/demo
sleep 2
kubectl get pods | grep demo | awk '{print $1}' | gxargs -l kubectl logs -f
