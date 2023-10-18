package io.datahive.ingestion.utils;


import org.apache.hadoop.fs.FSDataInputStream;
import org.apache.hadoop.fs.FSDataOutputStream;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;

import static org.junit.Assert.assertEquals;

import java.io.DataOutputStream;
import java.util.Optional;
import java.util.UUID;


@SpringBootTest
class HadoopUtilsTests {
    @Test
    @DisplayName("Create_File_In_HDFS")
    void createFile() {
        String path = "/" + UUID.randomUUID();
        Assertions.assertEquals(HadoopUtils.fileExists(path).get(), false);
        HadoopUtils.createFile(path);
        Assertions.assertEquals(HadoopUtils.fileExists(path).get(), true);
    }

    @Test
    @DisplayName("Write_File_In_HDFS")
    void writeFile() {
        String path = "/" + UUID.randomUUID();
//        Optional<FSDataOutputStream> outputStream = HadoopUtils.openFile(path);
    }
}
