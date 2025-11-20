import React, { useEffect, useState } from "react";
import Swal from "sweetalert2";
import { NavbarUsuario } from "../components/NavbarUsuario";
import { CursoCard } from "../components/CursoCard";
import "./Reservas.css";

const SEARCH_API_BASE = "http://localhost:8084";

export function Reservas() {
  const [hoteles, setHoteles] = useState([]);
  const [filteredHoteles, setFilteredHoteles] = useState([]);

  const [ubicacion, setUbicacion] = useState("");
  const [fechaEntrada, setFechaEntrada] = useState("");
  const [fechaSalida, setFechaSalida] = useState("");

  const [page, setPage] = useState(1);
  const [pageSize] = useState(10);
  const [totalResultados, setTotalResultados] = useState(0);

  useEffect(() => {
    const fetchHotelesInicial = async () => {
      try {
        const url = `${SEARCH_API_BASE}/search/hotels?page=${page}&page_size=${pageSize}`;
        console.log("➡️ Fetch inicial search-api:", url);

        const res = await fetch(url);
        const data = await res.json();
        console.log("➡️ Respuesta cruda de /search/hotels:", data);

        const lista = Array.isArray(data.results) ? data.results : [];

        setHoteles(lista);
        setFilteredHoteles(lista);
        setTotalResultados(data.total || lista.length);
      } catch (error) {
        console.error("Error al pedir /search/hotels (inicial):", error);
        Swal.fire({
          icon: "error",
          title: "Error al cargar habitaciones",
          text: "No pudimos obtener la lista de habitaciones. Intentalo nuevamente.",
          confirmButtonColor: "#00bcd4",
        });
      }
    };

    fetchHotelesInicial();
  }, [page, pageSize]);

  const handleBuscar = async (e) => {
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

    try {
      const q = encodeURIComponent(ubicacion.trim());
      const url = `${SEARCH_API_BASE}/search/hotels?q=${q}&page=1&page_size=${pageSize}`;

      console.log("➡️ Fetch search-api con filtro:", url);

      const res = await fetch(url);
      const data = await res.json();
      console.log("➡️ Respuesta filtrada de /search/hotels:", data);

      const lista = Array.isArray(data.results) ? data.results : [];

      setHoteles(lista);
      setFilteredHoteles(lista);
      setTotalResultados(data.total || lista.length);
      setPage(1);

      if (lista.length === 0) {
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
    } catch (error) {
      console.error("Error al buscar en search-api:", error);
      Swal.fire({
        icon: "error",
        title: "Error al buscar habitaciones",
        text: "No pudimos realizar la búsqueda. Intentalo nuevamente.",
        confirmButtonColor: "#00bcd4",
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

          <h2 className="titulo-principal">
            Nuestras Habitaciones{" "}
            {totalResultados > 0 && (
              <span style={{ fontSize: "0.9rem", color: "#c6c6c6" }}>
                ({totalResultados} resultados)
              </span>
            )}
          </h2>

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
