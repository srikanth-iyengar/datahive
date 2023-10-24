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

    public String startConsumerWithTransformations(String inTopicName, String groovyScript, String outTopicName, String pipelineId) {
        KafkaProducer<String, Object> producer = WorkerUtils.createProducer();
        return WorkerUtils.startConsumer(inTopicName, (data) -> {
            WorkerUtils.produceRecord(producer, outTopicName, WorkerUtils.transform(data, groovyScript));
        }, pipelineId);
    }

    public String startConsumerAndPushHadoop(String inTopicName, String hdfsFileName, String pipelineId) throws Exception {
        ObjectMapper objectmapper = new ObjectMapper();
        Optional<HadoopUtils.HadoopWriter> optWriter = HadoopUtils.openWriter(hdfsFileName);
        if(optWriter.isEmpty()) {
            logger.error("Cannot open hdfs file: {}", hdfsFileName);
            return "null";
        }
        HadoopUtils.HadoopWriter  writer = optWriter.get();
        return WorkerUtils.startConsumer(inTopicName, (data) -> {
            try {
                writer.write(objectmapper.writeValueAsString(data) + "\n");
            }
            catch(Exception e) {
                logger.error("Error while writing to file: {}", hdfsFileName);
            }
        }, pipelineId);
    }
}

