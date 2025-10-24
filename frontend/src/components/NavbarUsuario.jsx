import React from "react";
import { Link } from "react-router-dom";
import "./NavbarUsuario.css";
import logo from "../assets/logo-usuario.png";

export function NavbarUsuario() {
  return (
    <nav className="navbar-usuario">
      <div className="navbar-logo">
        <Link to="/usuario">
          <img src={logo} alt="Logo Hotel VEYOR" />
        </Link>
      </div>

      <ul className="nav-links">
        <li>
          <Link to="/usuario">Inicio</Link>
        </li>
        <li>
          <Link to="/usuario/reservas">Reservas</Link>
        </li>
      </ul>

      <div className="navbar-btn d-flex align-items-center">
        <Link to="/login" className="btn-comenzar me-3">
          Login
        </Link>
        <button type="button" className="btn-comenzar ms-3">
          Logout
        </button>
      </div>
    </nav>
  );
}
