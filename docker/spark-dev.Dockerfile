FROM openjdk:17-jdk-bullseye
RUN apt-get update
RUN apt-get install python3-pip -y
RUN pip3 install spark-submit
