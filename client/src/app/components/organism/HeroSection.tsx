import React from "react";
import TextContent from "../molecule/TextContent";

const HeroSection = () => {
  return (
    <div
      className="relative flex items-center justify-center w-full h-screen bg-cover bg-center"
      style={{ backgroundImage: "url('/landingPage_img/bg_img.png')" }}>
      <div className="absolute inset-0 bg-black bg-opacity-50"></div>
      <TextContent />
    </div>
  );
};

export default HeroSection;
