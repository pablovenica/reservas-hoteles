import React from "react";
import "./Admin.css";
import { AdminSidebar } from "../components/AdminSidebar";

export function Admin() {
  return (
    <div className="admin-container">
      <AdminSidebar />

      <main className="admin-content">
        <h1>Gestión de Cursos</h1>
        <p>Desde aquí podrás agregar, editar o eliminar cursos.</p>

        <div className="admin-panel">
          <button className="btn-agregar">+ Agregar nuevo curso</button>

          <table className="tabla-cursos">
            <thead>
              <tr>
                <th>ID</th>
                <th>Título</th>
                <th>Nivel</th>
                <th>Duración</th>
                <th>Docente</th>
                <th>Acciones</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>1</td>
                <td>Curso de Desarrollo Web</td>
                <td>Inicial</td>
                <td>10 semanas</td>
                <td>Lucía Fernández</td>
                <td>
                  <button className="btn-editar">Editar</button>
                  <button className="btn-eliminar">Eliminar</button>
                </td>
              </tr>
              <tr>
                <td>2</td>
                <td>Curso de JavaScript</td>
                <td>Intermedio</td>
                <td>10 semanas</td>
                <td>Martín Pérez</td>
                <td>
                  <button className="btn-editar">Editar</button>
                  <button className="btn-eliminar">Eliminar</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </main>
    </div>
  );
}
