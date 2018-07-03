# Admin API

### Configurations

Listen on isolated port, it gives a way to interact with mesher

```yaml
admin: 
  enable: true
  serverUri : 127.0.0.1:30102 # addr on listening
  goRuntimeMetrics : true # enable metrics
```

**enable**
>*(optional, bool)* default is false

**serverUri**
>*(optional, string)* Listen address,default is 0.0.0.0:30102

**goRuntimeMetrics**
>*(optional, bool)* default is false, enable to expose go runtime metrics in /v1/mesher/metrics

