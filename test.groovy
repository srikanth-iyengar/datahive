import groovy.json.JsonSlurper
import groovy.json.JsonOutput
def transform(record) {
	def jsonSlurper = new JsonSlurper();
	def jsonObject  = jsonSlurper.parseText(record);
	jsonObject.sum = jsonObject.a + jsonObject.b;
    return JsonOutput.toJson(jsonObject);
};

println transform('{"a": 1, "b": 2}')
