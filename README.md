# crd-nginx-vhost

add vhost by crd

## usages 

* go modules

```
# go mod vendor #在mod模式下添加旧的目录
# go mod tidy
```

* codegen

```
# cd hack
# ./update-codegen.sh
```

* build crd

```
# go build -o vhost .
```

* check vhost (no need)

```
# ./vhost --kubeconfig=/Users/wangsir/.kube/config
```

* create crd

```
# kubectl create -f example/crd.yml
```

* create vhost

```
kubectl create -f example/example-vhost.yml 
```

## issues

```bash
 ✗ go build -o vhost .
# k8s.io/client-go/rest
../../../pkg/mod/k8s.io/client-go@v11.0.0+incompatible/rest/request.go:598:31: not enough arguments in call to watch.NewStreamWatcher
        have (*versioned.Decoder)
        want (watch.Decoder, watch.Reporter)
```

client-go 版本兼容问题，可通过以下方式更改版本


```
# go get -u k8s.io/client-go@master
# go get -u k8s.io/client-go@kubernetes-1.15.0 #指定k8s版本

--- k8s.io/client-go v11.0.0+incompatible
+++ k8s.io/client-go v0.0.0-20190831074946-3fe2abece89e
```
