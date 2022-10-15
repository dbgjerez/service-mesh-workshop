import AppInfo from "../components/info/AppInfo";

const connections = [
    {
        "service": "Prueba",
        "url": "http://localhost:8080/api/v1/info"
    },
]

const Info = () => {
    return connections.map(
        conn => <div><AppInfo conn={conn} /></div>
    )
};

export default Info;