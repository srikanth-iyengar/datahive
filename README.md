# Datahive

**Datahive** is an ingenious, configuration-driven end-to-end data pipeline solution that simplifies the complexities of managing data workflows. Harnessing the power of Kafka, Hadoop, Apache Spark, Elasticsearch, Kibana, and an intuitive UI, Datahive empowers users to effortlessly manage and monitor their data stacks.

## Features

- 🚀 Streamlined data pipeline setup
- ☕ Automated data processing while you enjoy your coffee
- 📊 Utilizes Kafka, Hadoop, Apache Spark, Elasticsearch, and Kibana
- 🛠️ Easy configuration through YAML files
- 🔄 Supports both stream and batch processing

## How it Works

Define your data pipeline effortlessly using a simple YAML configuration file. Specify input and output schemas for each service, and let Datahive handle the rest. Below is a sample configuration for stream processing:

```yaml
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
      hdfsFileName: <your-hdfs-filename>
      hdfs: true

spark:
    - app-resource: <path-for-your-spark-build-file>
      driver.memory: 1g
      executor.memory: 2g
    - app-resource: <path-for-your-second-spark-build-file>
      driver-memory: 1g
      executor-memory: 2g
      res-location: <path-for-the-spark-job-code>
      main-class: <main-class-of-your-spark-job>
      job-name: <name-of-your-job>

elasticsearch:
    - 

kibana:
    dashboard-config:
```


## Getting Started
1. Clone the repository.
2. Install the required dependencies.
3. Configure Datahive using the provided YAML files.
4. Run the application.

## Screenshots
![image](https://github.com/srikanth-iyengar/datahive/assets/88551109/771a59ad-9665-411b-9cec-5eb226d21d2f)

![image](https://github.com/srikanth-iyengar/datahive/assets/88551109/fb2d4c6e-8666-45bb-95ea-60c12b318a0c)
![image](https://github.com/srikanth-iyengar/datahive/assets/88551109/60d45183-4a54-490b-93ae-bcdc170ffee5)
![image](https://github.com/srikanth-iyengar/datahive/assets/88551109/b4ba5a42-6b29-4934-aa58-d1c4755aba37)
![image](https://github.com/srikanth-iyengar/datahive/assets/88551109/7fabf218-ed1c-4ec1-a991-e0893d52db11)
![image](https://github.com/srikanth-iyengar/datahive/assets/88551109/615c2510-e447-40db-9360-2b4bd13db4b4)
![image](https://github.com/srikanth-iyengar/datahive/assets/88551109/396cf4ff-2d17-49c8-8897-717e45fef7eb)
![image](https://github.com/srikanth-iyengar/datahive/assets/88551109/f440a1ab-93c1-419c-bd1c-424ef6a03771)


