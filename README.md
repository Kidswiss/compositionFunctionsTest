# Commands

build
```
docker build . -t kidswiss/myfunc:latest
```

Test in docker
```
yq -o=json '.' redisFunctionIO.yaml | docker run -i kidswiss/myfunc:latest | jq
```

Test with crossplane runner
```
yq -o=json '.' redisFunctionIO.yaml |docker run -i --security-opt=seccomp=unconfined crossplane/xfn:v1.11.1 run -c /tmp kidswiss/myfunc:latest -
```

Bug: https://github.com/crossplane/crossplane/issues/3807
