import React from "react";
import "./Home.css";
import { RoleCard } from "../components/RoleCard";

import imgUsuario from "../assets/img-usuario.png";
import imgAdmin from "../assets/img-admin.png";

export function Home() {
  return (
    <div className="home-container">
      <div className="home-content">
        <h1 className="home-title">Sistema de Reservas para un Hotel</h1>
        <p className="home-subtitle">Elegí tu rol para continuar:</p>

        <div className="cards-container">
          <RoleCard
            img={imgUsuario}
            titulo="Usuario"
            link="/usuario"
            gradiente="verde"
          />
          <RoleCard
            img={imgAdmin}
            titulo="Administrador"
            link="/admin"
            gradiente="rojo"
          />
        </div>
      </div>
    </div>
  );
}
