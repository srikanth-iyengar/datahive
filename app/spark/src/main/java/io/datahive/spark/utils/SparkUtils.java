package io.datahive.spark.utils;

import java.util.Collections;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

import org.apache.spark.launcher.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import io.datahive.spark.exception.JobAlreadyStopped;
import io.datahive.spark.exception.JobNotFound;
import io.datahive.spark.proto.JobInfo;
import io.datahive.spark.proto.JobInfo.Status;

public class SparkUtils {
    private final static Map<String, Process> jobs = Collections.synchronizedMap(new HashMap<>());

    private final Logger logger = LoggerFactory.getLogger(SparkUtils.class);
    public static JobInfo StartSparkJob(
            String jobName,
            String driverMemory,
            String executorMemory,
            String resourceLocation,
            String mainClass) {
        try {
            var workerId = UUID.randomUUID().toString();
            if(!resourceLocation.startsWith("http://minio:9000")) {
                return JobInfo.newBuilder()
                            .setJobStatus(Status.INTERNAL_FAILRUE)
                            .setMessage("Invalid resource url, all the spark resource should start with http://minio:9000")
                            .build();
            }
            try {
                FileUtils.getFileFromMinIo(resourceLocation);
            }
            catch(Exception e) {
                return JobInfo.newBuilder()
                    .setJobStatus(Status.INTERNAL_FAILRUE)
                    .setMessage(e.getMessage())
                    .build();
            }
            var spark = new SparkLauncher()
                .setMaster("local[*]")
                .setMainClass(mainClass)
                .setSparkHome("/opt/spark")
                .setAppResource(resourceLocation)
                .setAppName(jobName + "-" + workerId)
                .setDeployMode("cluster")
                .launch();
            jobs.put(workerId, spark);
            return JobInfo.newBuilder()
                          .setWorkerId(workerId)
                          .setJobStatus(JobInfo.Status.RUNNING)
                          .build();
        }
        catch(Exception e) {
            e.printStackTrace();
            return JobInfo.newBuilder()
                          .setJobStatus(JobInfo.Status.SCHEDULING_FAILURE)
                          .setMessage(e.getMessage())
                          .build();
        }
    }

    public static JobInfo pingSparkWorker(String workerId) {
        try {
            Process p = jobs.get(workerId);
            if(p == null) {
                throw new JobNotFound();
            }
            else if(!p.isAlive()) {
                throw new JobAlreadyStopped();
            }
            return JobInfo.newBuilder()
                          .setWorkerId(workerId)
                          .setJobStatus(JobInfo.Status.RUNNING)
                          .build();
        }
        catch(JobNotFound e) {
            return JobInfo.newBuilder()
                          .setWorkerId(workerId)
                          .setJobStatus(JobInfo.Status.NOT_FOUND)
                          .build();
        }
        catch(JobAlreadyStopped e) {
            return JobInfo.newBuilder()
                          .setWorkerId(workerId)
                          .setJobStatus(JobInfo.Status.STOPPED)
                          .build();
        }
    }

    public static JobInfo stopSparkWorker(String workerId) {
        try {
            Process p = jobs.get(workerId);
            if(p == null) {
                throw new JobNotFound();
            }
            else if(!p.isAlive()) {
                throw new JobAlreadyStopped();
            }
            p.destroy();
            return JobInfo.newBuilder()
                          .setWorkerId(workerId)
                          .setJobStatus(JobInfo.Status.STOPPED)
                          .build();
                   
        }
        catch(JobNotFound e) {
            return JobInfo.newBuilder()
                          .setJobStatus(JobInfo.Status.NOT_FOUND).build();
        }
        catch(JobAlreadyStopped e) {
            return JobInfo.newBuilder()
                          .setWorkerId(workerId)
                          .setJobStatus(JobInfo.Status.TERMINATED).build();
        }
    }

    public void stopAll() {
        jobs.entrySet().forEach(entry -> {
            try {
                entry.getValue().destroy();
            } catch(Exception e) {
                logger.warn("Cannot kill process for the worker: {}, Killing it forcefully...", entry.getKey());
                entry.getValue().destroyForcibly();
            }
        });
    }
}
