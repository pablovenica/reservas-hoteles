import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import "./MisReservas.css";

const RESERVATIONS_API_BASE = "http://localhost:8083";

function MisReservas() {
  const [reservas, setReservas] = useState([]);
  const [loading, setLoading] = useState(true);
  const [errorMsg, setErrorMsg] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    const fetchReservas = async () => {
      try {
        const userIdStr = localStorage.getItem("userID");

        if (!userIdStr) {
          setErrorMsg("Tenés que iniciar sesión para ver tus reservas.");
          setLoading(false);
          return;
        }

        const url = `${RESERVATIONS_API_BASE}/reservations/user/${userIdStr}`;
        console.log("➡️ Fetch Mis Reservas:", url);

        const res = await fetch(url);
        const rawText = await res.text();

        console.log("Status /reservations/user:", res.status);
        console.log("Body /reservations/user:", rawText);

        if (!res.ok) {
          setErrorMsg(`Error al cargar tus reservas. Código: ${res.status}`);
          setReservas([]);
          setLoading(false);
          return;
        }

        let data;
        try {
          data = rawText ? JSON.parse(rawText) : [];
        } catch (e) {
          console.error("Error parseando JSON de reservas:", e);
          setErrorMsg("Respuesta inválida del servidor.");
          setReservas([]);
          setLoading(false);
          return;
        }

        console.log("Reservas del usuario:", data);
        setReservas(Array.isArray(data) ? data : []);
        setErrorMsg("");
      } catch (error) {
        console.error("Error en MisReservas:", error);
        setErrorMsg("No se pudieron cargar tus reservas.");
      } finally {
        setLoading(false);
      }
    };

    fetchReservas();
  }, []);

  const formatDate = (isoString) => {
    if (!isoString) return "—";
    const d = new Date(isoString);
    if (Number.isNaN(d.getTime())) return isoString;
    return d.toLocaleDateString("es-AR");
  };

  const handleVolver = () => {
    navigate(-1);
  };

  if (loading) {
    return (
      <main className="mis-reservas">
        <p>Cargando tus reservas...</p>
      </main>
    );
  }

  if (errorMsg) {
    return (
      <main className="mis-reservas">
        <h1>Mis reservas</h1>
        <p>{errorMsg}</p>
        <button
          className="detalle-hotel__btn"
          style={{ marginTop: "1.5rem" }}
          onClick={handleVolver}
        >
          Volver
        </button>
      </main>
    );
  }

  if (!reservas || reservas.length === 0) {
    return (
      <main className="mis-reservas">
        <h1>Mis reservas</h1>
        <p>No tenés reservas registradas todavía.</p>
        <button
          className="detalle-hotel__btn"
          style={{ marginTop: "1.5rem" }}
          onClick={handleVolver}
        >
          Volver
        </button>
      </main>
    );
  }

  return (
    <main className="mis-reservas">
      <h1>Mis reservas</h1>

      <div className="mis-reservas__lista">
        {reservas.map((reserva) => (
          <article key={reserva.id} className="mis-reservas__item">
            <h2>Hotel ID: {reserva.id_hoteles}</h2>

            <p>
              Check-in: <strong>{formatDate(reserva.fecha_ingreso)}</strong>
            </p>
            <p>
              Check-out: <strong>{formatDate(reserva.fecha_salida)}</strong>
            </p>
            <p>
              Estado: <strong>{reserva.estado}</strong>
            </p>
          </article>
        ))}
      </div>

      <button
        className="detalle-hotel__btn"
        style={{ marginTop: "2rem" }}
        onClick={handleVolver}
      >
        Volver
      </button>
    </main>
  );
}

export default MisReservas;
