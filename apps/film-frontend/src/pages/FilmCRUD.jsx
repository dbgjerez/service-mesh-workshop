import React from 'react'
import { Button, Checkbox, Form, Grid, Segment } from 'semantic-ui-react'

const FilmCRUD = () => (
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
)

export default FilmCRUD