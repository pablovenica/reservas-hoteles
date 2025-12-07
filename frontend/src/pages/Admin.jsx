import React, { useEffect, useState } from "react";
import "./Admin.css";

export function Button({ children, onClick, className }) {
  return (
    <button className={className} onClick={onClick}>
      {children}
    </button>
  );
}

export function HotelForm({ onSubmit, onClose, initialData }) {
  const [form, setForm] = useState(
    initialData || {
      nombre: "",
      provincia: "",
      direccion: "",
      precio: 0,
    }
  );

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(form);
  };

  return (
    <div className="modal-overlay">
      <div className="modal">
        <h2>{initialData ? "Editar Hotel" : "Agregar Hotel"}</h2>

        <form onSubmit={handleSubmit}>
          <input
            name="nombre"
            placeholder="Nombre"
            value={form.nombre}
            onChange={handleChange}
            required
          />

          <input
            name="provincia"
            placeholder="Provincia"
            value={form.provincia}
            onChange={handleChange}
            required
          />

          <input
            name="direccion"
            placeholder="Dirección"
            value={form.direccion}
            onChange={handleChange}
            required
          />

          <input
            type="number"
            name="precio"
            placeholder="Precio"
            value={form.precio}
            onChange={handleChange}
            required
          />

          <div className="modal-buttons">
            <button type="submit" className="btn-agregar">
              Guardar
            </button>
            <button type="button" className="btn-eliminar" onClick={onClose}>
              Cancelar
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}

export function Admin() {
  const [hotels, setHotels] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showForm, setShowForm] = useState(false);
  const [editHotel, setEditHotel] = useState(null);

  const fetchHotels = async () => {
    try {
      const res = await fetch("http://localhost:8082/hotels");
      const data = await res.json();

      const list = data.hoteles || data.hotels || [];
      setHotels(Array.isArray(list) ? list : []);
    } catch (error) {
      console.error("Error cargando hoteles", error);
      setHotels([]);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchHotels();
  }, []);

  const addHotel = async (hotel) => {
    const token = localStorage.getItem("token");

    await fetch("http://localhost:8082/hotels", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        ...hotel,
        precio: Number(hotel.precio),
        imagen: "",
        descripcion: "",
      }),
    });

    setShowForm(false);
    fetchHotels();
  };

  const updateHotel = async (hotel) => {
    const token = localStorage.getItem("token");

    await fetch(`http://localhost:8082/hotels/${editHotel.id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        ...hotel,
        precio: Number(hotel.precio),
        imagen: editHotel?.imagen || "",
        descripcion: editHotel?.descripcion || "",
      }),
    });

    setEditHotel(null);
    setShowForm(false);
    fetchHotels();
  };

  const deleteHotel = async (id) => {
    const token = localStorage.getItem("token");

    await fetch(`http://localhost:8082/hotels/${id}`, {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    fetchHotels();
  };

  if (loading) return <p>Cargando hoteles...</p>;

  return (
    <div className="admin-container">
      <main className="admin-content">
        <h1>Gestión de Hoteles</h1>

        <div className="admin-panel">
          <button className="btn-agregar" onClick={() => setShowForm(true)}>
            + Agregar nuevo hotel
          </button>

          <table className="tabla-cursos">
            <thead>
              <tr>
                <th>ID</th>
                <th>Nombre</th>
                <th>Provincia</th>
                <th>Dirección</th>
                <th>Precio</th>
                <th>Acciones</th>
              </tr>
            </thead>
            <tbody>
              {hotels.map((h) => (
                <tr key={h.id}>
                  <td>{h.id}</td>
                  <td>{h.nombre}</td>
                  <td>{h.provincia}</td>
                  <td>{h.direccion}</td>
                  <td>${h.precio}</td>
                  <td>
                    <button
                      className="btn-editar"
                      onClick={() => {
                        setEditHotel(h);
                        setShowForm(true);
                      }}
                    >
                      Editar
                    </button>
                    <button
                      className="btn-eliminar"
                      onClick={() => deleteHotel(h.id)}
                    >
                      Eliminar
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>

        {showForm && (
          <HotelForm
            initialData={editHotel}
            onSubmit={editHotel ? updateHotel : addHotel}
            onClose={() => {
              setEditHotel(null);
              setShowForm(false);
            }}
          />
        )}
      </main>
    </div>
  );
}
