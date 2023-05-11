# Go Messaging Service
Based on [this task](https://gist.github.com/zackbloom/57124a029f6bd1b8ab0e3ea5aff34d71). Currently
in-progress!

## Getting Started
Clone this repo and run `go run main.go` in the cloned directory. To make a GET request, try

```bash
curl http://localhost:3333
```

For a post request, try

```bash
curl -X POST -H 'Content-Type: application/json' -d '{"from": "zack", "to": "charles", "message": "pizza tonight?"}' http://localhost:3333
```

## Explanation/Blueprint
Here's my idea (written using Java names, but you get the gist):
- create a HashMap where the keys are alphabetically sorted (personA, personB) pairs, where personA and personB make up the sender and reciever of the given message.
- the values of the HashMap are ArrayLists of (timestamp, message) pairs. Always add to the back for time efficiency and to ensure the messages are in time-sorted order.
    - Edit: since I need to know who was the sender/reciever, store this in the array as well.
- when we have a GET request, look for the (personA, personB) key in the HashMap and then using that key's ArrayList arrays, binary search through by timestamp to find the starting index from which we should start printing messages!

## Reading
Here are the resources I used to build this. This was my first time writing anything in Go, so I
think this section should be beginner-friendly, assuming you have some programming background!

- [Go by Example](https://gobyexample.com/)
- [HTTP Server in Go](https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go)
- [http.ServeMux documentation](https://pkg.go.dev/net/http#ServeMux)
- [gorilla/mux](https://github.com/gorilla/mux) (not actively maintained as of Dec. 2022)
- [Blogpost about Go routers](https://www.alexedwards.net/blog/which-go-router-should-i-use)

## Backburner
1. How to make sure a POST request to root is denied?
2. Unit testing!
