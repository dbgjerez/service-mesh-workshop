import { useEffect, useState } from "react";
import { Grid, Header, Icon } from "semantic-ui-react";
//import './appinfo.scss'

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

    if (!isLoaded){
        return <div>Loading...</div>;
    } else {
        return (
        <Grid celled>
            <Grid.Row verticalAlign='middle'>
                <Grid.Column width={2}>
                    <Header as='h4' icon textAlign='center'>
                        <Icon name={error ? 'thumbs down outline' : 'thumbs up outline'} />
                    </Header>
                </Grid.Column>
                <Grid.Column width={3}>
                    <Header as='h2' icon textAlign='center'>
                        <Header.Content>{error ? conn.conn.service : data.app.name}</Header.Content>
                    </Header>
                </Grid.Column>
                <Grid.Column width={6}>
                    <Header as='h2' icon textAlign='left'>
                        <Header.Content>{error ? error.message : data.app.version}</Header.Content>
                    </Header>
                </Grid.Column>
            </Grid.Row>
        </Grid>
        );
    } 
};

export default AppInfo;