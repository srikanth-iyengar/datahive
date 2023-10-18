package io.datahive.ingestion.protoimpl;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import io.datahive.ingestion.proto.ConsumerRequestWithHadoop;
import io.datahive.ingestion.proto.ConsumerRequestWithTransformation;
import io.datahive.ingestion.proto.IngestionServiceGrpc;
import io.datahive.ingestion.proto.Response;
import io.datahive.ingestion.worker.DatahiveKafkaWorker;
import io.grpc.stub.StreamObserver;

@Service
public class KafkaServiceImpl extends IngestionServiceGrpc.IngestionServiceImplBase {

    private final Logger logger = LoggerFactory.getLogger(KafkaServiceImpl.class);

    @Autowired
    private DatahiveKafkaWorker kafkakWorker;
    
    @Override
    public void startKafkaConsumerWithTransformation(ConsumerRequestWithTransformation request,
            StreamObserver<Response> responseObserver) {
        try {
            kafkakWorker.startConsumerWithTransformations(request.getInTopic(), request.getGroovyScript(), request.getOutTopic());
            Response response = Response.newBuilder().setStatus("SUCCESS").build();
            responseObserver.onNext(response);
            logger.info("Started a kafka consumer for the topic: {}", request.getInTopic());
        }
        catch(Exception e) {
            Response response = Response.newBuilder().setStatus("FAIL").build();
            responseObserver.onNext(response);
            logger.warn("Failed to start a kafka consumer for the topic: {}, reason: {}", 
                    request.getInTopic(), e.getMessage());
        }
        finally {
            responseObserver.onCompleted();
        }
    }

    @Override
    public void startKafkaConsumerWithHDFSPlugin(ConsumerRequestWithHadoop request,
            StreamObserver<Response> responseObserver) {
        try {
            kafkakWorker.startConsumerAndPushHadoop(request.getInTopic(), request.getHdfsFileName());
            Response response = Response.newBuilder().setStatus("SUCCESS").build();
            responseObserver.onNext(response);
            logger.info("Start kafka consumer for topic: {}, hdfsFile: {}", request.getInTopic(), request.getHdfsFileName());
        }
        catch(Exception e) {
            Response response = Response.newBuilder().setStatus("FAIL").build();
            responseObserver.onNext(response);
            logger.warn("Failed to start a kafka consumer for the topic: {}, reason: {}", 
                    request.getInTopic(), e.getMessage());
        }
        finally {
            responseObserver.onCompleted();
        }
    }
}
