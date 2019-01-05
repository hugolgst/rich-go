This is forked from the original working version at [github.com/ananagame/rich-go](github.com/ananagame/rich-go).
The only difference is the removal of `fmt.Println` output


# rich-go

A simple Golang extension for Discord Rich Presence

## Installation

Install `github.com/ananagame/rich-go`:

```
$ go get github.com/ananagame/rich-go
```

## Usage

Login by sending the first handshake :
```crystal
client.Login("your_client_id")
```

And you can set the Rich Presence activity (parameters can be found :
```crystal
client.SetActivity(client.Activity{
		State:   "hey",
		Details: "i'm running on go",
	})
```

## Contributing

1. Fork it (https://github.com/ananagame/rich-go/fork)
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request

## Contributors

- [AnanaGame](https://github.com/ananagame) - creator, maintainer
