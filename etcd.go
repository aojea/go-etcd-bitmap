package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

type EtcdBackend struct {
	client *clientv3.Client
}

func NewEtcdBackend(endpoints []string) (*EtcdBackend, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return &EtcdBackend{
		client: cli,
	}, nil
}