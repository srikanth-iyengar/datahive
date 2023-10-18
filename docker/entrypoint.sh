#!/bin/bash -e
apt-get update
apt-get install python3-pip -y
pip3 install confluent-kafka
python3 /generate.py
