import React, { useEffect, useState } from "react";
import Swal from "sweetalert2";
import { NavbarUsuario } from "../components/NavbarUsuario";
import { CursoCard } from "../components/CursoCard";
import "./Reservas.css";

export function Reservas() {
  const [hoteles, setHoteles] = useState([]); // todos los hoteles del backend
  const [filteredHoteles, setFilteredHoteles] = useState([]); // hoteles filtrados

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
        Swal.fire({
          icon: "error",
          title: "Error al cargar habitaciones",
          text: "No pudimos obtener la lista de habitaciones. Intentalo nuevamente.",
          confirmButtonColor: "#00bcd4",
        });
      }
    };

    fetchHoteles();
  }, []);

  const handleBuscar = (e) => {
    e.preventDefault();

    // Validaciones: todos los campos obligatorios
    if (!ubicacion || !fechaEntrada || !fechaSalida) {
      Swal.fire({
        icon: "warning",
        title: "Datos incompletos",
        text: "Debés completar ubicación, fecha de entrada y fecha de salida.",
        confirmButtonColor: "#00bcd4",
      });
      return;
    }

    const entradaDate = new Date(fechaEntrada);
    const salidaDate = new Date(fechaSalida);

    if (salidaDate <= entradaDate) {
      Swal.fire({
        icon: "error",
        title: "Rango de fechas inválido",
        text: "La fecha de salida debe ser posterior a la fecha de entrada.",
        confirmButtonColor: "#00bcd4",
      });
      return;
    }

    // Por ahora filtramos en frontend por título (luego será search-api)
    const filtrados = hoteles.filter((hotel) =>
      hotel.titulo?.toLowerCase().includes(ubicacion.toLowerCase())
    );

    setFilteredHoteles(filtrados);

    if (filtrados.length === 0) {
      Swal.fire({
        icon: "info",
        title: "Sin resultados",
        text: "No se encontraron habitaciones para la búsqueda realizada.",
        confirmButtonColor: "#00bcd4",
      });
    } else {
      Swal.fire({
        icon: "success",
        title: "Búsqueda actualizada",
        text: "Te mostramos las habitaciones disponibles según tu búsqueda.",
        timer: 1800,
        showConfirmButton: false,
      });
    }

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
          {/* Barra de búsqueda */}
          <form className="filtros-reservas" onSubmit={handleBuscar}>
            <div className="campo-filtro ubicacion">
              <label className="campo-label">Ubicación</label>
              <input
                type="text"
                className="campo-input"
                placeholder="¿Dónde querés hospedarte?"
                value={ubicacion}
                onChange={(e) => setUbicacion(e.target.value)}
              />
            </div>

            <div className="campo-filtro entrada">
              <label className="campo-label">Entrada</label>
              <input
                type="date"
                className="campo-input"
                value={fechaEntrada}
                onChange={(e) => setFechaEntrada(e.target.value)}
              />
            </div>

            <div className="campo-filtro salida">
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
