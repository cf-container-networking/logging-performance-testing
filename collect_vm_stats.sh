#!/bin/bash

OUTPUT_FILE=$1
INSTANCE_IP=$(cf ssh proxy -c 'env | grep CF_INSTANCE_IP' | cut -d'=' -f 2)

while true; do
  bosh-cli vms --vitals | grep "$INSTANCE_IP" | cut -d' ' -f 4-9 >> "$OUTPUT_FILE";
  sleep 2;
done
