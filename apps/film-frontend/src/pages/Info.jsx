import AppInfo from "../components/info/AppInfo";
import connections from "../config/app.error.config.json"

const Info = () => {
    return (
        connections.map(
            conn => <AppInfo conn={conn}/> 
        )
    )
};

export default Info;