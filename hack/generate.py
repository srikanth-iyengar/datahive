import random
import time
import json
from confluent_kafka import Producer

KAFKA_BROKER = 'kafka:9092'
KAFKA_TOPIC = 'power_grid_data.in'

POWER_STATIONS = [
    {"name": "Station A", "location": "City X"},
    {"name": "Station B", "location": "City Y"},
    {"name": "Station C", "location": "City Z"},
    {"name": "Station D", "location": "City W"},
    {"name": "Station E", "location": "City W"},
    {"name": "Station F", "location": "City U"},
    {"name": "Station G", "location": "City P"},
    {"name": "Station H", "location": "City Q"},
]

MIN_VOLTAGE = random.uniform(200, 220)
MAX_VOLTAGE = random.uniform(230, 250)
MIN_CURRENT = random.uniform(5, 15)
MAX_CURRENT = random.uniform(30, 60)
MIN_POWER = random.uniform(1000, 2000)
MAX_POWER = random.uniform(4000, 8000)
INTERVAL_SECONDS = 1
NUM_DISTRIBUTOR_GRIDS = 15


def generate_power_grid_data(producer):
    while True:
        for station in POWER_STATIONS:
            station_data = {}
            station_name = station["name"]
            station_location = station["location"]
            station_data[station_name] = []

            for grid_id in range(1, NUM_DISTRIBUTOR_GRIDS + 1):
                voltage = round(random.uniform(MIN_VOLTAGE, MAX_VOLTAGE), 2)
                current = round(random.uniform(MIN_CURRENT, MAX_CURRENT), 2)
                power = round(random.uniform(MIN_POWER, MAX_POWER), 2)
                timestamp = time.strftime("%Y-%m-%d %H:%M:%S")

                grid_data = {
                    "power_station": {
                        "name": station_name,
                        "location": station_location
                    },
                    "distributor_grid": {
                        "grid_id": grid_id,
                        "voltage": voltage,
                        "current": current,
                        "power": power
                    },
                    "timestamp": timestamp
                }

                station_data[station_name].append(grid_data)
            station_data['timestamp'] = timestamp
            station_data['station'] = station

            message = json.dumps(station_data)
            producer.produce(KAFKA_TOPIC, key="power_grid_data", value=message)
            producer.flush()
            print(json.dumps(station_data))
            time.sleep(INTERVAL_SECONDS)


if __name__ == "__main__":
    try:
        kafka_config = {
            'bootstrap.servers': KAFKA_BROKER,
            'client.id': 'power_grid_data_producer'
        }

        producer = Producer(kafka_config)

        generate_power_grid_data(producer)
    except KeyboardInterrupt:
        print("Data generation stopped.")

