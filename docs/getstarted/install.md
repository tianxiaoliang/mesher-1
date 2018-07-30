# Quick start

### Local

1. Install ServiceComb [service-center](https://github.com/ServiceComb/service-center/releases)

2. Install go-chassis and run [rest server](https://github.com/go-chassis/go-chassis/tree/master/examples/rest/server)

3. Build mesher.go

4. Start mesher ./mesher
 
5. verify, in this case curl command is the consumer, mesher is consumer's sidecar, and rest server is provider

```shell
export http_proxy=http:127.0.0.1:30101
curl http://RESTServer/sayhello/peter
```

**Notice**:
>>You don't need to set service registry in chassis.yaml, because by default registry address is 127.0.0.1:30100, 
just same service center default listen address.
### Run on different infrastructure

Mesher does not bind to any platform or infrastructures, plz refer to 
https://github.com/go-chassis/mesher-examples/tree/master/Infrastructure
to know how to run mesher on different infra
