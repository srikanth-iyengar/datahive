import groovy.json.JsonBuilder

def transform(record) {
    def header = ["Order_Date", "Time", "Aging", "Customer_Id", "Gender", "Device_Type", "Customer_Login_type", "Product_Category", "Product", "Sales", "Quantity", "Discount", "Profit", "Shipping_Cost", "Order_Priority", "Payment_method"]

    def values = record.split(',')

    def jsonObject = [:]

    for (int i = 0; i < header.size(); i++) {
        jsonObject[header[i]] = values[i].trim()
    }

    return new JsonBuilder(jsonObject).toPrettyString()
}

// Example usage:
def csvRecord = "2018-01-02,10:56:33,8.0,37077,Female,Web,Member,Auto & Accessories,Car Media Players,140.0,1.0,0.3,46.0,4.6,Medium,credit_card"

def result = transform(csvRecord)
println(result)

