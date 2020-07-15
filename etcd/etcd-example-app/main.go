package main

import (
	"context"
	"fmt"

	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

// keyPrefix - used for select in etcd db
type etcdClient struct {
	ctx       context.Context
	cli       clientv3.Client
	kv        clientv3.KV
	keyPrefix string
	valuesTTL int64
}

// Create new etcd cklient
func NewEtcdClient(ctx context.Context, kv clientv3.KV, cli clientv3.Client) *etcdClient {
	return &etcdClient{
		ctx:       ctx,
		kv:        kv,
		cli:       cli,
		keyPrefix: "kta-",
		valuesTTL: 60,
	}
}

// delete all keys
func (e *etcdClient) deleteAllKeys(key string) {
	log.Println("Deleting all keys ...")

	e.kv.Delete(e.ctx, e.keyPrefix, clientv3.WithPrefix())
}

// Insert key into etcd
func (e *etcdClient) insertKey(key, value string) {
	// create lease, TTL, for the key
	lease, err := e.cli.Grant(e.ctx, e.valuesTTL)
	if err != nil {
		log.Printf("Error in generation lease: %s", err.Error())
	}
	res, err := e.kv.Put(e.ctx, e.keyPrefix+key, value, clientv3.WithLease(lease.ID))
	if err != nil {
		log.Printf("Error in inserting key into etcd: %s\n", err.Error())
		return
	}

	fmt.Printf("Key saved %s with action %s\n", key, res.Header.String())
}

func (e *etcdClient) getKey(key string) {
	res, err := e.kv.Get(e.ctx, "karel", clientv3.WithPrevKV())
	if err != nil {
		fmt.Printf("Error in get key %s : %s\n", key, err.Error())
	}

	if res == nil {
		log.Printf("Empty response GET KEY")
		return
	}
	for _, kv := range res.Kvs {
		log.Printf("Key: %s  --- Value: %s\n", kv.Key, kv.Value)
	}

}

func (e *etcdClient) createLease(cli *clientv3.Client) clientv3.LeaseID {
	lease, err := cli.Grant(e.ctx, 60)
	if err != nil {
		log.Printf("Unable to create lease: %s", err.Error())
	}
	fmt.Printf("Lease created :%+v\n", lease.ID)
	return lease.ID
}

// watch expired events from etcd and send them to returned channel
func (e *etcdClient) watchForExpired(cli *clientv3.Client) <-chan *clientv3.Event {
	events := make(chan *clientv3.Event)
	wCh := cli.Watch(e.ctx, e.keyPrefix, clientv3.WithPrefix(), clientv3.WithFilterPut())
	go func() {
		defer close(events)
		for wResp := range wCh {
			for _, ev := range wResp.Events {
				expired, err := e.isExpired(ev)
				if err != nil {
					log.Printf("Error in checking event expiration: %s\n", err.Error())
				} else if expired {
					events <- ev
				}

			}
		}
	}()

	return events
}

// Check if etcd generated event is of type expired
func (e *etcdClient) isExpired(ev *clientv3.Event) (bool, error) {
	if ev.Kv == nil {
		log.Println("EVENT KV is NILL")
		return false, nil
	}
	log.Printf("EVENT: %+v\n",ev.Kv)
	leaseID := clientv3.LeaseID(ev.Kv.Lease)
	if leaseID == clientv3.NoLease {
		log.Printf("LEASEID == clientv3.NoLease, LEASE-ID: %+v\n",leaseID)
		return true, nil
	}

	ttlResp, err := e.cli.TimeToLive(e.ctx, leaseID)
	if err != nil {
		log.Println("ERR IN checking TTL")
		return false, err
	}
	log.Printf("TTL values: %d\n",ttlResp.TTL)
	return ttlResp.TTL == 1, nil
}

func main() {
	// Do not use context with timeout (timeout kills expire events channel)
	ctx := context.Background()

	// config for etcd connection
	cfg := clientv3.Config{Endpoints: []string{"http://127.0.0.1:2379"}, DialTimeout: time.Second * 10}

	// new etcd client
	cli, err := clientv3.New(cfg)
	if err != nil {
		log.Fatalf("Unable to create client: %s\n", err.Error())
	}
	defer cli.Close()

	// new etcd key-value client
	kv := clientv3.NewKV(cli)

	etcdCli := NewEtcdClient(ctx, kv, *cli)

	// delete all keys, check existence, create new record check existence again
	key := "karel"
	value := "12345"
	etcdCli.deleteAllKeys(key)
	etcdCli.getKey(key)
	etcdCli.insertKey(key, value)
	etcdCli.getKey(key)

	// Renew lease, so value should stay for 90s (original ttl is 60s)
	go func() {
		time.Sleep(time.Second * 30)
		log.Printf("Key %s RENEWAL", etcdCli.keyPrefix+key)
		//etcdCli.insertKey(key, value)
	}()

	// wait for expire events
	for {
		select {
		case ev := <-etcdCli.watchForExpired(cli):
			log.Println("Expired event occured !!!")
			log.Printf("Received event is type: %T and value: %+v\n", ev, ev)
			fmt.Printf("Expire event for key %s and value %s", ev.Kv.Key, ev.Kv.Value)
		}
	}
}
