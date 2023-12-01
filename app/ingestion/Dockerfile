FROM openjdk:17.0.2
COPY ./target/ingestion-0.0.1-SNAPSHOT.jar /opt/datahive/ingestion.jar
ENTRYPOINT ["java", "-jar", "/opt/datahive/ingestion.jar"]
