package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/coreos/etcd/client"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

func createEnvMap(kapi client.KeysAPI, prefix string) map[string]string {
	recursive := client.GetOptions{Recursive: true}
	resp, err := kapi.Get(context.Background(), prefix, &recursive)
	if err != nil {
		log.Fatal(err)
	}
	base := make(map[string]string)
	for _, node := range resp.Node.Nodes {
		if node.Dir {
			continue
		}
		key := strings.TrimPrefix(node.Key, prefix+"/")
		base[key] = node.Value
	}

	return base
}

var (
	export  bool
	urls    string
	baseDir string
	envDir  string
)

func main() {
	flag.BoolVar(&export, "export", false, "prepend `export` to the lines")
	flag.Parse()

	viper.SetEnvPrefix("etcd")
	viper.AutomaticEnv()
	viper.SetDefault("urls", "http://127.0.0.1:2379")
	urls = viper.GetString("urls")
	baseDir = viper.GetString("base_dir")
	envDir = viper.GetString("env_dir")

	cfg := client.Config{
		Endpoints:               strings.Split(urls, ","),
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)

	base := createEnvMap(kapi, baseDir)
	env := createEnvMap(kapi, envDir)

	for key, tpl := range env {
		val := os.Expand(tpl, func(k string) string {
			return base[k]
		})

		if export {
			fmt.Printf("export %s=\"%s\"\n", key, val)
		} else {
			fmt.Printf("%s=\"%s\"\n", key, val)
		}
	}
}
