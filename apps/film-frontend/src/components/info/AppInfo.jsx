import { useEffect, useState } from "react";
import './appinfo.scss'

const AppInfo = () => {

    const [error, setError] = useState(null)
    const [isLoaded, setIsLoaded] = useState(false);
    const [data, setData] = useState([]);

    useEffect(() => {
        fetch("http://localhost:8080/api/v1/info")
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
        return <div>Error: {error.message}</div>;
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