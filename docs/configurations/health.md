# Local Health check
you can use health checker to check local service health,
when service instance is not healthy, mesher will update the instance status in registry service to "DOWN" 
so that other service
can not discover this instance. If the service is healthy again, mesher will update status to "UP", 
then other instance can discover it again. 
currently this function works only when you use service center as registry

examples:

Check local http service
```yaml
localHealthCheck:
  - portName: rest
    uri: /health
    interval: 30s
    match:
      status: 200
      body: ok
```

### Options


**portName**
>*(require, string)* must be one of the port name you define in mesher command line params "service-ports"
that name tells mesher the service port, currently just support rest-{suffix}, for other protocol, 
will use default TCP checker


**uri**
>*(optional, string)* uri start with /.


**interval**
>*(optional, string)* check interval, you can use number with unit: 1m, 10s. 

**match.status**
>*(optional, string)* the http response status must match status code

**match.body**
>*(optional, string)* the http response body must match body