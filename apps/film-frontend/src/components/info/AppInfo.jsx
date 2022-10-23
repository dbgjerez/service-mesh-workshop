import { useEffect, useState } from "react";
import './appinfo.scss'

export const Timeout = (time) => {
	let controller = new AbortController();
	setTimeout(() => controller.abort(), time * 1000);
	return controller;
};

const AppInfo = (conn) => {

    const [error, setError] = useState(null)
    const [isLoaded, setIsLoaded] = useState(false);
    const [data, setData] = useState([]);

    useEffect(() => {
        fetch(conn.conn.url, {
            signal: Timeout(1).signal
        })
            .then(res => res.json())
            .then(
                (result) => (
                    setIsLoaded(true),
                    setData(result)
                ),
                (error) => (
                    setIsLoaded(true),
                    setError(error)
                )
            )
    }, [])

    if (error){
        return (
            <div className="app">
                <p className='app__name'>{conn.conn.service}</p>
                <p className='app__version'>{error.message}</p>
            </div>
            );
    } else if (!isLoaded){
        return <div>Loading...</div>;   
    } else {
        return (
            <div className="app">
                <p className='app__name'>{data.app.name}</p>
                <p className='app__version'>{data.app.version}</p>
            </div>
            );
    }
};

export default AppInfo;