package io.datahive.kafkaprocessor.utils;

import static org.junit.Assert.assertEquals;

import java.util.UUID;

import org.junit.Test;
import org.junit.jupiter.api.DisplayName;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
class HadoopUtilsTest {

    @Test
    @DisplayName("Create_File_In_HDFS")
    public void createFile() {
        String path = "/" + UUID.randomUUID().toString();
        assertEquals(HadoopUtils.fileExists(path), false);
        HadoopUtils.createFile(path);
        assertEquals(HadoopUtils.fileExists(path), true);
    }
}
