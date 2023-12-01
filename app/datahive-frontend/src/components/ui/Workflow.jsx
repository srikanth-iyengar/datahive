import axios from "axios";
import React, { useEffect, useState } from "react"
import Node from "./Node";

const Workflow = ({ id = "1" }) => {
    const [worker, setWorkers] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        axios.get(`http://localhost:3000/worker/${id}`).then((resp) => {
            resp.data.map((w)=>{
                w.isUp = w.status == "RUNNING"
                if(w.type === "KafkaConsumer") {
                    w.img = "/kafka.png"
                }
                else if(w.type === "KafkaConsumerWithHdfs") {
                    w.img = "/spark.png"
                }
                else if(w.type === "Elastic") {
                    w.img = "/elastic.png"
                }
            })
            setWorkers(resp.data)
        })
    }, [])

    return (
        <>
            <h5 className="font-semibold text-2xl my-10">Worker Details: {id}</h5>
            <div className="grid grid-cols-3 gap-6">
                {worker.map((work)=>(
                    <Node appName={work.id} img={work.img} uptime="" isUp={work.isUp} />
                ))}
                <Node appName="elastic-index" img="/elastic.svg" isUp={true} uptime="" />
                <Node appName="Kibana-dashboard" img="/kibana.svg" isUp={true} uptime=""/>
            </div>
        </>
    )
}

export default Workflow
