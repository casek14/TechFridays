package main

import (
	"context"
	"fmt"

	"go.etcd.io/etcd/v3/client"
	"log"
	"time"
)

type etcdClient struct {
	ctx context.Context
	kv  client.KeysAPI
}

func NewEtcdClient(ctx context.Context, cli client.KeysAPI) *etcdClient {

	return &etcdClient{
		ctx: ctx,
		kv:  cli,
	}
}

// delete all keys
func (e *etcdClient) deleteAllKeys(key string) {
	log.Println("Deleting all keys ...")

	e.kv.Delete(e.ctx, key, &client.DeleteOptions{})
}

// Insert key into etcd
func (e *etcdClient) insertKey(key, value string) {
	res, err := e.kv.Set(e.ctx, "karel", "12345", &client.SetOptions{TTL: time.Second * 10})
	if err != nil {
		log.Printf("Error in inserting key into etcd: %s\n", err.Error())
		return
	}

	fmt.Printf("Key saved %s with action %s\n", key, res.Action)
}

func (e *etcdClient) getKey(key string) {
	res, err := e.kv.Get(e.ctx, "karel", &client.GetOptions{})
	if err != nil {
		fmt.Printf("Error in get key %s : %s\n", key, err.Error())
	}

	if res == nil {
		log.Printf("Empty response GET KEY")
		return
	}
	log.Printf("Key: %s  --- Value: %s\n", res.Node.Key, res.Node.Value)

}

/*
func (e *etcdClient) watchForExpired(keyPrefix string) <-chan *client.Response {
	reponses := make(chan *client.Response)
	watcher := e.kv.Watcher()
}
*/
func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	cfg := client.Config{Endpoints: []string{"http://127.0.0.1:2379"}, Transport: client.DefaultTransport}

	cli, err := client.New(cfg)
	if err != nil {
		log.Fatalf("Unable to create client: %s\n", err.Error())
	}

	kv := client.NewKeysAPI(cli)
	etcdCli := NewEtcdClient(ctx, kv)

	key := "karel"
	value := "12345"
	etcdCli.deleteAllKeys(key)
	etcdCli.getKey(key)
	etcdCli.insertKey(key, value)
	etcdCli.getKey(key)
	time.Sleep(15 * time.Second)
	fmt.Printf("Checking key %s after its 10s TTL expired\n", key)
	etcdCli.getKey(key)
}
