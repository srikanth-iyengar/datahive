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
import org.springframework.beans.factory.annotation.Value;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Component;

import org.apache.hadoop.conf.*;
import org.apache.hadoop.fs.FileSystem;
import org.apache.hadoop.fs.Path;
import org.apache.hadoop.util.Progressable;

import java.io.IOException;
import java.io.OutputStream;
import java.net.URI;
import java.net.URISyntaxException;
import java.time.Duration;
import java.util.Collections;
import java.util.Properties;
import java.util.UUID;

@Component
public class DatahiveConsumer {

    @Autowired
    @Qualifier("kafkaProps")
    private Properties props;

    @Value("${hadoop.host}")
    private String hadoopHost;

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
                    ProducerRecord<String, Object> transformedObject = new ProducerRecord<>(outTopicName, result);
                    producer.send(transformedObject);
                });
            }
        }
        catch (Exception e) {
            e.printStackTrace();
        }
    }

    @Async("consumerPool")
    public void startConsumerAndPushHadoop(String inTopicName, String hdfsFileName) throws IOException, URISyntaxException {
        props.setProperty(ConsumerConfig.GROUP_ID_CONFIG, "consumer-" + UUID.randomUUID().toString());
        KafkaConsumer<String, String> kafkaConsumer = new KafkaConsumer<>(props);
        Configuration hadoopConfiguration = new Configuration();
        FileSystem hdfs = FileSystem.get(new URI(hadoopHost), hadoopConfiguration);
        Path file = new Path(hdfsFileName);
        OutputStream oStream = hdfs.create(file, new Progressable() {
            public void progress() {
            }
        });
        kafkaConsumer.subscribe(Collections.singletonList(inTopicName));
        while(true) {
            ConsumerRecords<String, String> records = kafkaConsumer.poll(1000);
            records.forEach(record -> {
                try {
                    oStream.write(record.value().getBytes());
                    oStream.flush();
                }
                catch(IOException e) {
                    e.printStackTrace();
                }
            });
        }
    }
}
