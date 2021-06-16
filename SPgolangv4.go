/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	//"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"

	// cid library to be used for getting certificate
	// "github.com/hyperledger/fabric/core/chaincode/lib/cid"

)

//  var getCreator = func(certificate []byte) (string, string) {
//	data := certificate[strings.Index(string(certificate), "-----") : strings.LastIndex(string(certificate), "-----")+5]
//	block, _ := pem.Decode([]byte(data))
//	cert, _ := x509.ParseCertificate(block.Bytes)
//	organization := cert.Issuer.Organization[0]
//	commonName := cert.Subject.CommonName
//	organizationShort := strings.Split(organization, ".")[0]
//	return commonName, organizationShort
//}
//SmartContract is the definition of the chaincode structure.
//Smartcontract for Hyperledger Fabric

type SmartContract struct {
}

// Indexes store asset indexes in the blockchain
type Indexes struct {

	//indexTRU is the last index of the TRU
	IndexTRU string `json:"indexTRU"`
}

// TRU Define the standard part.  Structure tags are used by encoding/json library
type TRU struct {

	//TRUID is the unique id of a TRU in the blockchain records used by ledger
	TRUID string `json:"TRUID"`

	// GTIN or PO is used as identifier for the part instance, is the product number
	GTIN string `json:"GTIN"`

	//Quantity is used to indicate the quantity of TRU
	Quantity string `json:"Quantity"`

	//Batchnumber refers to the production batch
	BatchNumber string `json:"BatchNumber"`

	//LocalID is the local id of the TRU
	LocalID string `json:"LocalID"`
	
	//ExpiryDate is used tp store expiry date of particular TRU
	ExpiryDate string `json:"TRUExpiryDate"`

	//CoCID refers to CoC Spec of the part. COC stored as string in the first version.
	CoC string `json:"CoCID"`

	//Tracehash is used tp verify the trace of the part
	TraceHash string `json:"TraceHash"`

	//CoChash is used to verify the CoC
	CoCHash string `json:"CoChash"`

	//ReleaseFlag is used to verify the requestflag
	ReleaseFlag string `json:"ReleaseFlag"`

	//RequestFlag is used to verify the requestflag
	RequestFlag string `json:"RequestFlag"`
	
	//PaymentPeriod is to store the payment period for TRU
	PaymentPeriod string `json:"PaymentPeriod"`

}

// SCA is supply chain agent.  Structure tags are used by encoding/json library
type SCA struct {

	// Industry Specific ID OwnerGLN
	OrgID string `json:"OrgID"`

	//Type Manufacturer, Supplier, Orderer default = 111 assumes all three roles 111 all enabled
	SCAType string `json:"SCAType"`

	//Name of the Organization
	Name string `json:"Name"`

	// Supplier Certificates Names
	CertificateIDs string `json:"CertificateIDs"`
}

//CERT is certification log entries
type CERT struct {
	OrgID string `json:"CertownorgID"`

	CertificeteID string `json:"CertificateID"`

	IssuerPrefix string `json:"CertificateIssuer"`

	CertificateName string `json:"CertificateName"`

	//Producer, supplier or order certificate for a part
	Type string `json:"CertificateType"`

	ValidityDate string `json:"ValidityDate"`
}

//OWN represents Ownership relation
type OWN struct {
	OwnerOrgID string `json:"OwnerOrgID"`
	OwnerTRUID string `json:"OwnerTRUID"`

	//Change Date is recorded
	OwnDate string `json:"Date"`
}

