## Zeus Client ##

The Zeus client is used for interacting with the cloud application. You can override the values in the test to your own 
following these steps

1. In /test/configs -> create a config.yaml using the sample-config.yaml as a reference, the config.yaml should be in .gitignore by default so it doesn't commit your tokens
2. Add your bearer token to this config, otherwise config it directly in the client
3. Override the zeus_client_test variables that are used to point to your desired chart and kubernetes location
4. Then run the test

The 


