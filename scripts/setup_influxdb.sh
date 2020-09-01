#!/usr/bin/env bash

echo "1) launching influx"
docker run -d --name influxdb -p 9999:9999 quay.io/influxdb/influxdb:2.0.0-beta
while ! $(curl -sS 'http://localhost:9999/ready' | grep -q ready); do echo 'Waiting for influx...'; sleep 1; done

echo "2) exporting env var for terraform acceptance tests"
export $INFLUXDB_V2_URL="http://localhost:9999"
export $INFLUXDB_V2_USERNAME="test_acc_user"
export $INFLUXDB_V2_PASSWORD="test_acc_pass"