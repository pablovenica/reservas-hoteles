import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";

// === PAGES ===
import { Home } from "./pages/Home";
import { Usuario } from "./pages/Usuario";
import { Reservas } from "./pages/Reservas";
import { Admin } from "./pages/Admin";
import { Login } from "./pages/Login";
import { Registro } from "./pages/Registro";

// === ESTILOS ===
import "./App.css";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        {/* === PÁGINA PRINCIPAL === */}
        <Route path="/" element={<Home />} />

        {/* === SECCIÓN USUARIO === */}
        <Route path="/usuario" element={<Usuario />} />
        <Route path="/usuario/reservas" element={<Reservas />} />

        {/* === PANEL ADMIN === */}
        <Route path="/admin" element={<Admin />} />

        {/* === LOGIN === */}
        <Route path="/login" element={<Login />} />

        {/* === REGISTRO === */}
        <Route path="/registro" element={<Registro />} />

        {/* === ERROR 404 === */}
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
