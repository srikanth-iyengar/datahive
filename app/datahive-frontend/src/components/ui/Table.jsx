import React, { useEffect, useState } from "react"
import axios from "axios"
import Worker from "./Worker"

const Table = ({ header = [] }) => {
    const [rows, setRows] = useState([])
    const [configs, setConfigs] = useState({})
    const [worker, setWorker] = useState({})
    const boof = [1, 1, 1, 1, 1, 1]
    const [loading, setLoading] = useState(true)
    const [open, setOpen] = useState(0)
    useEffect(() => {
        axios.get("http://localhost:3000/pipeline/all").then((resp) => {
            const data = resp.data
            const rs = []
            const confs = {}
            data.map((r) => {
                confs[r.id] = r.configuration
                rs.push([
                    r.id,
                    r.name,
                    "Admin",
                    <button className="button">View</button>,
                    <Worker id={r.id} />,
                    <div className="flex flex-row gap-4">
                        <button id="modal-trigger" onClick={() => {
                            setOpen(r.id)
                            setWorker({ 2: 3 })
                            console.log(worker)
                        }} className="button color-secondary">Edit</button>
                        <button className="button color-primary">Delete</button>
                    </div>
                ])
            })
            setConfigs(confs)
            setRows(rs)
        })
    }, [])
    useEffect(() => {
        rows.map((r) => {
            const work = {}
            axios.get(`http://localhost:3000/worker/${r[0]}`).then((resp) => {
                let running = 0;
                resp.data.map((w) => {
                    if (w.status === "RUNNING") {
                        running += 1;
                    }
                })
                work[r[0]] = running
            })
            console.log(work)
            console.log(work)
            setWorker(work)
            setLoading(false)
        })
    }, [rows])
    return (<div className="relative shadow-md sm:rounded-lg w-11/12 my-4">
        <table className="w-full text-sm text-left">
            <thead className="text-xs uppercase">
                <tr className="text-center border-b">
                    {header.map((head) => (
                        <th scope="col" className="px-6 py-3 font-semibold text-xl">
                            <span> {head} </span>
                        </th>
                    ))}
                </tr>
            </thead>
            <tbody>
                {
                    rows.length === 0
                        ?
                        boof.map((val) => (
                            <tr className="animate-pulse">
                                {boof.map((val1) => (
                                    <td className="px-6 py-4 text-xl text-center">
                                        <div className="h-10 w-full bg-gray-700 mb-6 rounded-full"></div>
                                    </td>
                                ))}
                            </tr>
                        ))
                        :
                        rows.map((row) => (
                            <tr className="border-b">
                                {row.map((col) => (
                                    <td className="px-6 py-4 text-xl text-center">
                                        <span> {col} </span>
                                    </td>
                                ))}
                            </tr>
                        ))
                }
            </tbody>
        </table>
    </div>)
}

export default Table