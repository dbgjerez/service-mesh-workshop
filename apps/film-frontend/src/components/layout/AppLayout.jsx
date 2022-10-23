import { Outlet } from "react-router-dom";
import { Grid, Header, Icon } from "semantic-ui-react";
import Sidebar from "../sidebar/Sidebar";

const AppLayout = () => {
    return (
        <Grid celled>
            <Grid.Row>
            <Header as='h2' icon textAlign='center'>
                <Icon name='book' circular />
                <Header.Content>Service Mesh Workshop</Header.Content>
            </Header>
            </Grid.Row>
            <Grid.Row>
                <Grid.Column width={3}>
                    <Sidebar/>
                </Grid.Column>
                <Grid.Column width={13}>
                    <Outlet />
                </Grid.Column>
            </Grid.Row>
        </Grid>
    )
};

export default AppLayout;