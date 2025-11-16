import React, { useEffect, useState } from "react";
import { NavbarUsuario } from "../components/NavbarUsuario";
import { CursoCard } from "../components/CursoCard";
import "./Reservas.css";

export function Reservas() {
  const [hoteles, setHoteles] = useState([]);            // todos los hoteles del backend
  const [filteredHoteles, setFilteredHoteles] = useState([]); // hoteles filtrados para mostrar

  const [ubicacion, setUbicacion] = useState("");
  const [fechaEntrada, setFechaEntrada] = useState("");
  const [fechaSalida, setFechaSalida] = useState("");

  // Cargar hoteles desde booking-api al montar
  useEffect(() => {
    const fetchHoteles = async () => {
      try {
        const res = await fetch("http://localhost:8082/hotels");
        const data = await res.json();
        console.log("Hoteles desde backend:", data);
        setHoteles(data);
        setFilteredHoteles(data); // por defecto, mostrar todos
      } catch (error) {
        console.log("Error:", error);
      }
    };

    fetchHoteles();
  }, []);

  const handleBuscar = (e) => {
    e.preventDefault();

    // Validaciones
    if (!ubicacion || !fechaEntrada || !fechaSalida) {
      alert("Debés completar ubicación, fecha de entrada y fecha de salida.");
      return;
    }

    const entradaDate = new Date(fechaEntrada);
    const salidaDate = new Date(fechaSalida);

    if (salidaDate <= entradaDate) {
      alert("La fecha de salida debe ser posterior a la fecha de entrada.");
      return;
    }

    // más adelante vamos a llamar al search-api.
    // Por ahora hacemos un filtro simple por título para que veas el comportamiento.
    const filtrados = hoteles.filter((hotel) =>
      hotel.titulo?.toLowerCase().includes(ubicacion.toLowerCase())
    );

    setFilteredHoteles(filtrados);

    console.log("Payload preparado para search-api:", {
      ubicacion,
      fechaEntrada,
      fechaSalida,
    });
  };

  return (
    <>
      <NavbarUsuario />
      <section className="reservas-section">
        <div className="reservas-container">

          {/* 🔹 Barra de búsqueda (ubicación + fechas) */}
          <form className="filtros-reservas" onSubmit={handleBuscar}>
            <div className="campo-filtro">
              <label className="campo-label">Ubicación</label>
              <input
                type="text"
                className="campo-input"
                placeholder="¿Dónde querés hospedarte?"
                value={ubicacion}
                onChange={(e) => setUbicacion(e.target.value)}
              />
            </div>

            <div className="campo-filtro">
              <label className="campo-label">Entrada</label>
              <input
                type="date"
                className="campo-input"
                value={fechaEntrada}
                onChange={(e) => setFechaEntrada(e.target.value)}
              />
            </div>

            <div className="campo-filtro">
              <label className="campo-label">Salida</label>
              <input
                type="date"
                className="campo-input"
                value={fechaSalida}
                onChange={(e) => setFechaSalida(e.target.value)}
              />
            </div>

            <button type="submit" className="btn-buscar">
              Buscar
            </button>
          </form>

          <h2 className="titulo-principal">Nuestras Habitaciones</h2>

          <div className="cards-container">
            {filteredHoteles.length === 0 ? (
              <p className="mensaje-sin-resultados">
                No se encontraron habitaciones para la búsqueda realizada.
              </p>
            ) : (
              filteredHoteles.map((habitacion) => (
                <CursoCard
                  key={habitacion.id}
                  titulo={habitacion.titulo}
                  nivel={`Nivel: ${habitacion.nivel}`}
                  duracion={habitacion.duracion}
                  docente="Hotel"
                  precio={habitacion.precio}
                />
              ))
            )}
          </div>
        </div>
      </section>
    </>
  );
}
