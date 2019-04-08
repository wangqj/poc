# Rokie
a service for distributed key value management

## Conceptions

### Domain
same conception as tenant. 
all of resource is isolated by domain, it is optional if you don't have multi tenant requirement.
before you ask a kv value you can pull from other domain by specifying domain in request header 

### Realm
Realm is a unique name in one domain
it presents a config server endpoint, 
the config server could be a consul, spring cloud config server crip apollo cluster 
or even other rokie cluster.

