import React, { useEffect, useState } from "react";
import { NavbarUsuario } from "../components/NavbarUsuario";
import { CursoCard } from "../components/CursoCard";
import "./Reservas.css";

export function Reservas() {
  const [hoteles, setHoteles] = useState([]);

  useEffect(() => {
    const fetchHoteles = async () => {
      try {
        const res = await fetch("http://localhost:8082/hotels");
        const data = await res.json();
        setHoteles(data);
      } catch (error) {
        console.log("Error:", error);
      }
    };

    fetchHoteles();
  }, []);

  return (
    <>
      <NavbarUsuario />
      <section className="reservas-section">
        <div className="reservas-container">
          <h2 className="titulo-principal">Nuestras Habitaciones</h2>
          <div className="cards-container">
            {hoteles.map((hotel) => (
              <CursoCard
                key={hotel.id}
                titulo={hotel.titulo}
                nivel={`Nivel: ${hotel.nivel}`}
                duracion={hotel.duracion}
                docente="Hotel"
                precio={`$${hotel.precio}`}
              />
            ))}
          </div>
        </div>
      </section>
    </>
  );
}
