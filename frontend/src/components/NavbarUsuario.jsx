import React, { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import "./NavbarUsuario.css";
import logo from "../assets/logo-usuario.png";

export function NavbarUsuario() {
  const [userName, setUserName] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    const name = localStorage.getItem("userName");

    if (token && name) {
      setUserName(name);
    } else {
      setUserName(null);
    }
  }, []);

  const handleLogout = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("rol");
    localStorage.removeItem("userID");
    localStorage.removeItem("userName");

    setUserName(null);
    navigate("/");
  };

  const isLoggedIn = !!userName;

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
        <li>
          <Link to="/usuario/misreservas">Mis Reservas</Link>
        </li>
      </ul>

      <div className="navbar-btn d-flex align-items-center">
        {!isLoggedIn && (
          <Link to="/login" className="btn-comenzar">
            Login
          </Link>
        )}

        {isLoggedIn && (
          <>
            <span
              className="navbar-user-name me-3"
              style={{ marginRight: "0.75rem", fontWeight: 600 }}
            >
              {userName}
            </span>
            <button
              type="button"
              className="btn-comenzar ms-3"
              onClick={handleLogout}
            >
              Logout
            </button>
          </>
        )}
      </div>
    </nav>
  );
}
