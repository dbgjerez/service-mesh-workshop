import './appinfo.scss'

const AppInfo = (data) => {
    return <div className="app">
        <p className='app__name'>{data.app.name}</p>
        <p className='app__version'>{data.app.version}</p>
    </div>;
};

export default AppInfo;