from elasticsearch import Elasticsearch
import pandas as pd
import datetime

# Elasticsearch configuration
es = Elasticsearch([{'host': 'localhost', 'port': 9200, 'scheme': 'http'}])  # Replace with your Elasticsearch server details

# Define the Elasticsearch index and document type
index_name = "datahive_2"  # Replace with your desired index name

# Define the Elasticsearch mapping
mapping = {
    "mappings": {
        "properties": {
            "Order_Date": {"type": "date", "format": "yyyy-MM-dd"},
            "Time": {"type": "date", "format": "HH:mm:ss"},
            "Aging": {"type": "float"},
            "Customer_Id": {"type": "integer"},
            "Gender": {"type": "keyword"},
            "Device_Type": {"type": "keyword"},
            "Customer_Login_type": {"type": "keyword"},
            "Product_Category": {"type": "keyword"},
            "Product": {"type": "keyword"},
            "Sales": {"type": "float"},
            "Quantity": {"type": "integer"},
            "Discount": {"type": "float"},
            "Profit": {"type": "float"},
            "Shipping_Cost": {"type": "float"},
            "Order_Priority": {"type": "keyword"},
            "Payment_method": {"type": "keyword"}
        }
    }
}

# Read the CSV file
csv_file = "dataset.csv"  # Replace with the path to your CSV file
df = pd.read_csv(csv_file)

# Create the Elasticsearch index with the defined mapping
# es.indices.create(index=index_name, body=mapping)

# Convert DataFrame rows to Elasticsearch documents
id = 100000
for _, row in df.iterrows():
    try:
        document = row.to_dict()
        document['Order_date'] = f"{datetime.date.today()}"
        document_id = f"{id}"  # Use a unique ID for each document

    # Index the document into Elasticsearch
        es.index(index=index_name, id=document_id, body=document)
        id += 1
        print(id)
    except:
        continue

print(f"Successfully indexed {len(df)} documents into Elasticsearch.")

