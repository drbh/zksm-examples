# zksm-examples

This library explores examples of Zero Knowledge Set Membership implemented by ing-banking. While they have some other eamples for the ZKRP they do not have much documentation around ZKSM. 

Here we make use of a forked repo of ing-banking's `zkproofs` repo, where we make the internal varibales exportable - and useable for building small apps. This way we can expose a REST server where a user can interact with `zkproofs` with simple JSON payloads.

We also add a small Python client that allows us to programatically make calls to the REST service and simulate the proces of two clients seting up as ZKSM, creating a proof and then verifying that proof.  

## Get Examples

```bash
git clone https://github.com/drbh/zksm-examples.git
```

## Script Descriptions

**basic/main.go**  

This script just runs the test contained in `zkproofs` testing methods. Here we expose it so it can be easily edited.  

**server/base.go**  

This script contains a self validating set of endpoints. It only takes `GET` requests and has hardcoded values that always resolve as `True`  

**server/main.go**  
This script contains a server that recives I/O as JSON formatted `POST` requests. The following demo Python files use this REST server to setup, prove, and verify.

**demo/one-user.py**  
This script contains makes a small set and tests if a value is in the set. This interacts only with one instance of the zkproofs server - and simulates a single user using the ZKSM REST service

**demo/two-users.py**  
This script contains a demo that connects to two instances of the `zkproofs` REST service. This simulates two users running the proccesses on their own machines. The demo sets up the set on `server A` and then `server B` uses the public data to generate a proof. Then `server A` verfies the proof.  


## Run Demo

Your going to need to open 3 terminal windows to run the processes.  

Terminal A.
```bash
cd zksm-examples  
go run server/main.go 8080  
```

Terminal B. (or other computer)
```bash
cd zksm-examples  
go run server/main.go 8081  
```

```bash
cd zksm-examples  
python demo/two-users.py   
```

output
```json
{ 
    "message": true
}  
```