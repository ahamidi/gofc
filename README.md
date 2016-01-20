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

#### Usage

```go
import "github.com/140proof/gofc"

fc := gofc.NewClient("API_KEY_GOES_HERE")

res, err := fc.Person().GetByEmail("joe@example.com")
```

