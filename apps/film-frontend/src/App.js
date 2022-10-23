import './App.css';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Blank from './pages/Blank';
import AppLayout from './components/layout/AppLayout';
import Info from './pages/Info';
import FilmCRUD from './pages/FilmCRUD';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' element={<AppLayout />}>
          <Route index element={<Blank />} />
          <Route path='/films' element={<FilmCRUD />} />
          <Route path='/info' element={<Info />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
