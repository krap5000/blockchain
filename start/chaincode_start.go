/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ===============================================================================================
// Main
// ==================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
//=================================================================================================================


//=========================================================================================================================
// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")

		  err := stub.PutState("Project_Configuration_Name", []byte(args[0]))
          //err := stub.PutState("Project_Configuration_Owner", []byte(args[1]))
          //err := stub.PutState("Project_Configuration_Description", []byte(args[2))
          
        
        
        
    	  if err != nil {
        	return nil, err
    		}
	}

	return nil, nil
}
//====================================================================================================



//====================================================================================================
// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	  fmt.Println("invoke is running " + function)

    // Handle different functions
    if function == "init" {
        return t.Init(stub, "init", args)
    } else if function == "write" {
        return t.write(stub, args)
    }
    fmt.Println("invoke did not find func: " + function)

    return nil, errors.New("Received unknown function invocation")
}
//==================================================================================================================




//==================================================================================================================
// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	    fmt.Println("query is running " + function)

    // Handle different functions
    if function == "read" {                            //read a variable
        return t.read(stub, args)
    }
    fmt.Println("query did not find func: " + function)

    return nil, errors.New("Received unknown function query")
}
//=======================================================================================================================




//=================================================================================================================
//Write Function

func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    
	var Project_Configuration_Name, pcn_v string
	//Modification, Project_Configuration_Modify_Date, Project_Configuration_Modifier_Name, pcd_v, pcmd_v, pmmn_v string
    var err error
    
    //Project_Configuration_Modification, Project_Configuration_Modify_Date, Project_Configuration_Modifier_Name
    fmt.Println("running write()")

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting 1. name of the variable and value to set")
    }
    Project_Configuration_Name = args[0];
     pcn_v = args[1];
//    Project_Configuration_Modify_Date = args[2];
//    pcmd_v = args[3];
//    Project_Configuration_Modifier_Name = args[4];
//    pcmn_v = args[5];
    
       
    
    err = stub.PutState(Project_Configuration_Name, []byte(pcn_v))  //write the variable into the chaincode state
    //err = stub.PutState(Project_Configuration_Modify_Date, []byte(pcmd_v))  //write the variable into the chaincode state
    //err = stub.PutState(Project_Configuration_Modifier_Name, []byte(pcmn_v))  //write the variable into the chaincode state
  
        
    
    if err != nil {
        return nil, err
    }
    return nil, nil
}
//===================================================================================================



    
//==============================================================================================                                                                     
//Read Function

    func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    
		
	
	var Project_Configuration_Name string
	//var Project_Configuration_Owner, Project_Configuration_Description string
	//var Project_Configuration_Modification, Project_Configuration_Modify_Date, Project_Configuration_Modifier_Name string
	//Project_Configuration_Name, Project_Configuration_Owner, Project_Configuration_Description string
		
		
	var jsonResp string
    var err error
	
	
	
/*	type Message struct {
    Name string
    Body string
    Time int64
}
and an instance of Message

m := Message{"Alice", "Hello", 1294706395881547000}
we can marshal a JSON-encoded version of m using json.Marshal:

b, err := json.Marshal(m)
If all is well, err will be nil and b will be a []byte containing this JSON data:

b == []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)

   */
		
		
		if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
    }

    Project_Configuration_Name = args[0]
	
		
		
		
		
    valAsbytes, err := stub.GetState(Project_Configuration_Name)  //returns a json response. valsbytes is json.
		
		
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + Project_Configuration_Name + "\"}"
        return nil, errors.New(jsonResp)
    }

    return valAsbytes, nil
}

//=====================================================================================================

