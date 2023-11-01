package io.datahive.spark;

import java.io.IOException;
import java.util.stream.Collectors;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.CommandLineRunner;

import io.datahive.spark.protoimpl.SparkServiceImpl;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.protobuf.services.ProtoReflectionService;

@SpringBootApplication
public class SparkApplication implements CommandLineRunner {

    private final SparkServiceImpl sparkServiceImpl;
    private final Logger logger = LoggerFactory.getLogger(SparkApplication.class);

    public SparkApplication(SparkServiceImpl sparkServiceImpl) {
        this.sparkServiceImpl = sparkServiceImpl;
    }

	public static void main(String[] args) {
		SpringApplication.run(SparkApplication.class, args);
	}

    public void run(String ...args) throws IOException, InterruptedException {
        Server server = ServerBuilder
                        .forPort(Integer.valueOf(System.getProperty("server.port", "8080")))
                        .addService(sparkServiceImpl)
                        .addService(ProtoReflectionService.newInstance())
                        .build();
        server.start();
        logger.info("Started grpc server on port: {}", server.getPort());
        logger.info("Registered services on the grpc server: [{}]", 
                server.getServices().stream().map(svc -> svc.toString()).collect(Collectors.joining(" ")));
        server.awaitTermination();
    }

}
