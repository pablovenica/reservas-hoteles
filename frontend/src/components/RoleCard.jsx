import React from "react";
import { Link } from "react-router-dom";
import "./RoleCard.css";

export function RoleCard({ img, titulo, link, gradiente }) {
  return (
    <Link to={link} className={`role-card ${gradiente}`}>
      <img src={img} alt={titulo} className="role-card-img" />
      <h2>{titulo}</h2>
    </Link>
  );
}
