# Hyperledger Summer Internship: Analyzing Hyperledger Fabric Ledger, Transactions and Logs using Elasticsearch and Kibana

Mentor: Salman Baset [salmanbaset](https://github.com/salmanbaset)

Mentee: Balazs Prehoda [balazsprehoda](https://github.com/balazsprehoda)

## Contents
1. [Description](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#description)  
2. [Expected Outcome](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#expected-outcome)  
3. [Prerequisites](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#prerequisites)  
4. [Getting Started](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#getting-started)  
5. [Fabric Network](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#fabric-network)  
  5.1. [Basic network](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#basic-network)  
  5.2. [Multichannel network](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#multichannel-network)  
  5.3. [(Optional) Generate makefile configuration](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#optional-generate-makefile-configuration)  
  5.4. [Start the network](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#start-the-network)   
  5.5. [Stop the network](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#stop-the-network)  
6. [Dummyapp Application](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#dummyapp-application)  
  6.1. [Installation](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#installation)  
  6.2. [Configuration](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#configuration)  
  6.3. [User enrollment and registration](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#user-enrollment-and-registration)  
  6.4. [Invoke transactions](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#invoke-transactions)  
  6.5. [Query](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#query)  
7. [Elastic Stack](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#elastic-stack)  
  7.1. [Credit](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#credit)  
  7.2. [Description](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#description-1)  
  7.3. [Start](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#start)  
  7.4. [Stop and destroy](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#stop-and-destroy)  
8. [Beats Agent](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#beats-agent)  
  8.1. [Environment setup](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#environment-setup)  
  8.2. [Configure fabricbeat](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#configure-fabricbeat)  
  8.3. [Build fabricbeat](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#build-fabricbeat)  
  8.4. [Start fabricbeat](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#start-fabricbeat)   
  8.5. [Stop fabricbeat](https://github.com/balazsprehoda/hyperledger-elastic/tree/test-build-path#stop-fabricbeat)  
  

## Description

Each blockchain platform, including Hyperledger Fabric, provide a way to record information on blockchain in an immutable manner. In the case of Hyperledger Fabric, information is recorded as a `key-value` pair. All previous updates to a `key` are recorded in the ledger, but only the latest value of a `key` can be easily queried using CouchDB; the previous updates are only available in ledger files. This mechanism makes it challenging to perform analysis of updates to a `key`, a necessary requirement for information provenance.

The goal of this project is to

1. write a Elastic beats module (in Go), that will ship ledger data to Elasticsearch instance

2. create generic Kibana dashboards that will allow selection of a particular key and visualization updates to it (channel, id, timestamp etc)

Time permitting, the dashboards can be extended to analyze Fabric logs and in-progress transaction data, as well as creating dashboards similar to Hyperledger Explorer.

Of course, a blockchain solution can track information provenance in multiple ways. In one such mechanism, a solution may always write new key-value pairs to blockchain, and maintain the relationship among key-value pairs within the solution (off-chain), instead of blockchain. This project does not concern itself on how a solution manages relationship among key-value pairs.

## Expected Outcome

A open source implementation, eventually available as Hyperledger Labs, containing:

Elastic beats plugin for Hyperledger Fabric
Kibana dashboards
Dashboards similar to Hyperledger Explorer
Create a setup for generating various dummy data in various configurations
One peer / CA / order, single user for initial testing
A four peers/CA setup with two channels, and two users each associated with two peers. Select (e.g.) 10 keys (through configuration file), to which these users write data, for at least one value per key.

## Prerequisites

Please make sure that you have set up the environment for the project. Follow the steps listed in [Prerequisites](https://github.com/balazsprehoda/hyperledger-elastic/blob/test-build-path/Prerequisites.md).   

## Getting Started

To get started with the project, clone the git repository. It is important that you place it under $GOPATH/src/github.com  
```
cd $GOPATH/src/github.com  
git clone https://github.com/balazsprehoda/hyperledger-elastic.git  
```

## Fabric Network

The `network` directory contains two network setups, `basic` and `multichannel`. 

### Basic network

It is a simple test network with 4 organizations, 1 peer each, a solo orderer communicating over TLS and a sample chaincode called `dummycc`. It writes deterministically generated hashes and (optionally) previous keys as value to the ledger.  

### Multichannel network

It is a test network setup with 4 organizations, 2 peers each, a solo orderer communicating over TLS and two channels:

1. `fourchannel`:
  * members: all four organizations
  * chaincode: `dummycc`: It writes deterministically generated hashes and (optionally) previous keys as value to the ledger.
2. `twochannel`:
  * members: only `Org1` and `Org3`
  * chaincode: `fabcar`: The classic fabcar example chaincode extended with a `getHistoryForCar()` chaincode function.

### (Optional) Generate makefile configuration
```
make generate ABSPATH={ABSOLUTE_PATH_TO_SRC_DIR}
```

### Start the network

To generate crypto and setup the network, make sure that the `GOPATH` variable is set correctly. After that, set `ABSPATH` to point to the parent folder of `hyperledger-elastic` (e.g., `export ABSPATH=$GOPATH/src/github.com`). When we are done with setting these environment variables, we can start the network. Issue the following command in the `network/basic` directory to start the basic network, or in the `network/multichannel` directory to start the multichannel network:

```
make start  
```

We can enter the CLI by issuing the command 
```
docker exec -it cli bash  
```

Inside the CLI, the `/scripts` folder contains the scripts that can be used to install, instantiate and invoke chaincode (though the `make start` command takes care of installation and instantiation).

### Stop the network

To stop the network and delete all the generated data (crypto material, channel artifacts and dummyapp wallet), run

```
make destroy
```

## Dummyapp Application  
The dummyapp application is used to generate users and transactions. It can connect to both the basic and the multichannel networks.  
The commands in this section should be issued from the `hyperledger-elastic/apps/dummyapp` directory.

### Installation
Before the first run, we have to install the necessary node modules:
```
npm install
```

### Configuration
The `config.json` contains the configuration for the application. We can configure the channel and chaincode name that we want our application to use, the users we want to enroll and the transactions we want to initialize. Transactions have 4 fields:
1. `user`: This field is required. We have to specify which user to use when making the transaction.
2. `txFunction`: This field is required. We have to specify here the chaincode function that should be called.
3. `key`: This field is required. We have to specify here the key to be written to the ledger.
4. `previousKey`: This field is optional. We can specify here the key to which the new transaction (key-value pair) is linked.

###  User enrollment and registration

Setup path of the network configuration (basic or multichannel shipped with this repository)
```
export ABSPATH={ABSOLUTE_PATH}/src/github.com/hyperledger-elastic/network/basic/
```

To enroll admins, register and enroll users, run the following command:
```
make users
```

###  Invoke transactions
To add key-value pairs, run
```
make invoke
```

### Query
To query a specific key, run
```
make query KEY=key1
```
To query all key-value pairs, run
```
make query-all
```

## Elastic Stack
This project includes an Elasticsearch and Kibana setup to index and visualize blockchain data.  
The commands in this section should be issued from the `hyperledger-elastic/stack` folder.

### Credit
This setup is borrowed from
https://github.com/maxyermayank/docker-compose-elasticsearch-kibana

### Description
Kibana container and Elasticsearch cluster with nginx and 3 Elasticsearch containers. To view Kibana in browser, navigate to http://localhost:5601

### Start
To start the containers, issue
```
make start
```

### Stop and destroy
To stop the containers, issue
```
make destroy
```
To stop the containers and remove old data, run
```
make erase
```

## Beats Agent

The fabricbeat beats agent is responsible for connecting to a specified peer, periodically querying its ledger, processing the data and shipping it to Elasticsearch. Multiple instances can be run at the same time, each querying a different peer and sending its data to the Elasticsearch cluster.  
The commands in this section should be issued from the `hyperledger-elastic/agent/fabricbeat` directory.

### Environment setup

Before configuring and building the fabricbeat agent, we should make sure that the `GOPATH` variable is set correctly. Then, we have to add `$GOPATH/bin` to the `PATH`:
```
export PATH=$PATH:$GOPATH/bin
```
After that, we have to set the `BEAT_PATH` variable to point to the fabricbeat folder:
```
export BEAT_PATH=$GOPATH/src/github.com/hyperledger-elastic/agent/fabricbeat
```

We want to use vendoring instead of go modules, so we have to make sure `GO111MODULE` is set to `auto` (it is the default):  
```
export GO111MODULE=auto
```  

### Configure fabricbeat

We can configure the agent using the `fabricbeat.yml` file. If we want to update the generated config file, we can edit `fabricbeat/_meta/beat.yml`, then run
```
make update
```

### Build fabricbeat

To build the agent, issue the following command:
```
make
```

### Start fabricbeat

To start the agent, issue the following command from the `fabricbeat` directory:
```
./fabricbeat -e -d "*"
```
If we want to use the agent with one of the example networks from the `hyperledger-elastic/network` folder, we can start the agent using:
```
ORG_NUMBER=1 ORG_NAME=org1 ./fabricbeat -e -d "*"
```
The variables passed are used in the configuration (`fabricbeat.yml`). To connect to another peer, change the configuration (and/or the passed variables) accordingly.

### Stop fabricbeat

To stop the agent, simply type `Ctrl+C`