//TRACE represents the trace data
type TRACE struct {

	//OwnerOrgID is recorded
	OwnerOrgID string `json:"TraceOwnerOrgID"`

	//PartGTINSEQ includes GTIN with extension of sequences of events of a particular GTIN
	TraceTRUID string `json:"TraceTRUID"`

	//TraceSequenceNo is updated for every transaction
	//TraceSequenceNo string `json:"TraceSequence"`

	//PreviousownerID is recorded
	//PrevOwnerOrgID string `json:"PrevOwnerOrgID"`

	//Change Date is recorded
	ChangeDate string `json:"Date"`

	//Change Time is recorded
	//ChangeTime string `json:"Time"`

	//EventType is either produce /1/, receive /2/,  split /3/, ship /4/, install /5/
	EventType string `json:"EventType"`

	//EventID is test report, GTIN, GTIN of splitted batch, shipment number/ packingID, intalled unit ID.
	EventID string `json:"EventID"`

	//Optional: The trace can be private data
}
//MSPID Struct to store calling party of the transaction
type MSPID struct {
	Mspid   string `json:"MSPID"`
	IDBytes byte
}

// Init method is called when the Smart Contract "StandardParts" is instantiated by the blockchain network
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke method is called as a result of an application request to run the Smart Contract "StandardParts"
// The calling application program has also specified the particular smart contract function to be called, with arguments
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	//mspid := "ORG1"
	//cid := new ClientIdentity(APIstub);
	//mspid, err := cid.GetMSPID(APIstub)
	//if err != nil {
	//	fmt.Println("Error - GetMSPID()")
	//}
	//fmt.Print("MSPID: ")
	//fmt.Print(mspid)
	//fmt.Println("\n\n -")

	// Get necessary creator info for calling the chaincode
	//id, err := cid.GetID(APIstub)
	//if err != nil {
	//	fmt.Println("Error - cide.GetID()")
	//}
	//fmt.Print("ID: ")
	//fmt.Print(id)
	//fmt.Println("\n\n -")

	//mspid, err := cid.GetMSPID(APIstub)
	//if err != nil {
	//	fmt.Println("Error - cide.GetMSPID()")
	//}
	//fmt.Print("MSPID: ")
	//fmt.Print(mspid)
	//fmt.Println("\n\n -")

	//cert, err := cid.GetX509Certificate(APIstub)
	//if err != nil {
	//	fmt.Println("Error - cide.GetX509Certificate()")
	//}
	//fmt.Print("GetX509Certificate: \n\n")
	//fmt.Print("IPAddresses: ")
	//fmt.Println(cert.IPAddresses)

	//fmt.Print("IsCA: ")
	//fmt.Println(cert.IsCA)

	//fmt.Print("Issuer: ")
	//fmt.Println(cert.Issuer)

	//fmt.Print("IssuingCertificateURL: ")
	//fmt.Println(cert.IssuingCertificateURL)

	//fmt.Print("Public Key: ")
	//fmt.Println(cert.PublicKey)

	//fmt.Println("\n\n -")

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryTRU" {
		return s.queryTRU(APIstub, args)
	} else if function == "querySCA" {
		return s.querySCA(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "updateTRU" {
		return s.updateTRU(APIstub, args)
	} else if function == "createTRU" {
		return s.createTRU(APIstub, args)
	} else if function == "splitTRU" {
		return s.splitTRU(APIstub, args)
	} else if function == "updateSCA" {
		return s.updateSCA(APIstub, args)
	} else if function == "queryAllSCA" {
		return s.queryAllSCA(APIstub)
	} else if function == "queryAllTRU" {
		return s.queryAllTRU(APIstub)
	} else if function == "queryTRACE" {
		return s.queryTRACE(APIstub, args)
	} else if function == "updateTRACE" {
		return s.updateTRACE(APIstub, args)
	} else if function == "changeOWN" {
		return s.changeOWN(APIstub, args)
	} else if function == "searchTRU" {
		return s.searchTRU(APIstub, args)	
	} else if function == "queryNextTRU" {
		return s.queryNextTRU(APIstub)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {

	TRUs := []TRU{

		//Only one TRU is include as an example
		TRU{TRUID: "TRU1", GTIN: "H0W98.1.1000001", Quantity: "100", BatchNumber: "1001", LocalID: "00001", ExpiryDate: "01-01-2025", CoC: "1001", TraceHash: "TH", CoCHash: "CH", ReleaseFlag: "0", RequestFlag: "0", PaymentPeriod: "30"},
		TRU{TRUID: "TRU2", GTIN: "H0W98.2.1000001", Quantity: "200", BatchNumber: "1001", LocalID: "00002", ExpiryDate: "01-01-2025", CoC: "1001", TraceHash: "TH", CoCHash: "CH", ReleaseFlag: "0", RequestFlag: "0", PaymentPeriod: "30"}, 
		TRU{TRUID: "TRU3", GTIN: "F0210.1.1000001", Quantity: "400", BatchNumber: "100001", LocalID: "00003", ExpiryDate: "01-01-2025", CoC: "1002", TraceHash: "TH", CoCHash: "CH", ReleaseFlag: "0", RequestFlag: "0", PaymentPeriod: "30"},
	}

	//To be added to read the organizations in the blockchain
	i := 1
	for i < len(TRUs)+1 {
		fmt.Println("i is ", i)
		TRUsAsBytes, _ := json.Marshal(TRUs[i-1])
		APIstub.PutState("TRU"+strconv.Itoa(i), TRUsAsBytes)
		fmt.Println("Added", TRUs[i-1])
		i = i + 1
	}

	SCAs := []SCA{
		SCA{OrgID: "ORG1", SCAType: "111", Name: "Fokker Services B.V.", CertificateIDs: "CAA NL: PART-21 Subpart G"},
		SCA{OrgID: "ORG2", SCAType: "111", Name: "Airbus Helicopters", CertificateIDs: "EASA: 21G.0070"},
		SCA{OrgID: "ORG3", SCAType: "011", Name: "Honeywell International, Inc. D/B/AGREER Repair & Overhaul", CertificateIDs: "EASA: 145.4123"},
		SCA{OrgID: "ORG4", SCAType: "011", Name: "Avmax Aviation Services Inc.", CertificateIDs: "EASA: 145.7015"},
		SCA{OrgID: "ORG5", SCAType: "001", Name: "KLM Engineering & Maintenance", CertificateIDs: "CAA NL: 145.1113"},
		SCA{OrgID: "ORG6", SCAType: "001", Name: "Fokker Techniek B.V.", CertificateIDs: "CAA NL: 145.1358"},
	}

	j := 1
	for j < len(SCAs)+1 {
		fmt.Println("j is ", j)
		SCAsAsBytes, _ := json.Marshal(SCAs[j-1])
		APIstub.PutState("ORG"+strconv.Itoa(j), SCAsAsBytes)
		fmt.Println("Added", SCAs[j-1])
		j = j + 1
	}

	OWNs := []OWN{
		OWN{OwnerOrgID: "ORG1", OwnerTRUID: "TRU1", OwnDate: "2020-01-01"},
		OWN{OwnerOrgID: "ORG2", OwnerTRUID: "TRU2", OwnDate: "2020-01-01"},
		OWN{OwnerOrgID: "ORG2", OwnerTRUID: "TRU3", OwnDate: "2020-01-01"},
	}

	//TESTING INITIALIZATION CODE
	orgs1 := []string{"ORG1", "ORG2", "ORG2"}
	trus1 := []string{"TRU1", "TRU2", "TRU3"}

	k := 1
	for k < len(OWNs)+1 {
		fmt.Println("k is ", k)
		oWNsAsBytes, _ := json.Marshal(OWNs[k-1])
		APIstub.PutState("OWN:"+orgs1[k-1]+":"+trus1[k-1]+strconv.Itoa(k), oWNsAsBytes)
		fmt.Println("Added", OWNs[k-1])
		k = k + 1
	}

	TRACEs := []TRACE{
		TRACE{OwnerOrgID: "ORG1", TraceTRUID: "TRU1", ChangeDate: "2019-01-01", EventType: "1", EventID: "TestReport1001"},
		TRACE{OwnerOrgID: "ORG2", TraceTRUID: "TRU2", ChangeDate: "2019-01-01", EventType: "1", EventID: "TestReport1002"},
		TRACE{OwnerOrgID: "ORG2", TraceTRUID: "TRU3", ChangeDate: "2019-01-01", EventType: "1", EventID: "TestReport1003"},
	}

	l := 1
	for l < len(TRACEs)+1 {
		fmt.Println("l is ", l)
		tRACEsAsBytes, _ := json.Marshal(TRACEs[l-1])
		APIstub.PutState("TRACE:"+orgs1[l-1]+":"+trus1[l-1]+":"+"1", tRACEsAsBytes)
		fmt.Println("Added", TRACEs[l-1])
		l = l + 1
	}

	//Update TRU index using initialized TRUs.
	indexes2 := Indexes{IndexTRU: strconv.Itoa(k)}
	IndexAsBytes, _ := json.Marshal(indexes2)
	APIstub.PutState("nextTRU", IndexAsBytes)

	return shim.Success(nil)

}

func (s *SmartContract) queryTRU(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	TRUAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(TRUAsBytes)
}

func (s *SmartContract) searchTRU(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) == 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	startKey := "TRU1"
	endKey := "TRU999999"

	resultsIterator, _ := APIstub.GetStateByRange(startKey, endKey)
	
	for resultsIterator.HasNext() {
		
		queryResponse, _ := resultsIterator.Next()
		tRU := TRU{}
		json.Unmarshal(queryResponse.Value, &tRU)

		
		if tRU.LocalID == args[0] {
			TRUAsBytes, _ := APIstub.GetState(tRU.TRUID)
			return shim.Success(TRUAsBytes)
		}
	}

	return shim.Error("No such TRU")
}


func (s *SmartContract) createTRU(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 12 {
		return shim.Error("Incorrect number of arguments. Expecting 12")
	}

	//Later to be obtained from certificate
	mspid := args[11]

	//Get last TRU and information from the ledger\

	indexes := Indexes{}

	indexTRUAsBytes, _ := APIstub.GetState("nextTRU")
	json.Unmarshal(indexTRUAsBytes, &indexes)
	indexTRU, _ := strconv.Atoi(indexes.IndexTRU)

	var tRU = TRU{TRUID: ("TRU" + (strconv.Itoa(indexTRU))), GTIN: args[0], Quantity: args[1], BatchNumber: args[2], LocalID: args[3], ExpiryDate: args[4], CoC: args[5], TraceHash: args[6], CoCHash: args[7], ReleaseFlag: args[8], RequestFlag: args[9], PaymentPeriod: args[10]}

	TRUAsBytes, _ := json.Marshal(tRU)
	APIstub.PutState("TRU"+strconv.Itoa(indexTRU), TRUAsBytes)

	//Get Current Date
	currentTime := time.Now()

	//Write the ownership info

	var oWN = OWN{OwnerOrgID: mspid, OwnerTRUID: "TRU" + strconv.Itoa(indexTRU), OwnDate: currentTime.Format("2006-01-02")}

	OWNAsBytes, _ := json.Marshal(oWN)
	APIstub.PutState("OWN:"+mspid+":"+"TRU"+strconv.Itoa(indexTRU), OWNAsBytes)
	fmt.Println("Added", oWN)

	//Write the Trace info

	var tRACE = TRACE{OwnerOrgID: mspid, TraceTRUID: "TRU" + strconv.Itoa(indexTRU), ChangeDate: currentTime.Format("2006-01-02"), EventType: "Produce", EventID: args[4]}

	TRACEAsBytes, _ := json.Marshal(tRACE)
	APIstub.PutState("TRACE:"+mspid+":"+"TRU"+strconv.Itoa(indexTRU)+":"+"1", TRACEAsBytes)
	fmt.Println("Added", tRACE)

	indexTRU = indexTRU + 1
	indexes2 := Indexes{IndexTRU: strconv.Itoa(indexTRU)}
	IndexAsBytes, _ := json.Marshal(indexes2)
	APIstub.PutState("nextTRU", IndexAsBytes)

	//TO DO LATER: Get CoC pdf. Check for properties (external functions)
	//TO DO LATER: Save CoC pdf as private data.

	return shim.Success(TRUAsBytes)
}

//UpdateTRU. To be updated to change any trribute of TRU

func (s *SmartContract) updateTRU(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	//To be read from certificate later
	mspid := args[3]

	tRUAsBytes, _ := APIstub.GetState(args[0])
	tRU := TRU{}
	tRACE := TRACE{}

	json.Unmarshal(tRUAsBytes, &tRU)

	if args[1] == "BatchNumber" {
		tRACE.EventType = "Batch"
		tRACE.EventID = tRU.BatchNumber
		tRU.BatchNumber = args[2]
	} else if args[1] == "LocalID" {
		tRACE.EventType = "LocalID"
		tRACE.EventID = tRU.LocalID
		tRU.LocalID = args[2]
	}else if args[1] == "CoC" {
		tRACE.EventType = "CoC"
		tRACE.EventID = tRU.CoC
		tRU.CoC = args[2]
	} else if args[1] == "ExpiryDate" {
		tRACE.EventType = "ExpiryDate"
		tRACE.EventID = tRU.ExpiryDate
		tRU.ExpiryDate = args[2]
	} else if args[1] == "TraceHash" {
		tRACE.EventType = "TraceHash"
		tRACE.EventID = tRU.TraceHash
		tRU.TraceHash = args[2]
	} else if args[1] == "CoCHash" {
		tRACE.EventType = "CoCHash"
		tRACE.EventID = tRU.CoCHash
		tRU.CoCHash = args[2]
	} else if args[1] == "RequestFlag" {
		tRACE.EventType = "Order"
		tRACE.EventID = args[2]
		tRU.RequestFlag = args[2]
	} else if args[1] == "ReleaseFlag" {
		tRACE.EventType = "Ship"
		tRACE.EventID = args[2]
		tRU.ReleaseFlag = args[2]
	} else if args[1] == "Quantity" {
		tRACE.EventType = "Quantity Update"
		tRACE.EventID = args[2]
		tRU.Quantity = args[2]
	} else if args[1] == "PaymentPeriod" {
		tRACE.EventType = "Payment Period Update"
		tRACE.EventID = args[2]
		tRU.PaymentPeriod = args[2]
	}

	TRUAsBytes, _ := json.Marshal(tRU)
	APIstub.PutState(args[0], TRUAsBytes)

	//TO BE CONSIDERED LATER: If GTIN is changed then possibly new TRU is created. Trace records for the new GTIN are copied and appended.

	//Get Current Date
	currentTime := time.Now()

	//Record Trace change. Get current trace  and copy to new one
	tRACE.OwnerOrgID = mspid
	tRACE.TraceTRUID = args[0]
	tRACE.ChangeDate = currentTime.Format("2006-01-02")

	startKey := "TRACE:" + mspid + ":" + args[0] + ":" + "1"
	endKey := "TRACE:" + mspid + ":" + args[0] + ":" + "999"
	n := 1

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		resultsIterator.Next()
		n = n + 1
	}

	TRACEAsBytes, _ := json.Marshal(tRACE)
	APIstub.PutState("TRACE:"+mspid+":"+args[0]+":"+strconv.Itoa(n), TRACEAsBytes)
	fmt.Println("Added", tRACE)

	return shim.Success(nil)
}

func (s *SmartContract) splitTRU(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	//Optional: Read certificate and compare
	mspid := args[3]

	var q1, q2 int

	// Check the similarity with Create TRU
	indexes := Indexes{}
	indexTRUAsBytes, _ := APIstub.GetState("nextTRU")
	json.Unmarshal(indexTRUAsBytes, &indexes)
	indexTRU, _ := strconv.Atoi(indexes.IndexTRU)

	//TO DO: Check submitting organization is the owner. Error: TRU does not belong to your organization.

	//Get TRU info to be split
	tRU := TRU{}
	tRUAsBytes, _ := APIstub.GetState(args[0])
	json.Unmarshal(tRUAsBytes, &tRU)

	//Find new quantities
	q1, _ = strconv.Atoi(tRU.Quantity)
	q2, _ = strconv.Atoi(args[1])

	//Check the reuired quantity is smaller than original.
	if q1 <= q2 {
		return shim.Error("The split can not be completed. Quantity of new TRU should be lower than the original TRU.")
	}

	tRU.Quantity = strconv.Itoa(q1 - q2)

	//Current Time
	currentTime := time.Now()

	//Write the old TRU infornmation
	TRUAsBytes, _ := json.Marshal(tRU)
	APIstub.PutState(args[0], TRUAsBytes)

	//Form the second TRU information and write it to second TRU
	tRU2 := TRU{TRUID: "TRU" + strconv.Itoa(indexTRU), GTIN: tRU.GTIN, Quantity: args[1], BatchNumber: tRU.BatchNumber, LocalID: args[2], ExpiryDate: tRU.ExpiryDate, CoC: tRU.CoC, TraceHash: tRU.TraceHash, CoCHash: tRU.CoCHash, ReleaseFlag: tRU.ReleaseFlag, RequestFlag: tRU.RequestFlag, PaymentPeriod: tRU.PaymentPeriod}

	TRUAsBytes2, _ := json.Marshal(tRU2)
	APIstub.PutState("TRU"+strconv.Itoa(indexTRU), TRUAsBytes2)

	//Write the ownership information for the new TRU
	oWN2 := OWN{OwnerOrgID: mspid, OwnerTRUID: "TRU" + strconv.Itoa(indexTRU), OwnDate: currentTime.Format("2006-01-02")}

	OWNAsBytes2, _ := json.Marshal(oWN2)
	APIstub.PutState("OWN:"+mspid+":"+"TRU"+strconv.Itoa(indexTRU), OWNAsBytes2)

	//Record Trace change. Get current trace  and copy to new one
	startKey := "TRACE:" + mspid + ":" + args[0] + ":" + "1"
	endKey := "TRACE:" + mspid + ":" + args[0] + ":" + "999"
	n := 1

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		APIstub.PutState("TRACE:"+mspid+":TRU"+strconv.Itoa(indexTRU)+":"+strconv.Itoa(n), queryResponse.Value)
		n = n + 1
	}

	//Ownership event is added to trace of splitted TRU. EventID is the TRUID of the new TRU.
	tRACE := TRACE{}
	tRACE.OwnerOrgID = mspid
	tRACE.TraceTRUID = args[0]
	tRACE.ChangeDate = currentTime.Format("2006-01-02")
	tRACE.EventType = "Split-Original"
	tRACE.EventID = strconv.Itoa(indexTRU)
	tRACEAsBytes, _ := json.Marshal(tRACE)
	APIstub.PutState("TRACE:"+mspid+":"+args[0]+":"+strconv.Itoa(n), tRACEAsBytes)

	//Ownership event is added to new TRU,
	tRACE2 := TRACE{}
	tRACE2.OwnerOrgID = mspid
	tRACE2.TraceTRUID = args[0]
	tRACE2.ChangeDate = currentTime.Format("2006-01-02")
	tRACE2.EventType = "Split-New"
	tRACE2.EventID = strconv.Itoa(indexTRU)
	tRACEAsBytes2, _ := json.Marshal(tRACE2)
	APIstub.PutState("TRACE:"+mspid+":TRU"+strconv.Itoa(indexTRU)+":"+strconv.Itoa(n), tRACEAsBytes2)

	//Update NextTRU.
	indexTRU = indexTRU + 1
	indexes2 := Indexes{IndexTRU: strconv.Itoa(indexTRU)}
	IndexAsBytes, _ := json.Marshal(indexes2)
	APIstub.PutState("nextTRU", IndexAsBytes)

	return shim.Success(TRUAsBytes2)
}

