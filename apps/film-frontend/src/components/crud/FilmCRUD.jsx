import { useEffect, useState } from "react";
import { Button, Checkbox, Form, Grid, Table } from 'semantic-ui-react'

const FilmCRUD = (config) => {
    const [error, setError] = useState(null)
    const [isLoaded, setIsLoaded] = useState(false);
    const [data, setData] = useState([]);

    useEffect(() => {
        fetch(`http://localhost:8080/api/v1/film`)  
            .then(
                (res) => {
                    if(!res.ok) throw new Error(res.status)
                    else return res.json()
                }
            ).then(
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

    return (
        <Grid celled>
            <Grid.Row>
                <Grid.Column width={11}>
                    <Table celled padded>
                        <Table.Header>
                            <Table.Row>
                                <Table.HeaderCell singleLine>Movie title</Table.HeaderCell>
                                <Table.HeaderCell>Efficacy</Table.HeaderCell>
                                <Table.HeaderCell>Duration</Table.HeaderCell>
                                <Table.HeaderCell>Premium</Table.HeaderCell>
                            </Table.Row>
                        </Table.Header>
                        <Table.Body>
                            {
                                isLoaded ? 
                                    data.map((data)=> {
                                        return (
                                            <Table.Row>
                                                <Table.Cell>{data.title}</Table.Cell>
                                                <Table.Cell>{data.efficacy}</Table.Cell>
                                                <Table.Cell>{data.duration}</Table.Cell>
                                                <Table.Cell>{data.premium ? "SI" : "NO"}</Table.Cell>
                                            </Table.Row>
                                        )
                                    }) : 
                                    (
                                        <Table.Row>
                                            <Table.Cell></Table.Cell>
                                            <Table.Cell></Table.Cell>
                                            <Table.Cell></Table.Cell>
                                            <Table.Cell></Table.Cell>
                                        </Table.Row>
                                    )
                            }
                        </Table.Body>
                    </Table>
                </Grid.Column>
                <Grid.Column width={5}>
                    <Form>
                        <Form.Input
                            icon='pencil alternate'
                            iconPosition='left'
                            label='Title'
                            placeholder='Movie title'
                            />
                        <Form.Input 
                            icon='film'
                            iconPosition='left'
                            label='Duration'
                            placeholder='Minutes'
                            type='number'
                        />
                        <Form.Input
                            icon='trophy'
                            iconPosition='left'
                            type='checkbox'
                            label='Premium'
                        />
                        <Button type='submit'>Submit</Button>
                    </Form>
                </Grid.Column>
            </Grid.Row>
        </Grid>
    );
}

export default FilmCRUD;