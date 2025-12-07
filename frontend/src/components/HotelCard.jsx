import React from "react";
import { Link } from "react-router-dom";
import "./HotelCard.css";

export function HotelCard({ id, nombre, imagen, descripcion, provincia, precio }) {
  return (
    <div className="hotel-card">
      <img src={imagen} alt={nombre} className="hotel-img" />

      <div className="hotel-body">
        <h3>{nombre}</h3>

        <p className="provincia"><strong>{provincia}</strong></p>
        <p className="descripcion">{descripcion}</p>

        <p className="precio">USD {precio}</p>

        <Link to={`/usuario/hotel/${id}`} className="btn-detalle">
          Ver detalle
        </Link>
      </div>
    </div>
  );
}
