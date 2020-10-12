# radius-auth-exporter

With this exporter you can check if your radius server correctly authenticates a given user and returns the correct attributes.

## Configuration and usage
Use the example config file and adapt to your needs, then run the tool with the path to your config:
```
./radius-auth-export -config /path/to/config.yaml
```

If everything went well, you should see your metrics at [localhost:2112/metrics](http://localhost:2112/metrics)

## Example metrics
```
# HELP radius_authentication_duration Amount of time to authetnicate against radius server
# TYPE radius_authentication_duration gauge
radius_authentication_duration{username="bob"} 0.000646936
radius_authentication_duration{username="unknown"} 1.000541992
radius_authentication_duration{username="user1"} 0.000477272
# HELP radius_authentication_response_attributes List of attributes sent in the reponse
# TYPE radius_authentication_response_attributes gauge
radius_authentication_response_attributes{attribute="Tunnel-Private-Group-Id",username="bob"} 0
radius_authentication_response_attributes{attribute="Tunnel-Private-Group-Id",username="user1"} 2
# HELP radius_authentication_success Indicates a successful authentication against the radius server
# TYPE radius_authentication_success gauge
radius_authentication_success{username="bob"} 1
radius_authentication_success{username="unknown"} 0
radius_authentication_success{username="user1"} 1

```
