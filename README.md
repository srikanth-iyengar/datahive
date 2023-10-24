# Datahive 
Datahive is a configuration-driven end-to-end data pipeline solution
Datahive utilizes Kafka, Hadoop, Apache Spark, Elasticsearch, Kibana and a UI for getting the status of all the stacks

You define your configuration as by defining resource input and output schema for each service and you drink coffee meanwhile datahive does the job for you

```yaml
# sample datahive configuration
type: stream
kafka:
    - inTopic: <your-topic-name>
      outTopic: <your-topic-name>
      hdfs: false
      transform: | 
        def transform(record) {
            def jsonObject = record
            // do your transformation logic in a groovy script
            return jsonObject
        }
    - inTopic: <your-topic-name>
      hdfsFileName: <your-hdfs-filename> # the hdfs file path in which you want to save the file
      hdfs: true

spark:
    - 

elasticsearch:
    - 

kibana:
    dashboard-config:
```

```yaml
# sample datahive configuration for batch processing
```
