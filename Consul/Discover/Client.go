package Discover

import (
	"github.com/hashicorp/consul/api"
)

type Client interface{

	Register(tags []string) error
    Service(service,tag string)([]*api.ServiceEntry,*api.QueryMeta,error)
}

type client struct{
	client *api.Client
	address string
	name string
	port int
}

func NewClient(config *api.Config,address string,name string,port int)(*client,error){
	c,err := api.NewClient(config)

	if err != nil{
		return nil,err
	}

	ClientEntry := new(client)
	ClientEntry.client = c
	ClientEntry.name = name
	ClientEntry.address = address
	ClientEntry.port = port

	return ClientEntry,nil


}
