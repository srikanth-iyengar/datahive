package io.datahive.spark.utils;

import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.net.HttpURLConnection;
import java.net.MalformedURLException;
import java.net.URL;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class FileUtils {

    private final static Logger logger  = LoggerFactory.getLogger(FileUtils.class);
    
    public static void getFileFromMinIo(String url) {
        String uri = "/data" + url.substring(17);
        File resource = new File(uri);
        if(!resource.getParentFile().exists()) {
            resource.getParentFile().mkdirs();
        }
        try {
            resource.createNewFile();
            URL fileLocation = new URL(url);
            HttpURLConnection conn = (HttpURLConnection) fileLocation.openConnection();
            if(conn.getResponseCode() != 200) {
                throw new IOException("File not found in remote storage bucket: " + url);
            }
            try(InputStream inStream = conn.getInputStream(); 
                FileOutputStream outStream = new FileOutputStream(resource)){
                int bytesRead = -1;
                byte []buffer = new byte[1000];
                while((bytesRead = inStream.read(buffer)) != -1) {
                    outStream.write(buffer, 0, bytesRead);
                }
            }
            finally {
            }
        }
        catch(MalformedURLException e) {
            e.printStackTrace();
            logger.error("Malformed url exception: {}", e.getMessage());
        }
        catch(IOException e) {
            e.printStackTrace();
            logger.error("Failed to create file: {}", uri);
        }
    }
}
