### Test task 

Create a console application in Go,
which will receive a list of cat breeds via the API (https://catfact.ninja/breeds).
Group all breed names by origin in the country,
sort the breed name by length and write the resulting data to the JSON file out.json.
Pay special attention to code quality and error handling.

--------------------------------------------------------------------------------------

### How to run app

#### Use makefile: 

- For build program run command `make build`
- For program execution run command `make run`
- For clear cache and remove binary file use command `make clear` 

#### Build and run manually:

1. Build
`$ go build -o ./bin/app ./cmd`  

2. Run
`$ ./bin/app`

#### Run test:

- `$ go test -v ./...`
