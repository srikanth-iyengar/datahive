package io.datahive.ingestion.hooks;


import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;
import io.datahive.ingestion.utils.WorkerUtils;


@Component
public class ShutdownHook implements CommandLineRunner {

    @Override
    public void run(String ...args) {
        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            WorkerUtils.stopAll();
        }));
    }
}
