#!/bin/bash
cd /opt/datahive/kafka-processor
./mvnw clean package
./mvnw spring-boot:run
