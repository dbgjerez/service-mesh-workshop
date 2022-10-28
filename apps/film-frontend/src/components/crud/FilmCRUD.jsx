import { useEffect, useState } from "react";
import { Button, Form, Grid, Icon, Table } from 'semantic-ui-react'

const FilmCRUD = (config) => {
    const [error, setError] = useState(null)
    const [isLoaded, setIsLoaded] = useState(false);
    const [data, setData] = useState([]);
    const [item, setItem] = useState({});

    const handleTitle = (e) => {
        setItem({
            ...item,
            title: e.target.value
        })
    }
    const handleDuration = (e) => {
        setItem({
            ...item,
            duration: e.target.value
        })
    }
    const handlePremium = (e) => {
        setItem({
            ...item,
            premium: e.target.checked 
        })
    }

    const updateFilm = () => {
        console.log(item)
    }

    const findById = (e) => {
        console.log(e.target.dataset.id)
        fetch(`http://localhost:8080/api/v1/film/`+e.target.dataset.id)
            .then(
                (res) => {
                    if(!res.ok) throw new Error(res.status)
                    else return res.json()
                }
            ).then(
                (result) => (
                    setItem(result)
                ),
                (error) => (
                    setItem(error)
                )
            )
    }

    const deleteById = (e) => {
        console.log(e.target.dataset.id)
        fetch(`http://localhost:8080/api/v1/film/`+e.target.dataset.id, { method: 'DELETE' })
            .then(updateScreen)
    }

const updateScreen = () => {
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
}

    useEffect(() => {
        updateScreen()
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
                                <Table.HeaderCell>Options</Table.HeaderCell>
                            </Table.Row>
                        </Table.Header>
                        <Table.Body>
                            {
                                isLoaded ? 
                                    data.map((data)=> {
                                        return (
                                            <Table.Row key={data.id}>
                                                <Table.Cell>{data.title}</Table.Cell>
                                                <Table.Cell>{data.efficacy}</Table.Cell>
                                                <Table.Cell>{data.duration}</Table.Cell>
                                                <Table.Cell>{data.premium ? "SI" : "NO"}</Table.Cell>
                                                <Table.Cell>
                                                    <Icon name="edit" data-id={data.id} onClick={ findById } />
                                                    <Icon name="delete" data-id={data.id} onClick={ deleteById } />
                                                </Table.Cell>
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
                    <Form onSubmit={updateFilm}>
                        <Form.Input
                            icon='pencil alternate'
                            iconPosition='left'
                            label='Title'
                            placeholder='Movie title'
                            value={item.title}
                            onChange={handleTitle}
                            />
                        <Form.Input 
                            icon='film'
                            iconPosition='left'
                            label='Duration'
                            placeholder='Minutes'
                            type='number'
                            value={item.duration}
                            onChange={handleDuration}
                        />
                        <Form.Input
                            icon='trophy'
                            iconPosition='left'
                            type='checkbox'
                            label='Premium'
                            checked={item.premium}
                            onChange={handlePremium}
                        />
                        <Form.Button type='submit'>{item.id ? 'Update' : 'Create'}</Form.Button>
                        <Button onClick={setItem}>Clear</Button>
                    </Form>
                </Grid.Column>
            </Grid.Row>
        </Grid>
    );
}

export default FilmCRUD;