package io.datahive.spark;

import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

import io.datahive.spark.utils.SparkUtils;

public class CliRunner implements CommandLineRunner {

    @Override
    public void run(String... args) throws Exception {
        System.out.println("Hello from CLIRunner");
        SparkUtils.StartSparkJob(
                    "test job",
                    "1g", 
                    "2g",
                    "/opt/spark/work-dir/spark.jar",
                    "Test"
                );
        Thread.sleep(5000);
    }

}
