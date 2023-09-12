package io.datahive.kafkaprocessor.worker;
import groovy.lang.GroovyShell;
import groovy.lang.Script;
import org.apache.kafka.clients.consumer.ConsumerConfig;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.KafkaConsumer;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Component;

import java.lang.reflect.Method;
import java.time.Duration;
import java.util.Collections;
import java.util.Properties;
import java.util.UUID;

@Component
public class DatahiveConsumer {

    @Autowired
    @Qualifier("kafkaProps")
    private Properties props;

    @SuppressWarnings("unchecked")
    @Async("consumerPool")
    public void startConsumerWithTransformations(String inTopicName, String groovyScript, String outTopicName) {
        try {
            props.setProperty(ConsumerConfig.GROUP_ID_CONFIG, "consumer-" + UUID.randomUUID().toString());
            KafkaConsumer<String, Object> consumer = new KafkaConsumer<>(props);
            KafkaProducer<String, Object> producer = new KafkaProducer<>(props);
            consumer.subscribe(Collections.singletonList(inTopicName));
            while (true) {
                ConsumerRecords<String, Object> records = consumer.poll(Duration.ofMillis(1000));
                records.forEach(record -> {
                    GroovyShell shell = new GroovyShell();
                    Script script = shell.parse(groovyScript);
                    Object result = script.invokeMethod("transform", record.value());
                    ProducerRecord<String, Object> transformedObject = new ProducerRecord<>("test-out-topic", result);
                    producer.send(transformedObject);
                });
            }
        }
        catch (Exception e) {
            e.printStackTrace();
        }
    }

    @Async("consumerPool")
    public void startConsumerAndPushHadoop(String inTopicName, String hdfsFileName) {

    }
}
