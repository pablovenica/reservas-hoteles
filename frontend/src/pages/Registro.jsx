import React, { useState } from "react";
import { Link } from "react-router-dom";
import "./Registro.css";

export function Registro() {
  const [formData, setFormData] = useState({
    nombre: "",
    email: "",
    password: "",
    tipo_usuario: "normal",
  });

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const res = await fetch("http://localhost:8080/users", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });

      const data = await res.json();

      if (!res.ok) {
        alert(data.error || "Error registrando usuario");
        return;
      }

      alert("Usuario registrado con éxito");
      window.location.href = "/login";
    } catch (error) {
      console.error(error);
      alert("Error en el servidor");
    }
  };

  return (
    <section className="registro-section">
      <div className="registro-container">
        <h2 className="registro-title">Crear cuenta</h2>
        <form className="registro-form" onSubmit={handleSubmit}>
          
          <div className="form-group">
            <label htmlFor="nombre">Nombre</label>
            <input
              type="text"
              id="nombre"
              name="nombre"
              placeholder="Tu nombre"
              value={formData.nombre}
              onChange={handleChange}
              required
            />
          </div>

          <div className="form-group">
            <label htmlFor="email">Correo electrónico</label>
            <input
              type="email"
              id="email"
              name="email"
              placeholder="tucorreo@example.com"
              value={formData.email}
              onChange={handleChange}
              required
            />
          </div>

          <div className="form-group">
            <label htmlFor="password">Contraseña</label>
            <input
              type="password"
              id="password"
              name="password"
              placeholder="********"
              value={formData.password}
              onChange={handleChange}
              required
            />
          </div>

          <button type="submit" className="btn-registro">
            Registrarse
          </button>
        </form>

        <p className="login-text">
          ¿Ya tenés cuenta?{" "}
          <Link to="/login" className="login-link">
            Iniciá sesión
          </Link>
        </p>
      </div>
    </section>
  );
}
