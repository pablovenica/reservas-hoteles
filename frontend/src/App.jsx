import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";

import { Home } from "./pages/Home";
import { Usuario } from "./pages/Usuario";
import { Reservas } from "./pages/Reservas";
import { Admin } from "./pages/Admin";
import { Login } from "./pages/Login";
import { Registro } from "./pages/Registro";

import "./App.css";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/usuario" element={<Usuario />} />
        <Route path="/usuario/reservas" element={<Reservas />} />
        <Route path="/admin" element={<Admin />} />
        <Route path="/login" element={<Login />} />
        <Route path="/registro" element={<Registro />} />
        <Route
          path="*"
          element={
            <h1 style={{ textAlign: "center", marginTop: "50px" }}>
              Página no encontrada
            </h1>
          }
        />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
