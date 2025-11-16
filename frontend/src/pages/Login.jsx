import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
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
        alert(data.error || "Credenciales incorrectas");
        return;
      }

      // 🔹 Guardamos token y datos mínimos del usuario
      localStorage.setItem("token", data.token);
      localStorage.setItem("rol", data.rol);
      localStorage.setItem("userID", data.userID);

      // 🔹 Nombre a mostrar en el navbar (antes del @ del email)
      const nombreParaMostrar = email.split("@")[0];
      localStorage.setItem("userName", nombreParaMostrar);

      alert("Login exitoso");

      // 🔹 Redirigimos según rol (si querés podés dejar todo a /usuario)
      if (data.rol === "admin") {
        navigate("/admin");
      } else {
        navigate("/usuario");
      }
    } catch (error) {
      console.error("Error:", error);
      alert("Error en el servidor");
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
