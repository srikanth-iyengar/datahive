import groovy.json.JsonSlurper
import groovy.json.JsonOutput
def transform(record) {
	def jsonSlurper = new JsonSlurper();
	def kafkaData = jsonSlurper.parseText(record);
    def totalVoltage = 0.0;
    def totalCurrent = 0.0;
    def totalPower = 0.0;
    def gridCount = 0;

    def stationName = kafkaData.station.name
    println kafkaData[stationName][0]
	kafkaData[stationName].each {gridData ->
	        totalVoltage += gridData.distributor_grid.voltage;
	        totalCurrent += gridData.distributor_grid.current;
	        totalPower += gridData.distributor_grid.power;
	        gridCount++;
	};

    def averageVoltage = totalVoltage / gridCount;
    def averageCurrent = totalCurrent / gridCount;
    def averagePower = totalPower / gridCount;

    def newPayload = [
        timestamp: kafkaData['timestamp'],
        station_details: kafkaData['station'],
        average_metrics: [
            voltage: averageVoltage,
            current: averageCurrent,
            power: averagePower
        ]
    ];

    def newJsonPayload = JsonOutput.toJson(newPayload);
    return newJsonPayload;
};

println transform('''
{
    "timestamp": "2023-09-13 22:40:09",
	"station": {
		"name": "Station A",
		"location": "City X"
	},
	"Station A": [ 
        {
			"power_station": {
				"name": "Station A",
				"location": "City X"
			},
			"distributor_grid": {
				"grid_id": 1,
				"voltage": 213.02,
				"current": 28.31,
				"power": 5761.8
			},
			"timestamp": "2023-09-13 22:40:09"
		},
        {
			"power_station": {
				"name": "Station A",
				"location": "City X"
			},
			"distributor_grid": {
				"grid_id": 1,
				"voltage": 213.02,
				"current": 28.31,
				"power": 5761.8
			},
			"timestamp": "2023-09-13 22:40:09"
		}
    ],
}
''')
