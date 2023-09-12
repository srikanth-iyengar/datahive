package io.datahive.kafkaprocessor;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.core.task.TaskExecutor;
import org.springframework.scheduling.annotation.EnableAsync;
import org.springframework.scheduling.concurrent.ThreadPoolTaskExecutor;

import java.util.concurrent.ThreadPoolExecutor;

@SpringBootApplication
@EnableAsync
public class KafkaProcessorApplication {

	@Bean(name = "consumerPool")
	public TaskExecutor consumerPoolExecutor() {
		ThreadPoolTaskExecutor threadPoolTaskExecutor = new ThreadPoolTaskExecutor();
		threadPoolTaskExecutor.setThreadNamePrefix("kafkaConsumer-");
		threadPoolTaskExecutor.setCorePoolSize(1000);
		threadPoolTaskExecutor.setMaxPoolSize(2000);
		threadPoolTaskExecutor.setQueueCapacity(3000);
		return threadPoolTaskExecutor;
	}

	public static void main(String[] args) {
		SpringApplication.run(KafkaProcessorApplication.class, args);
	}

}
