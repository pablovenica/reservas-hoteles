import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import Swal from "sweetalert2";
import "./DetalleHotel.css";

function DetalleHotel() {
  const { id } = useParams();
  const navigate = useNavigate();

  const [hotel, setHotel] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchHotel = async () => {
      try {
        const res = await fetch(`http://localhost:8082/hotels/${id}`);
        if (!res.ok) throw new Error("Error al cargar el hotel");
        const data = await res.json();

        console.log("Respuesta cruda de /hotels/:id:", data);

        setHotel(data);
      } catch (error) {
        console.error(error);
        Swal.fire("Error", "No se pudo cargar la información del hotel", "error");
      } finally {
        setLoading(false);
      }
    };

    fetchHotel();
  }, [id]);

  const handleReservar = async () => {
    try {
      const userIdStr = localStorage.getItem("userID");

      if (!userIdStr) {
        Swal.fire("Debes iniciar sesión", "Inicia sesión para reservar un hotel.", "warning");
        return;
      }

      const userId = parseInt(userIdStr, 10);

      const ahora = new Date();
      const mañana = new Date();
      mañana.setDate(ahora.getDate() + 1);

      const reservaPayload = {
        id_usuarios: userId,
        id_hoteles: hotel.id,
        fecha_ingreso: ahora.toISOString(),
        fecha_salida: mañana.toISOString(),
        estado: "ACTIVA",
      };

      console.log("Enviando reserva:", reservaPayload);

      const res = await fetch("http://localhost:8083/reservations", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(reservaPayload),
      });

      const rawText = await res.text();
      console.log("Respuesta al crear reserva. Status:", res.status, "Body:", rawText);

      if (!res.ok) {
        let msg = "No se pudo crear la reserva";
        try {
          const errData = rawText ? JSON.parse(rawText) : null;
          if (errData && errData.error) msg = errData.error;
        } catch (_) {}
        throw new Error(msg);
      }

      Swal.fire("Reserva creada", "Tu reserva se registró correctamente.", "success");
    } catch (error) {
      console.error(error);
      Swal.fire("Error", error.message || "No se pudo crear la reserva", "error");
    }
  };

  const handleVolver = () => {
    navigate(-1);
  };

  if (loading) {
    return (
      <main className="detalle-hotel">
        <p>Cargando hotel...</p>
      </main>
    );
  }

  if (!hotel) {
    return (
      <main className="detalle-hotel">
        <p>No se encontró información del hotel.</p>
        <button className="detalle-hotel__btn" onClick={handleVolver}>
          Volver
        </button>
      </main>
    );
  }

  const imagen =
    hotel.imagen && hotel.imagen.trim() !== ""
      ? hotel.imagen
      : "/hotel-default.jpg";

  return (
    <main className="detalle-hotel">
      <div className="detalle-hotel__container">
        <div className="detalle-hotel__imagen">
          <img src={imagen} alt={hotel.nombre} />
        </div>

        <div className="detalle-hotel__info">
          <h1>{hotel.nombre}</h1>

          <p className="detalle-hotel__ubicacion">{hotel.provincia}</p>

          {hotel.precio && (
            <p className="detalle-hotel__precio">
              Precio por noche: <span>${hotel.precio}</span>
            </p>
          )}

          <p className="detalle-hotel__descripcion">
            <strong>Descripción:</strong> {hotel.descripcion}
          </p>

          <ul className="detalle-hotel__extras">
            {hotel.capacidad && (
              <li>
                <span>Capacidad:</span> {hotel.capacidad} personas
              </li>
            )}
            {hotel.servicios && (
              <li>
                <span>Servicios:</span> {hotel.servicios}
              </li>
            )}
          </ul>

          <div style={{ display: "flex", gap: "0.75rem", marginTop: "1.6rem" }}>
            <button className="detalle-hotel__btn" onClick={handleReservar}>
              Reservar
            </button>

            <button
              className="detalle-hotel__btn"
              style={{ background: "#374151", boxShadow: "none" }}
              onClick={handleVolver}
            >
              Volver
            </button>
          </div>
        </div>
      </div>
    </main>
  );
}

export default DetalleHotel;
