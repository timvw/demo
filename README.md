
```bash
docker build -t timvw/demo .
docker run --rm -p 80:80  timvw/demo
```


```
kubectl create deployment demo --image=timvw/demo
kubectl expose deployment demo --port=80
kubectl apply -f ingress.yaml
```