func (s *SmartContract) changeOWN(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3 - owner of TRU and TRUID.")
	}

	//To be updated later
	mspid := args[2]

	tRUAsBytes, _ := APIstub.GetState(args[1])
	tRU := TRU{}

	json.Unmarshal(tRUAsBytes, &tRU)

	//Check if TRU is released for the organization
	if tRU.ReleaseFlag != mspid {
		return shim.Error("Item is not released to your organization.")
	}

	currentTime := time.Now()

	//Record ownership change
	oWN := OWN{}
	oWN.OwnerOrgID = mspid
	oWN.OwnerTRUID = args[1]
	oWN.OwnDate = currentTime.Format("2006-01-02")
	oWNAsBytes, _ := json.Marshal(oWN)
	APIstub.PutState("OWN:"+mspid+":"+args[1], oWNAsBytes)

	//Record Trace change. Get current trace  and copy to new one
	startKey := "TRACE:" + args[0] + ":" + args[1] + ":" + "1"
	endKey := "TRACE:" + args[0] + ":" + args[1] + ":" + "999"
	n := 1

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		queryResponse, _ := resultsIterator.Next()
		APIstub.PutState("TRACE:"+mspid+":"+args[1]+":"+strconv.Itoa(n), queryResponse.Value)
		n = n + 1
	}

	//Ownership event is added to trace. EventID is old organization ID.
	tRACE := TRACE{}
	tRACE.OwnerOrgID = mspid
	tRACE.TraceTRUID = args[1]
	tRACE.ChangeDate = currentTime.Format("2006-01-02")
	tRACE.EventType = "Ownership"
	tRACE.EventID = args[0]
	tRACEAsBytes, _ := json.Marshal(tRACE)
	APIstub.PutState("TRACE:"+mspid+":"+args[1]+":"+strconv.Itoa(n), tRACEAsBytes)

	//TO DO LATER:After ownership change is performed it cannot be performed again,

	return shim.Success(nil)

}

