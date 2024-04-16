# 0g-storage-scan

Implementation of a Storage Scan on 0g Storage Node.

[![MIT licensed][1]][2]

[1]: https://img.shields.io/badge/license-MIT-blue.svg
[2]: LICENSE

## Prerequisites

### Hardware Requirements

Recommended:

* Fast CPU with 4+ cores
* 8GB+ RAM
* High Performance SSD with at least 100GB free space

### Software Requirements

Storage Scan requires the following third party dependencies:

* Go v1.21.x+ *(used for building code)*
* Docker Compose v2.24.x+ *(used for running services)*
* MySQL Server v8.0.x+ *(mainly used for off-storage-node persistent storage)*

## Running Storage Scan

Storage Scan is comprised of several components as below:

* Data Sync *(synchronizes storage node/blockchain data to persistent storage)*
* Statistics *(statistics storage node/blockchain data which have been persisted to storage)*
* Open API *(provides a collection of common rest apis)*

### Quick Start

You can use the `run.sh` script in the `script` directory to start storage scan.

```shell
$ chmod u+x ./script/run.sh
$ ./script/run.sh
```

It will start up services as below:

```shell
$ docker-compose ps 

    Name               Command        SERVICE         State                              
---------------------------------------------------------------
0g-scan-api       "./0g-scan api"       api            Up  
0g-scan-stat      "./0g-scan stat"      stat           Up   
0g-scan-sync      "./0g-scan sync"      sync           Up   
```