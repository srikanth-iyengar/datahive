import axios from "axios";
import { useEffect, useState } from "react";
import Node from "./Node";

function getHumanReadableDiff(timestamp1, timestamp2) {
    const diff = timestamp2 - timestamp1;
    const diffInSeconds = diff / 1000;
    const diffInMinutes = diffInSeconds / 60;
    const diffInHours = diffInMinutes / 60;
    const diffInDays = diffInHours / 24;
    let humanReadableDiff = "";
    function round(value, precision) {
        var multiplier = Math.pow(10, precision || 0);
        return Math.round(value * multiplier) / multiplier;
    }
    
    if (diffInDays >= 1) {
        humanReadableDiff += round(diffInDays, 1) + "D ";
    }
    else if (diffInHours >= 1) {
        humanReadableDiff += round(diffInHours, 1) + "H ";
    }
    else if (diffInMinutes >= 1) {
        humanReadableDiff += round(diffInMinutes, 1) + "M ";
    }
    else if (diffInSeconds >= 1) {
        humanReadableDiff += round(diffInSeconds, 1) + "S";
    }
    
    return humanReadableDiff;
}
const NodeGrid = () => {
    const [data, setData] = useState(Array(0))

    useEffect(() => {
        axios.get(`http://localhost:3000/health`).then((resp) => {
            resp.data.map((d)=>{
                if(d.name == "Kibana") {
                    d.img="/kibana.svg"
                }
                else if(d.name=="Elastic") {
                    d.img="/elastic.svg"
                }
                else if(d.name=="Apache Spark") {
                    d.img = "/spark.png"
                }
                else if(d.name=="MinIo") {
                    d.img="/minio.png"
                }
                else if(d.name=="Datahive Api"){
                    d.img="/api.svg"
                }
                else if(d.name =="Apache Hadoop") {
                    d.img="/hadoop.svg"
                }
                else if(d.name == "Apache Kafka") {
                    d.img ="/kafka.webp"
                }
                else if(d.name == "Datahive Ingestor") {
                    d.img = "/ingestor.svg"
                }
            })
            setData(resp.data)
            data.sort((a, b) => a.id - b.id)
        })
    }, [])

    return (
    <div className="my-10">
        <h5 className="text-4xl font-bold my-8">Stack Information</h5>
        <div className="grid grid-cols-3 gap-6">
            {data.map((st)=>(
                    <Node isUp={st.isUp} img={st.img} appName={st.name} uptime={getHumanReadableDiff(st.earliestSuccess, Math.floor(new Date().getTime() / 1000))} url={st.address} />
                )
            )}
        </div>
    </div>
    )
}

export default NodeGrid
