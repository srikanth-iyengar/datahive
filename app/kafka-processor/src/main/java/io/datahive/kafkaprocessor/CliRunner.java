package io.datahive.kafkaprocessor;

import io.datahive.kafkaprocessor.worker.DatahiveConsumer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

@Component
public class CliRunner implements CommandLineRunner {

    @Autowired
    private DatahiveConsumer datahiveConsumer;

    @Override
    public void run(String ...args) {
        String groovyScript = "import groovy.json.JsonOutput\n" +
                "import groovy.json.JsonSlurper\n" +
                "def transform(record) {\n"
                + "	def jsonSlurper = new JsonSlurper();\n"
                + "	def jsonObject  = jsonSlurper.parseText(record);\n"
                + "	jsonObject.sum = jsonObject.a + jsonObject.b;\n"
                + " return JsonOutput.toJson(jsonObject);\n"
                + "};\n"
                + "this";
        datahiveConsumer.startConsumerWithTransformations("test-in-topic", groovyScript, "test-out-topic");
    }
}
