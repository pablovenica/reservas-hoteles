import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import Swal from "sweetalert2";
import "./Registro.css";

export function Registro() {
  const [formData, setFormData] = useState({
    nombre: "",
    email: "",
    password: "",
    tipo_usuario: "normal",
  });

  const navigate = useNavigate();

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
        Swal.fire({
          icon: "error",
          title: "Error al registrarse",
          text: data.error || "No se pudo crear el usuario. Intentá nuevamente.",
          confirmButtonColor: "#00bcd4",
        });
        return;
      }

      await Swal.fire({
        icon: "success",
        title: "Usuario registrado",
        text: "Tu cuenta fue creada correctamente. Ahora podés iniciar sesión.",
        confirmButtonColor: "#00bcd4",
      });

      navigate("/login");
    } catch (error) {
      console.error(error);
      Swal.fire({
        icon: "error",
        title: "Error en el servidor",
        text: "Ocurrió un problema al registrar tu usuario. Intentá más tarde.",
        confirmButtonColor: "#00bcd4",
      });
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
