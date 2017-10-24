#!/bin/bash

#------------------------------------------------------------------------------
# Start MongoDB Service Broker :
#------------------------------------------------------------------------------
start_monggodb_service_broker() {
    echo 'Starting MongoDB-Service-Broker ...'
    bee run -gendoc=true -downdoc=false
}

#------------------------------------------------------------------------------
# MongoDB Service Broker Main:
#------------------------------------------------------------------------------
start_monggodb_service_broker
