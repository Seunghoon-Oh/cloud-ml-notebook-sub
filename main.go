package main

import (
	"log"
	"sync"

	"github.com/nats-io/nats.go"
)

// NOTE: Can test with demo servers.
// nats-sub -s demo.nats.io <subject>

// func printMsg(m *nats.Msg, i int) {
// 	log.Printf("[#%d] Received on [%s]: '%s'", i, m.Subject, string(m.Data))
// }

func main() {
	log.SetFlags(0)
	// flag.Parse()
	// args := flag.Args()

	// Connect Options.
	// opts := []nats.Option{nats.Name("NATS Subscriber")}
	// opts = setupConnOptions(opts)
	for {
		// time.Sleep(1 * time.Second)
		// // Connect to NATS
		// // nc, err := nats.Connect("nats://nats.cloud-ml-mgmt:4222")
		// nc, err := nats.Connect(nats.DefaultURL)
		// log.Println("Connected")
		// if err != nil {
		// 	log.Println(err)
		// 	continue
		// }

		// sub, _ := nc.SubscribeSync("updates")
		// msg, _ := sub.NextMsg(10 * time.Millisecond)

		// // Use the response
		// log.Printf("Reply: %s", msg.Data)

		// // log.Println(msg)
		// // fmt.Println(msg)
		// // fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)
		// [begin subscribe_sync]
		// nc, err := nats.Connect(nats.DefaultURL)
		nc, err := nats.Connect("nats://nats.cloud-ml-mgmt:4222")
		if err != nil {
			log.Fatal(err)
		}
		defer nc.Close()

		// Use a WaitGroup to wait for a message to arrive
		wg := sync.WaitGroup{}
		wg.Add(1)

		// Subscribe
		if _, err := nc.Subscribe("notebook", func(m *nats.Msg) {
			log.Printf("Reply: %s", m.Data)
			// wg.Done()
		}); err != nil {
			log.Fatal(err)
		}

		// Wait for a message to come in
		wg.Wait()

		// [end subscribe_async]

	}

	// subj, i := args[0], 0

	// nc.Subscribe(subj, func(msg *nats.Msg) {
	// 	i += 1
	// 	printMsg(msg, i)
	// })
	// nc.Flush()

	// if err := nc.LastError(); err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("Listening on [%s]", subj)

}

// func setupConnOptions(opts []nats.Option) []nats.Option {
// 	totalWait := 10 * time.Minute
// 	reconnectDelay := time.Second

// 	opts = append(opts, nats.ReconnectWait(reconnectDelay))
// 	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
// 	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
// 		log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
// 	}))
// 	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
// 		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
// 	}))
// 	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
// 		log.Fatalf("Exiting: %v", nc.LastError())
// 	}))
// 	return opts
// }
