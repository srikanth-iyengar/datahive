package io.datahive.kafkaprocessor.utils;

import java.time.Duration;
import java.util.ArrayList;
import java.util.Collections;
import java.util.Iterator;
import java.util.List;
import java.util.Properties;
import java.util.UUID;

import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.KafkaConsumer;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import groovy.lang.GroovyShell;
import groovy.lang.Script;

public class WorkerUtils {

    private static final Logger logger = LoggerFactory.getLogger(WorkerUtils.class);

    private static final List<Thread> runningThreads = Collections.synchronizedList(new ArrayList<Thread>());
    private static final List<KafkaProducer<String, Object>> runningProducers = Collections.synchronizedList(new ArrayList<KafkaProducer<String, Object>>());

    private static final Properties kafkaProps = new Properties(){{
        setProperty("bootstrap.servers", System.getProperty("bootstrap.servers", "kafka:9092"));
        setProperty("replication.factor", "1");
        setProperty("key.serializer", "org.apache.kafka.common.serialization.StringSerializer");
        setProperty("key.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
        setProperty("value.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
        setProperty("value.serializer", "org.apache.kafka.common.serialization.StringSerializer");
    }};

    public static void startConsumer(String topicName, AfterConsumer afterConsumer) {
        logger.info("Starting the consumer for the topic: " + topicName);
        Thread consumerThread = new Thread(new KafkaConsumerRunnable(topicName, afterConsumer));
        consumerThread.start();
        runningThreads.add(consumerThread);
    }

    public static KafkaProducer<String, Object> createProducer() {
        KafkaProducer<String, Object> producer = new KafkaProducer<>(kafkaProps);
        runningProducers.add(producer);
        return producer;
    }

    public static void produceRecord(KafkaProducer<String, Object> kafkaProducer, String topicName, Object data) {
        ProducerRecord<String, Object> record = new ProducerRecord<String,Object>(topicName, data);
        kafkaProducer.send(record);
    }

    public static Object transform(Object data, String groovyScript) {
        GroovyShell shell = new GroovyShell();
        Script script = shell.parse(groovyScript);
        return script.invokeMethod("transform", data);
    }

    private static class KafkaConsumerRunnable implements Runnable {
        public String topic;
        public AfterConsumer afterConsumer;
        private volatile boolean isRunning;

        private KafkaConsumerRunnable(String topic, AfterConsumer afterConsumer) {
            this.topic = topic;
            this.afterConsumer = afterConsumer;
        }

        public void run() {
            if(this.topic == null || afterConsumer == null) {
                return;
            }
            Properties props = new Properties(kafkaProps){{ 
            }};
            kafkaProps.forEach((key, value) -> {
                props.setProperty((String)key, (String)value);
            });
            props.setProperty("group.id", "consumer-" + UUID.randomUUID().toString());
            this.isRunning = true;
            KafkaConsumer<String, Object> consumer = new KafkaConsumer<>(props);
            consumer.subscribe(Collections.singletonList(this.topic));
            while(isRunning)  {
                try {
                    ConsumerRecords<String,Object> records = consumer.poll(Duration.ofSeconds(3));
                    Iterator<ConsumerRecord<String, Object>> itr = records.iterator();
                    while(itr.hasNext()) {
                        ConsumerRecord<String, Object> record = itr.next();
                        afterConsumer.run(record.value());
                    }
                }
                catch(Exception e) {
                    // e.printStackTrace();
                    Thread.currentThread().interrupt();
                    this.isRunning = false;
                }
            }
            consumer.close();
        }
    }

    public static void stopAll() {
        logger.info("Stopping all the consumer threads gracefully...");
        runningThreads.forEach(thread -> {
            if(thread.isAlive()) {
                thread.interrupt();
            }
        });
        logger.info("Cleanup task successful for KafkaConsumers");
    }

    public static void stopAllProducers() {
        logger.info("Stopping all the producers");
        runningProducers.forEach((producer) -> {
            producer.close();
        });
        logger.info("Cleanup task successful for KafkaProducer");
    }

    @FunctionalInterface
    public static interface AfterConsumer {
        public void run(Object data);
    }
}
