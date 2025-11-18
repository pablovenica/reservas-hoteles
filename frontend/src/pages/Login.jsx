import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import Swal from "sweetalert2";
import "./Login.css";

export function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const res = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email,
          password,
        }),
      });

      const data = await res.json();

      if (!res.ok) {
        Swal.fire({
          icon: "error",
          title: "Credenciales incorrectas",
          text: data.error || "Revisá tu correo y contraseña.",
          confirmButtonColor: "#00bcd4",
        });
        return;
      }

      localStorage.setItem("token", data.token);
      localStorage.setItem("rol", data.rol);
      localStorage.setItem("userID", data.userID);

      const nombreParaMostrar = email.split("@")[0];
      localStorage.setItem("userName", nombreParaMostrar);

      await Swal.fire({
        icon: "success",
        title: "Login exitoso",
        text: "Bienvenido a VEYOR Hotels.",
        timer: 1600,
        showConfirmButton: false,
      });

      if (data.rol === "admin") {
        navigate("/admin");
      } else {
        navigate("/usuario");
      }
    } catch (error) {
      console.error("Error:", error);
      Swal.fire({
        icon: "error",
        title: "Error en el servidor",
        text: "Ocurrió un problema al intentar iniciar sesión. Probá más tarde.",
        confirmButtonColor: "#00bcd4",
      });
    }
  };

  return (
    <section className="login-section">
      <div className="login-container">
        <h2 className="login-title">Iniciar Sesión</h2>
        <form className="login-form" onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="email">Correo electrónico</label>
            <input
              type="email"
              id="email"
              placeholder="tucorreo@example.com"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>

          <div className="form-group">
            <label htmlFor="password">Contraseña</label>
            <input
              type="password"
              id="password"
              placeholder="********"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>

          <button type="submit" className="btn-login">
            Ingresar
          </button>
        </form>

        <p className="register-text">
          ¿No te registraste todavía?{" "}
          <Link to="/registro" className="register-link">
            Registrate
          </Link>
        </p>
      </div>
    </section>
  );
}
