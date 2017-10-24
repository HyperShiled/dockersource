#!/bin/bash

#------------------------------------------------------------------------------
# Start DB2 Service Broker :
#------------------------------------------------------------------------------
start_db2_service_broker() {
    echo 'Starting DB2-Service-Broker Service ...'
    su - db2inst1 -c "cd /home/db2inst1/go/src/github.com/compassorg/db2-service-broker && bee run -gendoc=true -downdoc=false"
}

#------------------------------------------------------------------------------
# HOR DB2 Service Broker Main:
#------------------------------------------------------------------------------
start_db2_service_broker
