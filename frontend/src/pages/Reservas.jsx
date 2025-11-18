import React, { useEffect, useState } from "react";
import Swal from "sweetalert2";
import { NavbarUsuario } from "../components/NavbarUsuario";
import { CursoCard } from "../components/CursoCard";
import "./Reservas.css";

export function Reservas() {
  const [hoteles, setHoteles] = useState([]);
  const [filteredHoteles, setFilteredHoteles] = useState([]);

  const [ubicacion, setUbicacion] = useState("");
  const [fechaEntrada, setFechaEntrada] = useState("");
  const [fechaSalida, setFechaSalida] = useState("");

  useEffect(() => {
    const fetchHoteles = async () => {
      try {
        const res = await fetch("http://localhost:8082/hotels");
        const data = await res.json();
        console.log("➡️ Respuesta cruda de /hotels:", data);

        let lista = [];

        if (data && Array.isArray(data.hoteles)) {
          lista = data.hoteles;
        }

        else if (Array.isArray(data)) {
          lista = data;
        } else {
          lista = [];
        }

        console.log("Lista de hoteles:", lista);

        setHoteles(lista);
        setFilteredHoteles(lista);
      } catch (error) {
        console.error("Error al pedir /hotels:", error);
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

    const filtro = ubicacion.toLowerCase();

    const filtrados = hoteles.filter(
      (hotel) =>
        hotel.provincia?.toLowerCase().includes(filtro) ||
        hotel.nombre?.toLowerCase().includes(filtro)
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
  };

  return (
    <>
      <NavbarUsuario />
      <section className="reservas-section">
        <div className="reservas-container">
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
              filteredHoteles.map((hotel) => (
                <CursoCard
                  key={hotel.id}
                  id={hotel.id}
                  titulo={hotel.nombre}
                  nivel={hotel.provincia}
                  duracion={hotel.direccion}
                  docente={hotel.descripcion}
                  precio={`$ ${hotel.precio}`}
                />
              ))
            )}
          </div>
        </div>
      </section>
    </>
  );
}
