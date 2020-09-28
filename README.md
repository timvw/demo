
```bash
docker build -t timvw/demo .
docker run --rm -p 80:80  timvw/demo
```


```
kubectl create deployment demo --image=timvw/demo
kubectl expose deployment demo --port=80
kubectl apply -f ingress.yaml
```

```
docker build -t timvw/demo . 
docker push timvw/demo:latest
kubectl rollout restart deployment/demo
kubectl get pods | grep demo | awk '{print $1}' | gxargs -l kubectl logs -f
echo "done"
```

```bash
git tag "test-$(date +%s)" && git push --tags
```


curl http://localhost/github -X POST -d @testbody.txt -H "Content-Type: application/json"






