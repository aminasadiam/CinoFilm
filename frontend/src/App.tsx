import React, { useEffect, useState } from 'react';
import './App.css';
import Register from './pages/Register';
import Login from './pages/Login';
import Home from './pages/Home';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Nav from './components/Nav';

function App() {
  const [username, setUsername] = useState("");

  useEffect (() => {
    (
      async () => {
        const response = await fetch("http://localhost:3650/api/user", {
          headers: {'Content-Type': 'application/json'},
          credentials: 'include',
        });

        const content = await response.json();

        setUsername(content.username)
      }
    )();
  });

  return (
    <div className="App">
        <BrowserRouter>
            <Nav name={username} setName={setUsername}/>

            <Routes>
                <Route path="/" element={<Home name={username}/>}/>
                <Route path="/login" element={<Login setName={setUsername}/>}/>
                <Route path="/register" element={<Register />}/>
            </Routes>
        </BrowserRouter>
    </div>
);
}

export default App;
