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
        String groovyScript = "import groovy.json.JsonSlurper\n" +
                "import groovy.json.JsonOutput\n" +
                "def transform(record) {\n" +
                "\tdef jsonSlurper = new JsonSlurper();\n" +
                "\tdef kafkaData = jsonSlurper.parseText(record);\n" +
                "    def totalVoltage = 0.0;\n" +
                "    def totalCurrent = 0.0;\n" +
                "    def totalPower = 0.0;\n" +
                "    def gridCount = 0;\n" +
                "\n" +
                "    def stationName = kafkaData.station.name\n" +
                "\tkafkaData[stationName].each {gridData ->\n" +
                "\t        totalVoltage += gridData.distributor_grid.voltage;\n" +
                "\t        totalCurrent += gridData.distributor_grid.current;\n" +
                "\t        totalPower += gridData.distributor_grid.power;\n" +
                "\t        gridCount++;\n" +
                "\t};\n" +
                "\n" +
                "    def averageVoltage = totalVoltage / gridCount;\n" +
                "    def averageCurrent = totalCurrent / gridCount;\n" +
                "    def averagePower = totalPower / gridCount;\n" +
                "\n" +
                "    def newPayload = [\n" +
                "        timestamp: kafkaData['timestamp'],\n" +
                "        station_details: kafkaData['station'],\n" +
                "        average_metrics: [\n" +
                "            voltage: averageVoltage,\n" +
                "            current: averageCurrent,\n" +
                "            power: averagePower\n" +
                "        ]\n" +
                "    ];\n" +
                "\n" +
                "    def newJsonPayload = JsonOutput.toJson(newPayload);\n" +
                "    return newJsonPayload;\n" +
                "};\n" +
                "this";
        datahiveConsumer.startConsumerWithTransformations("power_grid_data.in", groovyScript, "power_grid_data.out");
        try {
            datahiveConsumer.startConsumerAndPushHadoop("power_grid_data.out", "/data/random.json");
        }
        catch(Exception e) {
        }
    }

}
