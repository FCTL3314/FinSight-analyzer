#!/bin/bash

bash -c '
  CLUSTER_ID=$(kafka-storage random-uuid)
  echo "Formatting storage with CLUSTER_ID=${CLUSTER_ID}"
  kafka-storage format -t ${CLUSTER_ID} -c /etc/kafka/kafka.properties
  /etc/confluent/docker/run
'
