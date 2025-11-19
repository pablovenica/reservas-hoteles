import React from "react";
import "./AdminSidebar.css";

export function AdminSidebar() {
  return (
    <aside className="sidebar">
      <div className="sidebar-header">
        <h2>Panel Admin</h2>
      </div>
      <hr />
      <ul className="sidebar-nav">
        <li><a href="#">🏠 Inicio</a></li>
        <li><a href="#">📊 Dashboard</a></li>
        <li><a href="#">📚 Cursos</a></li>
        <li><a href="#">👥 Estudiantes</a></li>
        <li><a href="#">⚙️ Configuración</a></li>
      </ul>
      <hr />
    </aside>
  );
}
