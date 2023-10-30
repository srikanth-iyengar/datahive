const Sidebar = () => {
    return (
        <div className="inner__brand">
          <div className="bg-neutral-900 radius-large space-8">
            <img
              className={"brand__logo radius-large"}
              src={"/favicon.svg"}
              alt=""
            style={{
                backgroundColor: "#f5f5f5",
            }}
            width={200}
            />
          </div>
          <div className="brand__text">
            <span>Datahive Dashboard</span>
            <span>Manage all your mission critical Data Pipelines</span>
          </div>
        </div>
    )
}

export default Sidebar