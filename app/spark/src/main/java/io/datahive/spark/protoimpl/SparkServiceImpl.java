package io.datahive.spark.protoimpl;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

import io.datahive.spark.proto.JobInfo;
import io.datahive.spark.proto.PingJob;
import io.datahive.spark.proto.SparkJobRequest;
import io.datahive.spark.proto.AnalysisServiceGrpc.AnalysisServiceImplBase;
import io.datahive.spark.utils.SparkUtils;
import io.grpc.stub.StreamObserver;

@Service
public class SparkServiceImpl extends AnalysisServiceImplBase {

    private final Logger logger = LoggerFactory.getLogger(SparkServiceImpl.class);

    @Override
    public void startSparkJob(SparkJobRequest request, StreamObserver<JobInfo> responseObserver) {
        JobInfo resp = SparkUtils.StartSparkJob(request.getJobName(),
                                        request.getDriverMemory(),
                                        request.getExecutorMemory(),
                                        request.getResourceLocation(),
                                        request.getMainClass());
        logger.info("Scheduling a worker with JobName: {}, WorkerId: {}", request.getJobName(), resp.getWorkerId());
        responseObserver.onNext(resp);
        responseObserver.onCompleted();
    }

    @Override
    public void checkJob(PingJob request, StreamObserver<JobInfo> responseObserver) {
        logger.info("Status check for the worker: {}", request.getWorkerId());
        JobInfo resp = SparkUtils.pingSparkWorker(request.getWorkerId());
        responseObserver.onNext(resp);
        responseObserver.onCompleted();
    }

    @Override
    public void stopSparkJob(PingJob request, StreamObserver<JobInfo> responseObserver) {
        JobInfo resp = SparkUtils.stopSparkWorker(request.getWorkerId());
        responseObserver.onNext(resp);
        responseObserver.onCompleted();
    }
}
