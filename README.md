# Proof of Work

Simple implementation of proof-of-work algorithm to limit request per second. To get a piece of wisdom (make server do some work) the client app solving task based on data recieved from server. 

[![Go Version](https://img.shields.io/badge/Go-v1.19-blue)](https://golang.org/dl/)
[![Go Report Card](https://goreportcard.com/badge/github.com/kormiltsev/proofofwork)](https://goreportcard.com/report/github.com/kormiltsev/proofofwork)
[![License](https://img.shields.io/github/license/kormiltsev/proofofwork)](https://github.com/kormiltsev/proofofwork/blob/main/LICENSE)

Final result is random quote like this:

```
If A equals success, then the formula is A = X + Y + Z. X is work. Y is play. Z is keep your mouth shut.
		-- Albert Einstein
```

## Installation

Use Make to build an app binary. The same file starts both server and a client. 

```bash
make build
```

Server and a client can run in separate containers. There are 2 Dockerfiles. 

Dockerfile is for server app. Example: 

```bash
docker build -t words-of-wisdom .   
docker run -d --rm -p 12000:8080 --name w-o-w words-of-wisdom
```

Dockerfile.client is for client app. Example: 

```bash
docker build -t words-of-wisdom-client -f Dockerfile.client . 
docker run -it --rm --net=host --name w-o-w-c words-of-wisdom-client
```

## Usage

### Binary usage

Server

```bash
words-of-wisdom-amd64 run
```

```bash
words-of-wisdom-amd64 run --socket=:8080 --initial-difficulty=16 --cache-limit=1000
```

Client

```bash
words-of-wisdom-amd64 client
```

```bash
words-of-wisdom-amd64 client --url=http://localhost:8080/words --endless
```

### Container usage

Client operates in endless mode byt default, sending requests right after response.

## Result

```
🟢 0000e539e597ef41604cb59f5f276f16a9e0375720c6d45fadbccf13de2ea899
result: ========================================================
When the Universe was not so out of whack as it is today, and all the
stars were lined up in their proper places, you could easily count them
from left to right, or top to bottom, and the larger and bluer ones
were set apart, and the smaller yellowing types pushed off to the
corners as bodies of a lower grade ...
		-- Stanislaw Lem, "Cyberiad"
```

## Roadmap

- dinamic difficulty based on request quantity
- tests
- one logger to log them all
- improve memory usage by replacing strings usage etc.

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
