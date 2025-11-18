import React from "react";
import "./HeroSection.css";

export function HeroSection({ background }) {
  return (
    <section
      className="hero-section"
      style={{
        backgroundImage: `url(${background})`,
        backgroundSize: "cover",
        backgroundPosition: "center",
      }}
    >
      <div className="hero-overlay"></div>
      <div className="hero-content">
        <div className="hero-text">
          <h1>Descubre tu estadía ideal</h1>
          <p>
            Reserva fácilmente en los mejores hoteles y vive experiencias únicas
            de descanso, confort y elegancia.
          </p>
        </div>
      </div>
    </section>
  );
}
