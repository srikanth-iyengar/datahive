package io.datahive.ingestion.utils;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Optional;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.locks.ReentrantLock;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.fs.FSDataOutputStream;
import org.apache.hadoop.fs.FileSystem;
import org.apache.hadoop.fs.Path;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class HadoopUtils {

    private static final Logger logger = LoggerFactory.getLogger(HadoopUtils.class);
    private static final List<FSDataOutputStream> outputStreams = Collections.synchronizedList(new ArrayList<FSDataOutputStream>());

    private static Configuration hadoopConf = new Configuration() {{
        set("fs.defaultFS", "hdfs://" + System.getProperty("hadoop.host", "namenode:8020") + "/");
        set("dfs.replication", "1");
        set("dfs.client.block.write.replace-datanode-on-failure.policy", "DEFAULT");
    }};

    public static Optional<HadoopWriter> openWriter(String path) throws Exception {
        Optional<Boolean> fileExists = fileExists(path);
        if(fileExists.isEmpty()) {
            logger.error("Error while checking on the file : {}", path);
            return Optional.empty();
        }
        if(!fileExists.get()) {
            createFile(path);
        }
        FileSystem fs = FileSystem.get(hadoopConf);
        HadoopWriter writer = new HadoopWriter(path, fs);
        writer.startWriter();
        return Optional.of(writer);
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
            logger.warn("Cannot connect to hdfs while checking: {}", path);
        }
        return Optional.empty();
    }

    public static class HadoopWriter extends Thread {
        private final FileSystem fs;
        private final String path;
        private FSDataOutputStream outstream;

        private List<String> messageCache;

        private final ReentrantLock cacheLock;
        private final ReentrantLock fsLock;
        private volatile boolean isRunning;

        protected HadoopWriter(String path, FileSystem fs) throws Exception {
            this.fs = fs;
            this.path = path;
            this.outstream = this.fs.append(new Path(path));
            this.cacheLock = new ReentrantLock();
            this.messageCache = new ArrayList<>();
            this.fsLock = new ReentrantLock();
            this.isRunning = false;
        }

        public void write(String s) throws IOException{
            if(!isRunning) throw new IOException("The writer is not ready for this operation");
            cacheLock.lock();
            try {
                messageCache.add(s);
            } finally {
                cacheLock.unlock();
            }
        }

        protected void startWriter() {
            this.isRunning = true;
            this.start();
            logger.info("Started writer for file: {}, status: {}", this.path, this.isAlive());
        }

        private void flush() throws IOException {
            this.cacheLock.lock();
            List<String> cache = messageCache;
            this.messageCache = new ArrayList<>();
            this.fsLock.lock();
            cache.forEach(msg -> {
                try {
                    this.outstream.write(msg.getBytes());
                } catch (IOException e) {
                    logger.error("Error while writing the cached messages to : {}", this.path);
                }
            });
            this.outstream.close();
            this.outstream = fs.append(new Path(path));
            this.fsLock.unlock();
            this.cacheLock.unlock();
        }

        @Override
        public void run() {
            long timeout = 1000000;
            while(isRunning) {
                try {
                    boolean canLockFs = this.fsLock.tryLock(timeout, TimeUnit.MICROSECONDS);
                    if(canLockFs) {
                        flush();
                    }
                    else {
                        
                    }
                    timeout += timeout / 50;
                }
                catch (Exception e) {
                    logger.error("Error while trying to flush the output: {}", e.getMessage());
                }
            }
        }
    }
}
