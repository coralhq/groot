# groot

> read key/value pairs from etcd and convert it to dotenv file with expansion

## installation

```
go get -u -v github.com/coralhq/groot
go install github.com/coralhq/groot
```

## sample use case

```
## etcd contents

/env/global/GITHUB_USERNAME
akhy

/env/global/GITHUB_TOKEN
abcdefghijkl

/env/myservice/GITHUB_URL
https://github.com/coralhq/groot

/env/myservice/GITHUB_AUTH
${GITHUB_USERNAME}:${GITHUB_TOKEN}
```


```sh
~$ export ETCD_URLS="https://etcd-01:2379,https://etcd-02:2379"
~$ export ETCD_BASE_DIR="/env/global"
~$ export ETCD_ENV_DIR="/env/myservice"

~$ groot
GITHUB_URL="https://github.com/coralhq/groot"
GITHUB_TOKEN="akhy:abcdefghijkl"

~$ groot --export
export GITHUB_URL="https://github.com/coralhq/groot"
export GITHUB_TOKEN="akhy:abcdefghijkl"
```

### generating dotenv file

```sh
groot > .env
```

### exporting results in current shell

```sh
eval $(groot --export)
```
