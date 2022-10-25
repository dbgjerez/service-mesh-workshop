import AppInfo from "../components/info/AppInfo";
import connections from "../config/app.info.config.json"

const Info = () => {
    return (
        connections.map(
            conn => <AppInfo conn={conn}/> 
        )
    )
};

export default Info;