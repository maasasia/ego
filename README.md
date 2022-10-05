# ego

Package for error wrapping, unwrapping, transforming to gql or http status code. (Inspired by [pkg/errors](https://github.com/pkg/errors))

`go get github.com/maasasia/ego`

## Adding context to an error

```go
_, err := ioutil.ReadAll(r)
if err != nil {
        return ego.Wrap(err, "read failed")
}
if err != nil {
        extra := 123
        return ego.Wrapf(err, "read failed with some code %d", extra)
}
```

## Retrieving the cause of an error

```go
func ErrorFunction() error {
	return ego.NewBadInputError("not found")
}

func Call1() error {
	err := ErrorFunction()
	if err != nil {
		return ego.Wrap(err, "FindById fail")
	}
	return nil
}

func Call2() error {
	err := Call1()
	if err != nil {
		return ego.Wrap(err, "Call1 fail")
	}
	return nil
}

func main() {
	err := Call2()
	if err != nil {
		rootErr := ego.Cause(err)
		switch rootErr.(type) {
		case ego.InternalServerError:
			fmt.Println("internal")
		case ego.ForbiddenError:
			fmt.Println("forbidden")
		case ego.BadInputError:
			fmt.Println("bad")
		default:
			fmt.Println("default")
		}
	}
}

// -> print "bad"
```