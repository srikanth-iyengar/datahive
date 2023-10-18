package io.datahive.ingestion.worker;



import io.datahive.ingestion.utils.HadoopUtils;
import io.datahive.ingestion.utils.WorkerUtils;

import org.apache.kafka.clients.producer.KafkaProducer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Component;

import com.fasterxml.jackson.databind.ObjectMapper;

import java.util.Optional;

@Component
public class DatahiveKafkaWorker {

    private final Logger logger = LoggerFactory.getLogger(DatahiveKafkaWorker.class);

    public void startConsumerWithTransformations(String inTopicName, String groovyScript, String outTopicName) {
        KafkaProducer<String, Object> producer = WorkerUtils.createProducer();
        WorkerUtils.startConsumer(inTopicName, (data) -> {
            WorkerUtils.produceRecord(producer, outTopicName, WorkerUtils.transform(data, groovyScript));
        });
    }

    public void startConsumerAndPushHadoop(String inTopicName, String hdfsFileName) throws Exception {
        ObjectMapper objectmapper = new ObjectMapper();
        Optional<HadoopUtils.HadoopWriter> optWriter = HadoopUtils.openWriter(hdfsFileName);
        if(optWriter.isEmpty()) {
            logger.error("Cannot open hdfs file: {}", hdfsFileName);
            return;
        }
        HadoopUtils.HadoopWriter  writer = optWriter.get();
        WorkerUtils.startConsumer(inTopicName, (data) -> {
            try {
                writer.write(objectmapper.writeValueAsString(data) + "\n");
            }
            catch(Exception e) {
                e.printStackTrace();
                logger.error("Error while writing to file: {}", hdfsFileName);
            }
        });
    }
}

