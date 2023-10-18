package io.datahive.kafkaprocessor.factory;

import java.util.Properties;
import java.util.UUID;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Component;


@Component
public class KafkaSessionFactory {

    @Value("${kafka-properties.bootstrap.servers}")
    private String bootstrapServers;

    @Bean(name = "kafkaProps")
    public Properties initKafkaConnectionFactory() {
        Properties props = new Properties();
        props.setProperty("bootstrap.servers", bootstrapServers);
        props.setProperty("replication.factor", "1");
        props.setProperty("key.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
        props.setProperty("value.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
        props.setProperty("key.serializer", "org.apache.kafka.common.serialization.StringSerializer");
        props.setProperty("value.serializer", "org.apache.kafka.common.serialization.StringSerializer");
        return props;
    }
}