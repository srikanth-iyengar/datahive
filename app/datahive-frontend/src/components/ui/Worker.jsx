import axios from "axios"
import { useEffect, useState } from "react"

const Worker = ({id="1"}) => {
    const [data, setData] = useState(0)
    const [loading, setLoading] = useState(true)
    const [success, setSuccess] = useState(false)
    useEffect(()=>{
        axios.get(`http://localhost:3000/worker/${id}`).then((resp)=>{
            let running = 0
            resp.data.map((work)=>{
                running += work.status === "RUNNING" ? 1 : 0
            })
            setData(running+"/"+resp.data.length)
            setSuccess(running === resp.data.length)
            setLoading(false)
        })
    }, [])
    return (
        <div>
            {loading ? 
            <div className="w-10 h-10 loader animate-spin border-t-2 rounded-full"></div> : 
            <h5>{data} {
                success ?
                <h5>✅</h5>
                :
                <h5>❌</h5>
            }</h5>}
        </div>
    )
}

export default Worker