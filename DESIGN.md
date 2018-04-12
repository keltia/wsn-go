# Design of the WS-N protocol

## Description

WS-N is the OASIS way to implement the [Publish-Subscribe](https://en.wikipedia.org/wiki/Publish%E2%80%93subscribe_pattern) pattern.

## Mode of operations

It can operate in two modes: Push and Pull.  

The former is analogous of the [FTP protocol](https://en.wikipedia.org/wiki/File_Transfer_Protocol) in active mode where the server connects *back* to the client to send the data.  Here the WS-N client subscribe to a given topic giving an URI as a callback.  The server will then periodically do an [HTTP POST](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods) call for each packet of data.

In the latter mode, the client will ask for a PullPoint creation and then regularly poll it for data.  The advantage is that it can pass firewalls more easily, not having IP/hostname embedded into the Subscribe call.  This is again analogous to FTP in passive mode (without the embedded IP).

### Push mode

Here is a schema describing the operation:

    —> subscribe to topic with endpoint as argument
    <— get the de-subscription address
    <— data collection through periodic POST by the server to the endpoint 
    …
    —> unsubscribe from topic

### Pull mode

Here is a schema describing the operation:

    —> create PullPoint
    <— get the PullPoint URI as answer
    —> subscribe a topic to PullPoint
    loop
        —> read 1 or more messages, buffering done by the broker
        <— get packets
    —> unsubscribe topic from PullPoint
    —> destruction of the PullPoint

## Differences

As you can see, overhead is higher in Pull mode: on the one hand, you have to manage the PullPoint but on the other hand, you do not need an HTTP server to run (along with firewall rules and authorizations associated with inbound connections and possibly NAT handling).

## Implementation

For the moment the [wsn-go](https://github.com/keltia/wsn-go) package only implements the Push model in APIv2.

    client := wsn.NewClient(config)     // new client instance
    client.Subscribe(topic, callback)   // subscribe to a topic with that callback
    client.AddHandler(func)             // if you want to write into a file for example
    client.StartServer()                // start listening for POST for each callback
    for {
        server POST to http://ENDPOINT/callback -> data
    }
    client.Unsubscribe(topic)

## Pull mode

Wondering how to implement the Pull model, I came with two different ways, not sure which one would be more _idiomatic_ Go:

1. split `Client` type into `PushClient` & `PullClient` types with some code duplication between the two


    client := wsn.NewPushClient(config)     // new client instance
    client.Subscribe(topic, callback)       // subscribe to a topic with that callback
    client.AddHandler(func)                 // if you want to write into a file for example
    client.StartServer()                    // start listening for POST for each callback
    client.Unsubscribe(topic)

    and

    client := wsn.NewPullClient(config)     // new pull client
    pull := client.NewPull()                // create pull point
    pull.Subscribe(topic)                   // subscribe topic to pull point
    for {
        data := pull.GetMessage()           // get data
    }
    pull.Unsubscribe(topic)                 // finally ub-subscribe
    pull.Destroy()                          // cleanup

2. keep a common `Client` type with `NewClient()` taking a type parameter to create either mode.


    client := wsn.NewClient(PULL_MODE | PUSH_MODE)
    --> fill in a ClientOps type
    client.Subscribe(topic)                 // where is the callback? stored in the client context I guess
    client.Start()                          // start the HTTP server in push mode
    <loop ?>                                // easy in pull, use channel in push? io.Reader?
    client.Unsubscribe(topic)               // 
    client.Stop()                           // destroy the pull point if needed

I've gone over the 1st side for now.
