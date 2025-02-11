# nem - nuzur entity manager

manages the system entities

## Getting Started

### Prerequisites

* Golang

### Running

```
CONFIG=./config/base.yaml go run ./api
```

#### Running Details

* The service provides a GRPC API on port 6005

* The service provides an AWS S3 upload API on port 8085
* The service publishes events via kafka - see the yaml file for the base configuration