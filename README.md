# goforward
> forward tcp ipv4/6:port to ipv4/6:port

```
Usage:
  goforward [flags]

Flags:
  -c, --config string         config file of frps
  -f, --forward_link string   forward link, e.g. -f=127.0.0.1:12345->[2400:3200::1]:443,127.0.0.1:12346->[2400:3200:baba::1]:443
  -h, --help                  help for goforward
  -l, --listen_addr string    forward api listen addr (default "0.0.0.0:5668")
```

## Examples
- forward local ipv4 to remote ipv4
```
goforward -f 127.0.0.1:1111->1.1.1.1:443
```
- forward local ipv4 to remote ipv6
```
goforward -f 127.0.0.1:12345->[2400:3200::1]:443
```
- forward multiple mixes
```
goforward -f 127.0.0.1:12345->[2400:3200::1]:443,127.0.0.1:12346->[2400:3200:baba::1]:443
```
- forward with config
```
goforward -c config.yaml
```
`config.yaml`
```yaml
listen_addr: 127.0.0.1:5668
forward_links: 
  - 127.0.0.1:12345->[2400:3200::1]:443
  - 127.0.0.1:12346->[2400:3200:baba::1]:443
```

## API
> See Makefile for details

- get all forward
- get specified forward
- add forward
  - support override
- del forward


