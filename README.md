# Hello World Go/Beanstalk

Very simple Go/AWS Elastic Beanstalk application, based on `geetarista/go-http-hello-world`. Project used to create one Beanstalk app, `hello-go`, and two Beanstalk environments, `hello-go-v1` and `hello-go-v2`, using Elastic Beanstalk CLI (`eb`).

## Build and Run Locally

```bash
go build hello_world.go
./hello-world
curl http://localhost:3000
```

## Build and Run on AWS

```bash
eb --help

eb init hello-go

eb create hello-go-v1
eb deploy hello-go-v1

eb create hello-go-v2
eb deploy hello-go-v2

eb scale 2 hello-go-v1
eb scale 2 hello-go-v2

eb open hello-go-v1

# swap source dns with destination
eb swap hello-go-v1 -n hello-go-v2

eb config save hello-go-v1 --cfg hello-go-v1-sc

eb terminate hello-go-v1 --force --verbose
eb terminate hello-go-v2 --force --verbose
```
