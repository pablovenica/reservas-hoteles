import React from "react";
import { NavbarUsuario } from "../components/NavbarUsuario";
import { CursoCard } from "../components/CursoCard";
import "./Reservas.css";

export function Reservas() {
  const habitaciones = [
    {
      id: 1,
      titulo: "Suite Ejecutiva",
      nivel: "Capacidad: 2 personas",
      duracion: "Vista al mar",
      docente: "Incluye desayuno y spa",
      precio: "$85.000 / noche",
    },
    {
      id: 2,
      titulo: "Habitación Doble Deluxe",
      nivel: "Capacidad: 3 personas",
      duracion: "Vista a la ciudad",
      docente: "Incluye desayuno buffet",
      precio: "$65.000 / noche",
    },
    {
      id: 3,
      titulo: "Suite Familiar Premium",
      nivel: "Capacidad: 5 personas",
      duracion: "Vista panorámica",
      docente: "Incluye desayuno y piscina",
      precio: "$110.000 / noche",
    },
  ];

  return (
    <>
      <NavbarUsuario />
      <section className="reservas-section">
        <div className="reservas-container">
          <h2 className="titulo-principal">Nuestras Habitaciones</h2>
          <div className="cards-container">
            {habitaciones.map((habitacion) => (
              <CursoCard key={habitacion.id} {...habitacion} />
            ))}
          </div>
        </div>
      </section>
    </>
  );
}
