from kafka import KafkaProducer
import csv
import time

kafka_bootstrap_servers = 'kafka:9092'
kafka_topic = 'ecommerce-data.in'
csv_file_path = '/stream.csv'

producer = KafkaProducer(
    bootstrap_servers=[kafka_bootstrap_servers],
    value_serializer=lambda v: str(v).encode('utf-8')
)

while True:
    with open(csv_file_path, 'r') as csv_file:
        csv_reader = csv.reader(csv_file)
        for row in csv_reader:
            csv_string = ','.join(row)  # Convert row to comma-separated string
            print(csv_string)  # For debugging
            producer.send(kafka_topic, value=csv_string)
            time.sleep(1)

producer.close()

