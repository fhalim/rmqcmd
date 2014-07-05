package main

import (
	"fmt"
	"github.com/michaelklishin/rabbithole"
	"github.com/voxelbrain/goptions"
)

type RabbitMQAdminOpts struct {
	AdminUrl   string        `goptions:"-au, --adminUrl, description='Admin URL'"`
	Username   string        `goptions:"-u, --password, description='Username'"`
	Password   string        `goptions:"-p, --password, description='Password'"`
	Vhost   string        `goptions:"-v, --vhost, description='Vhost'"`
	Help       goptions.Help `goptions:"-h, --help, description='Show this help'"`

	Verb    goptions.Verbs
	ListQueues struct {
	} `goptions:"listqueues"`
	ListExchanges struct {
	} `goptions:"listexchanges"`
	ListBindings struct {
	} `goptions:"listbindings"`

}

func parseOptions() RabbitMQAdminOpts {
	options := RabbitMQAdminOpts {
		AdminUrl: "http://localhost:15672",
		Username: "guest",
		Password: "guest",
		Vhost: "/",
	}
	goptions.ParseAndFail(&options)
	return options
}

func listQueues(vhost string, client *rabbithole.Client) {
	queues, _ := client.ListQueuesIn(vhost)
	for idx := range queues {
		fmt.Println("name: %s\n", queues[idx].Name)
	}
}
func listExchanges(vhost string, client *rabbithole.Client) {
	exchanges, _ := client.ListExchangesIn(vhost)
	for idx := range exchanges {
		fmt.Printf("name: %s\n", exchanges[idx].Name)
	}
}
func listBindings(vhost string, client *rabbithole.Client) {
	bindings, _ := client.ListBindingsIn(vhost)
	for idx := range bindings {
		binding := bindings[idx]
		fmt.Printf("source: %s, destination: %s, type: %s, routingkey: %s\n", binding.Source, binding.Destination, binding.DestinationType, binding.RoutingKey)
	}
}
func main() {
	options := parseOptions()
	rmqc, _ := rabbithole.NewClient(options.AdminUrl, options.Username, options.Password)

	verb := options.Verb
	switch {
	case verb == "listqueues": listQueues(options.Vhost, rmqc)
	case verb == "listexchanges": listExchanges(options.Vhost, rmqc)
	case verb == "listbindings": listBindings(options.Vhost, rmqc)
	default:
		fmt.Println("Unrecognized command")
	}
}
