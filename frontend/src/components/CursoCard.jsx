import React from "react";
import { Link } from "react-router-dom";
import "./CursoCard.css";

export function CursoCard({ id, titulo, nivel, duracion, docente, precio }) {
  return (
    <div className="curso-card">
      <div className="curso-header">
        <p className="categoria">Hotel & Reservas</p>
      </div>

      <h3>{titulo}</h3>

      <p>
        <strong>{nivel}</strong>
      </p>
      <p>
        <strong>{duracion}</strong>
      </p>
      <p>
        <strong>{docente}</strong>
      </p>

      <p className="precio">{precio}</p>

      <Link to={`/usuario/hotel/${id}`} className="btn-programa">
        Ver detalle
      </Link>
    </div>
  );
}
