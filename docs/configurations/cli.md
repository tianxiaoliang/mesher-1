# Mesher command Line 
when you start mesher process, you can use mesher command line to specify configurations like below
```shell
mesher --config=mesher.yaml --service-ports=rest:8080
```


### Options


**--config**
>*(optional, string)* the path to mesher configuration file, default value is {current_bin_work_dir}/conf/mesher.yaml


**--mode**
>*(optional, string)* mesher has 2 work mode, sidecar and per-host, default is sidecar


**--service-ports**
>*(optional, string)* running as sidecar, mesher need to know local service ports, 
this is to tell mesher service port list, 
if service has multiple protocol, you can write like rest:8080,grpc:9000. 
default is empty, mesher can not call your local service