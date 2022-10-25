import React from 'react'
import config from "../config/app.film.crud.json"
import FilmCRUD from '../components/crud/FilmCRUD';

const Film = () => {
    return <FilmCRUD config={config}/> 
};

export default Film;