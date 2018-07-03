# Minimum Configurations

To run mesher along with your service, you need to set minimum configurations as below:

1. Give mesher your service port list by ENV SERVICE_PORTS or CLI --service-ports
2. Give mesher your service name in microservice.yaml file 
3. Set service discovery service(service center, Istio etc) configurations in chassis.yaml
4. export HTTP_PROXY=http://127.0.0.1:30101 as your service runtime environment

**Notice**:
 >> consumer need to use http://provider_service_name/ to access provider,instead of http://ip:port/