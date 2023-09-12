# Datahive 
Datahive is a configuration-driven end-to-end data pipeline solution
Datahive utilizes Kafka, Hadoop, Apache Spark, Elasticsearch, Kibana and a UI for getting the status of all the stacks

You define your configuration as by defining resource input and output schema for each service and you drink coffee meanwhile datahive does the job for you

```yaml
# sample datahive configuration
kafka:
    in-topic-name: <your-topic-name>
    out-topic-name: <your-topic-name>
    transform: | 
        def transform(record) {
            def jsonObject = record
            // do your transformation logic in a groovy script
            return jsonObject
        }
    send-data-to:
        - hadoop:
            hdfs-file-name: <hdfs-file-name>
        - elasticsearch:
            index-name: <elasticsearch-index>
# your elasticsearch mapping json value
            mappings: |
            {
                
            }
```
