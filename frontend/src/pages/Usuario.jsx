import React from "react";
import { NavbarUsuario } from "../components/NavbarUsuario";
import { HeroSection } from "../components/HeroSection";
import imgPortada from "../assets/img-portada.png";

export function Usuario() {
  return (
    <>
      <NavbarUsuario />
      <HeroSection
        titulo="Hotel VEYOR"
        subtitulo="Descubre tu estadía ideal"
        background={imgPortada}
      />
    </>
  );
}
