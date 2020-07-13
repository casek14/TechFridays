package main

import (
	"context"
	"fmt"

	"log"
	"time"
	"github.com/coreos/etcd/client"
)

type etcdClient struct {
	ctx         context.Context
	kv          client.KeysAPI
}

func NewEtcdClient(ctx context.Context, cli client.KeysAPI) *etcdClient {

	return &etcdClient{
		ctx:         ctx,
		kv:          cli,
	}
}

// delete all keys
func (e *etcdClient) deleteAllKeys() {
	log.Println("Deleting all keys ...")

	e.kv.Delete(e.ctx,"key",&client.DeleteOptions{} )
}

// Insert key into etcd
func (e *etcdClient) insertKey(key, value string) {
	res, err := e.kv.Set(e.ctx,"karel","12345",&client.SetOptions{})
	if err != nil {
		log.Printf("Error in inserting key into etcd: %s\n", err.Error())
		return
	}

	fmt.Printf("Key saved %s with action %s\n", key, res.Action)
}

func (e *etcdClient) getKey(key string) {
	res, err := e.kv.Get(e.ctx, "karel",&client.GetOptions{})
	if err != nil {
		fmt.Printf("Error in get key %s : %s\n",key, err.Error())
	}

	if res == nil {
		log.Printf("Empty response GET KEY")
		return
	}
	log.Printf("Key: %s  --- Value: %s\n",res.Node.Key, res.Node.Value)

}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	cfg := client.Config{Endpoints: []string{"127.0.0.1:2379"}, Transport: client.DefaultTransport}

	cli, _ := client.New(cfg)

	kv := client.NewKeysAPI(cli)
	etcdCli := NewEtcdClient(ctx, kv)

	etcdCli.deleteAllKeys()
	etcdCli.getKey("karel")
	etcdCli.insertKey("karel","12345")
	etcdCli.getKey("karel")
}
