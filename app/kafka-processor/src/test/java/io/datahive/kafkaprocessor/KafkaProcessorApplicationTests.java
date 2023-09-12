package io.datahive.kafkaprocessor;

import io.datahive.kafkaprocessor.worker.DatahiveConsumer;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import javax.jms.*;

@SpringBootTest
class KafkaProcessorApplicationTests {

	@Autowired
	private DatahiveConsumer datahiveConsumer;

	@Test
	void contextLoads() {
	}


//	@Test
//	void consumerWithTransformations() {
//		String groovyScript = "import groovy.json.JsonSlurper\n" +
//				              "def transform(record) {\n"
//							+ "	def jsonSlurper = new JsonSlurper();\n"
//							+ "	def jsonObject  = jsonSlurper.parseText(record);\n"
//							+ "	jsonObject.sum = jsonObject.a + jsonObject.b;\n"
//							+ " return jsonObject;\n"
//							+ "};"
//				;
//		try {
//			datahiveConsumer.startConsumerWithTransformations("test-in-topic", groovyScript, "test-out-topic");
//			Thread.sleep(2000);
//			Session session = connection.createSession(false, Session.AUTO_ACKNOWLEDGE);
//			MessageProducer producer = session.createProducer(session.createTopic("test-in-topic"));
//			TextMessage msg1 = session.createTextMessage();
//			TextMessage msg2 = session.createTextMessage();
//			msg2.setText("{\"a\": 1, \"b\": 2 }");
//			msg1.setText("{\"a\": 1, \"b\": 2 }");
//			producer.send(msg1);
//			producer.send(msg2);
//		}
//		catch (Exception e) {
//			e.printStackTrace();
//		}
//	}

}
