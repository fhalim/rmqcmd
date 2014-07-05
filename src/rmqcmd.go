package main

import (
	"encoding/json"
	"fmt"
	"github.com/michaelklishin/rabbit-hole"
	"github.com/voxelbrain/goptions"
)

type RabbitMQAdminOpts struct {
	AdminUrl string        `goptions:"-au, --adminUrl, description='Admin URL'"`
	Username string        `goptions:"-u, --password, description='Username'"`
	Password string        `goptions:"-p, --password, description='Password'"`
	Vhost    string        `goptions:"-v, --vhost, description='Vhost'"`
	Help     goptions.Help `goptions:"-h, --help, description='Show this help'"`

	Verb       goptions.Verbs
	ListQueues struct {
	} `goptions:"listqueues"`
	ListExchanges struct {
	} `goptions:"listexchanges"`
	ListBindings struct {
	} `goptions:"listbindings"`
	ListNodes struct {
	} `goptions:"listnodes"`
}

func parseOptions() RabbitMQAdminOpts {
	options := RabbitMQAdminOpts{
		AdminUrl: "http://localhost:15672",
		Username: "guest",
		Password: "guest",
		Vhost:    "/",
	}
	goptions.ParseAndFail(&options)
	return options
}
func print(object interface{}) {
	byteArray, _ := json.Marshal(object)
	fmt.Println(string(byteArray))
}
func printList(objects []interface{}) {
	for idx := range objects {
		node := objects[idx]
		print(node)
	}
}
func listQueues(vhost string, client *rabbithole.Client) {
	queues, _ := client.ListQueuesIn(vhost)
	printList(queues)
}
func listExchanges(vhost string, client *rabbithole.Client) {
	exchanges, _ := client.ListExchangesIn(vhost)
	printList(exchanges)
}
func listBindings(vhost string, client *rabbithole.Client) {
	bindings, _ := client.ListBindingsIn(vhost)
	printList(bindings)

}
func listNodes(client *rabbithole.Client) {
	nodes, _ := client.ListNodes()
	printList(nodes)
}

func main() {
	options := parseOptions()
	rmqc, _ := rabbithole.NewClient(options.AdminUrl, options.Username, options.Password)

	verb := options.Verb
	switch {
	case verb == "listqueues":
		listQueues(options.Vhost, rmqc)
	case verb == "listexchanges":
		listExchanges(options.Vhost, rmqc)
	case verb == "listbindings":
		listBindings(options.Vhost, rmqc)
	case verb == "listnodes":
		listNodes(rmqc)
	default:
		fmt.Println("Unrecognized command")
	}
}
