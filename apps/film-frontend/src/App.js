import './App.css';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Blank from './pages/Blank';
import AppLayout from './components/layout/AppLayout';
import Info from './pages/Info';
import Film from './pages/Film';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' element={<AppLayout />}>
          <Route index element={<Blank />} />
          <Route path='/films' element={<Film />} />
          <Route path='/info' element={<Info />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
