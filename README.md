# Micro Services for fun


based on Writing microservices with Go Micro https://medium.com/microhq/writing-microservices-with-go-micro-b6e0880b8829

Imports

go get -v import github.com/micro/go-micro
go get github.com/micro/protobuf/{proto,protoc-gen-go}

Source 

main.go
proto/greeter.proto    
protoc --go_out=plugins=micro:. greeter.proto

Service Discovery

The service discovery mechanism will need to be running so the service can register to be discovered by clients and other services. https://www.consul.io/

brew install -v consul 
brew services start consul
http://127.0.0.1:8500/ui/dc1/services

### Imports

    go get -v import github.com/micro/go-micro
    go get github.com/micro/protobuf/{proto,protoc-gen-go}

### Generate the proto sources 

    protoc --go_out=plugins=micro:. greeter.proto

### Service Discovery

The service discovery mechanism will need to be running so the service can register to be discovered by clients and other services. 

-> https://www.consul.io/


    brew install -v consul 

    brew services start consul

    http://127.0.0.1:8500/ui/dc1/services
