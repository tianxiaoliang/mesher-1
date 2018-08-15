# Minimum Configurations

To run mesher along with your service(a provider), you need to set minimum configurations as below:


1. Give mesher your service name in microservice.yaml file 
2. Set service discovery service(service center, Istio etc) configurations in chassis.yaml
3. export HTTP_PROXY=http://127.0.0.1:30101 as your service runtime environment
4. (optinal)Give mesher your service port list by ENV SERVICE_PORTS or CLI --service-ports

**Notice**:
 >> consumer need to use http://provider_name:provider_port/ to access provider,
 instead of http://provider_ip:provider_port/. 
 if you choose to set step4, then you can simply use http://provider_name/ to access provider