package io.datahive.kafkaprocessor.utils;

import java.io.IOException;
import java.util.Optional;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.fs.FSDataOutputStream;
import org.apache.hadoop.fs.FileSystem;
import org.apache.hadoop.fs.Path;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class HadoopUtils {

    private static final Logger logger = LoggerFactory.getLogger(HadoopUtils.class);

    private static Configuration hadoopConf = new Configuration() {{
        set("fs.defaultFS", "hdfs://namenode:8020/");
    }};

    public static Optional<FSDataOutputStream> openFile(String path) {
        Optional<Boolean> fileExists = fileExists(path);

        if(fileExists.isEmpty()) {
            logger.error("Error while checking on the file : {}", path);
            return Optional.empty();
        }

        if(!fileExists.get()) {
            createFile(path);
        }

        try(FileSystem fs = FileSystem.get(hadoopConf)) {
            return Optional.of(fs.append(new Path(path), 1024, ()-> {
            }));
        } catch(IOException e) {
            e.printStackTrace();
            logger.error("Error while openning the file: {}", path);
            return Optional.empty();
        }
    }

    public static void createFile(String path) {
        try(FileSystem fs = FileSystem.get(hadoopConf)) {
            fs.createNewFile(new Path(path));
        }
        catch(Exception e) {
            logger.error("Error while creating file in hdfs: {}", path);
        }
    }

    public static Optional<Boolean> fileExists(String path) {
        try(FileSystem fs = FileSystem.get(hadoopConf)) {
            return Optional.of(fs.exists(new Path(path)));
        }
        catch(Exception e) {
            logger.warn("Cannot connect to hdfs");
        }
        return Optional.empty();
    }
}
