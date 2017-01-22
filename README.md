# mqueue (Go)

In-memory message broker in Go over an HTTP API using a concurrent, thread-safe publisher/subscriber architecture with multiple topics. 

## Installation

```sh
$ go get github.com/ridwanmsharif/mqueue
$ cd $GOPATH/src/github.com/ridwanmsharif/mqueue/
$ go build -o mqueue
$ ./mqueue
```
Listening on `:8081`, ready for client/terminal use.

## Examples/Usage

**Subscriber (client library)**
```go

  ch, err := client.Subscribe("topic_of_your_choice")
  if err != nil {
    log.Println(err)
    return
  }

  for e := range ch {
    log.Println(string(e))
  }

  log.Println("Channel closed")

```

**Publisher (client library)**
```go

  err := client.Publish("topic_of_your_choice", []byte("Arbitrary Message"))
  if err != nil {
      log.Println(err)
  }

```

**Subscriber (command line)**
```sh
curl -k -i -N -H "Connection: Upgrade" \
    -H "Upgrade: websocket" \
    -H "Host: localhost:8081" \
    -H "Origin:http://localhost:8081" \
    -H "Sec-Websocket-Version: 13" \
    -H "Sec-Websocket-Key: MQ" \
    "https://localhost:8081/sub?topic=
    topic_of_your_choice"
```

**Publisher (command line)** 
```sh 
curl -d "Any message you wish to post""http://localhost:8081/pub?topic=topic_of_your_choice"
```

## Disclaimer

Purely experimental project. Designed for learning purposes not production use. 
This weas my first project in Go, the implementation here is largely based on [asim/mq](https://www.github.com/asim/mq), I'm tracking his incremental progress and repeating the same

## Contributing

Bug reports and pull requests are welcome on GitHub at [@ridwanmsharif](https://www.github.com/ridwanmsharif)

## Author

Ridwan M. Sharif:[ E-mail](mailto:ridwanmsharif@hotmail.com), [@ridwanmsharif](https://www.github.com/ridwanmsharif)

## License

The command line utility is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT)

