package io.datahive.kafkaprocessor.worker;

import io.datahive.kafkaprocessor.utils.HadoopUtils;
import io.datahive.kafkaprocessor.utils.WorkerUtils;

import org.apache.hadoop.fs.FSDataOutputStream;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Component;

import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;
import java.io.OutputStream;
import java.net.URISyntaxException;
import java.util.Optional;

@Component
public class DatahiveConsumer {

    private final Logger logger = LoggerFactory.getLogger(DatahiveConsumer.class);

    public void startConsumerWithTransformations(String inTopicName, String groovyScript, String outTopicName) {
        KafkaProducer<String, Object> producer = WorkerUtils.createProducer();
        WorkerUtils.startConsumer(inTopicName, (data) -> {
            WorkerUtils.produceRecord(producer, outTopicName, WorkerUtils.transform(data, groovyScript));
        });
    }

    public void startConsumerAndPushHadoop(String inTopicName, String hdfsFileName) {
        Optional<FSDataOutputStream> optOutstream = HadoopUtils.openFile(hdfsFileName);
        if(optOutstream.isEmpty()) {
            logger.error("Cannot open hdfs file: {}", hdfsFileName);
            return;
        }
        FSDataOutputStream outstream = optOutstream.get();
        ObjectMapper objectmapper = new ObjectMapper();
        WorkerUtils.startConsumer(inTopicName, (data) -> {
            try {
                outstream.write(objectmapper.writeValueAsString(data).getBytes());
                outstream.flush();
            }
            catch(Exception e) {
                e.printStackTrace();
                logger.error("Error while writing to file: {}, status: {}", hdfsFileName, outstream.getPos());
            }
        });
    }
}
