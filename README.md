# Mqueue (Go)

In-memory message broker in Go over an HTTP API using 
concurrent, thread-safe publisher/subscriber architecture
with multiple topics. 

## Subscribe
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

## Publish 
```sh 
curl -d "Any message you wish to post""http://localhost:8081/pub?topic=topic_of_your_choice"
```

## Disclaimer

Purely experimental project. Designed for learning purposes not production use.

## Contributing

Bug reports and pull requests are welcome on GitHub at [@ridwanmsharif](https://www.github.com/ridwanmsharif)

## Author

Ridwan M. Sharif:[ E-mail](ridwanmsharif@hotmail.com), [@ridwanmsharif](https://www.github.com/ridwanmsharif)

## License

The command line utility is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT)

