### Gofc - [FullContact](http://fullcontact.com) Package for Go

#### TODO

* [x] Authentication
* [x] Person Lookup
    * [x] By Email
    * [x] By Phone
    * [x] By Twitter
    * [x] By Facebook
* [ ] Docs
* [ ] Tests

#### Package Usage
If you're embedding the Gofc package in your own application.

```go
import "github.com/140proof/gofc"

fc := gofc.NewClient("API_KEY_GOES_HERE")

res, err := fc.Person().GetByEmail("joe@example.com")
```

#### CLI Usage
If you'd like to use the Gofc CLI tool

```
sh> go build cmd/gofc ./...
sh> ./gofc -h
sh> ./gofc -method "twitter" -value "someone" -apikey "API_KEY"
```
