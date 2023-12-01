const Node = ({ appName = "Hadoop", url = "hadoop.srikanthk.tech", isUp = true, img = "/hadoop.svg", uptime = "1D" }) => {
    return (
        <div className="container">
            <div className="flex-row flex gap-5">
                {isUp  ?
                <img
                    src={img}
                    className={`w-14 h-14 rounded-full ring-2 ring-green-500 shadow-md shadow-green-500 backdrop-blur-sm border-white`}
                />
                :
                <img
                    src={img}
                    className={`w-14 h-14 rounded-full ring-2 ring-red-500 shadow-md shadow-red-500 backdrop-blur-sm border-white`}
                />
                }
                <div className="grid grid-cols-2 gap-4">
                    <div>
                        <p className="font-semibold">{appName.substring(0, 12)}..</p>
                        <p>Host: {url.substring(0, 7)}...</p>
                    </div>
                    <p className="my-auto font-extrabold text-4xl">{uptime}</p>
                </div>
            </div>
        </div>
    )
}

export default Node