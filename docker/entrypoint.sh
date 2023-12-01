#!/bin/bash -e
apt-get update
apt-get install python3-pip -y
pip3 install kafka-python
python3 /stream.py
