import AppInfo from "../components/info/AppInfo";

const apps = [{
    "name": "Prueba",
    "version": "1.1"
},{
    "name": "Prueba",
    "version": "1.2"
},]

const Info = () => {
    return apps.map(
        app => <div><AppInfo app={app} /></div>
    )
};

export default Info;