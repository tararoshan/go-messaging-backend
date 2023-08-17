# Go Messaging Service
Test it out for yourself using [http://gomessagingbackend-dev.us-east-2.elasticbeanstalk.com/](http://gomessagingbackend-dev.us-east-2.elasticbeanstalk.com/)! *Note: this is no longer up because I was spending too many free credits! You can still run this locally*

Based on [this task](https://gist.github.com/zackbloom/57124a029f6bd1b8ab0e3ea5aff34d71). Currently working on adding Cloudflare!

Time estimate: 16hr

Time spent: 17hr

## Getting Started
> **Note**
> You shouldn't copy the `$` symbol below. The `$` is just a reminder to type the commands into the terminal. Eg. if I write `$ echo 'hi!'` you should *actually* paste `echo 'hi!'` in the terminal.
Clone this repo and, from the terminal of the cloned reposity, run
```bash
$ go mod init mux
$ go mod tidy
```
You might need to download a few packages, like `gorilla/mux`. After you've done this once, to run the program, run
`$ go install *.go && go run main.go messagemap.go`. To make a GET request, try
```bash
$ curl http://localhost:3333/zack/charles/0
```
For a post request, try
```bash
$ curl -X POST -H 'Content-Type: application/json' -d '{"from": "zack", "to": "charles", "message": "pizza tonight?"}' http://localhost:3333
```

## Explanation/Blueprint
Here's my idea (written using Java names, but you get the gist):
- create a HashMap where the keys are alphabetically sorted (personA, personB) pairs, where personA and personB make up the sender and reciever of the given message.
- the values of the HashMap are ArrayLists of (timestamp, message) pairs. Always add to the back for time efficiency and to ensure the messages are in time-sorted order.
    - Edit: since I need to know who was the sender/reciever, store this in the array as well.
- when we have a GET request, look for the (personA, personB) key in the HashMap and then using that key's ArrayList arrays, binary search through by timestamp to find the starting index from which we should start printing messages!

5/14/23 Update:
Considering using [Redis](https://redis.io/docs/about/) to take care of concurrency issues. I'll stick with using a global HashMap and lock ([reader and writer](https://en.wikipedia.org/wiki/Readers%E2%80%93writer_lock)) instead, though. I want to stay simple for now.

5/16/23 Update:
It's not simple at all, but I still want to try.

5/21/23 Update:
Turns out there's a [RWMutex](https://pkg.go.dev/sync#RWMutex) in Go, so it might be that simple after all.

## Reading
Here are the resources I used to build this. This was my first time writing anything in Go, so I think this section should be beginner-friendly, assuming you have some programming background!

- [Go by Example](https://gobyexample.com/)
- [DigitalOcean's "HTTP Server in Go" Article](https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go)
- [http.ServeMux documentation](https://pkg.go.dev/net/http#ServeMux)
- [gorilla/mux GitHub repo](https://github.com/gorilla/mux) (not actively maintained as of Dec. 2022)
- [Alex Edwards's blogpost about Go routers](https://www.alexedwards.net/blog/which-go-router-should-i-use)
- [Go json package decode(v) documentation](https://pkg.go.dev/encoding/json#Decoder.Decode)
- [Alex Edwards's blogpost about parsing a Go request body](https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body)
- [JSON and Go](https://go.dev/blog/json)
- [sort package](https://pkg.go.dev/sort)
- [5 advanced testing techniques in Go](https://segment.com/blog/5-advanced-testing-techniques-in-go/)
- [starter CI/CD workflows for GitHub](https://github.com/actions/starter-workflows/tree/main/ci)
- [Understanding HTTP Requests & Responses in Golang](https://ciaranmcveigh5.medium.com/understanding-http-requests-responses-in-golang-a13e5e92bc4f)
- [https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/go-environment.html](Using the Elastic Beanstalk Go platform)
- [YouTube: Deploying A Go App to AWS Elastic Beanstalk](https://www.youtube.com/watch?v=1WXJTlkf0S4)
