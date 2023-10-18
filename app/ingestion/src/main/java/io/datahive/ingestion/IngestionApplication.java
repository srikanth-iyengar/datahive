package io.datahive.ingestion;

import java.util.stream.Collectors;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import io.datahive.ingestion.protoimpl.KafkaServiceImpl;
import io.grpc.Server;
import io.grpc.ServerBuilder;

@SpringBootApplication
public class IngestionApplication implements CommandLineRunner {

    private final Logger logger = LoggerFactory.getLogger(IngestionApplication.class);

    @Autowired
    private KafkaServiceImpl kafkaServiceImpl;

	public static void main(String[] args) {
		SpringApplication.run(IngestionApplication.class, args);
	}

    @Override
	public void run(String... args) throws Exception {
        Server server = ServerBuilder
            .forPort(Integer.parseInt(System.getProperty("server.port", "8080")))
            .addService(kafkaServiceImpl)
            .build();
        server.start();
        logger.info("Started grpc server on port: {}", server.getPort());
        logger.info("Registered services on the grpc server: [{}]", 
                server.getServices().stream().map(svc -> svc.toString()).collect(Collectors.joining(" ")));
        server.awaitTermination();
	}
}