func (s *SmartContract) updateSCA(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	//serializedID, _ := APIstub.GetCreator()
	//msPID := MSPID{}
	//json.Unmarshal(serializedID, &msPID)
	//callerID := msPID.Mspid
	//fmt.Printf("mspid" + callerID)

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	//To be obtained from caller certificate
	mspid := args[2]

	sCAAsBytes, _ := APIstub.GetState(mspid)
	sCA := SCA{}
	json.Unmarshal(sCAAsBytes, &sCA)

	if args[0] == "Name" {
		sCA.Name = args[1]
	} else if args[0] == "CertificateIDs" {
		sCA.CertificateIDs = args[1]
	} else if args[0] == "SCAType" {
		sCA.SCAType = args[1]
	}

	SCAAsBytes, _ := json.Marshal(sCA)
	APIstub.PutState(args[0], SCAAsBytes)

	return shim.Success(nil)

}

func (s *SmartContract) querySCA(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	//serializedID, _ := APIstub.GetCreator()
	//msPID := MSPID{}
	//json.Unmarshal(serializedID, &msPID)
	//callerID := msPID.Mspid

	//fmt.Printf("mspid" + callerID)

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	sCAAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(sCAAsBytes)

}

func (s *SmartContract) queryAllTRU(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "TRU1"
	endKey := "TRU999999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllTRU:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())

}

func (s *SmartContract) queryAllSCA(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "ORG1"
	endKey := "ORG999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllSCA:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())

}

func (s *SmartContract) queryTRACE(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	startKey := "TRACE:" + args[0] + ":" + args[1] + ":1"
	endKey := "TRACE:" + args[0] + ":" + args[1] + ":999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- Trace:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())

}

func (s *SmartContract) updateTRACE(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	tRACE := TRACE{}
	tRACE.OwnerOrgID = args[0]
	tRACE.TraceTRUID = args[1]
	tRACE.ChangeDate = args[2]
	tRACEAsBytes, _ := json.Marshal(tRACE)

	//TODO: Change trace considering trace sequence
	APIstub.PutState("TRACE:"+args[0]+":"+args[1]+":"+args[2], tRACEAsBytes)

	return shim.Success(nil)

}

func (s *SmartContract) queryNextTRU(APIstub shim.ChaincodeStubInterface) sc.Response {

	NextAsBytes, _ := APIstub.GetState("nextTRU")
	return shim.Success(NextAsBytes)
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting Smart Contract: %s", err)
	}
}