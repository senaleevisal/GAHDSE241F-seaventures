import React from "react";
import FooterLinks from "../molecule/footerLinks";
import AppLinks from "../molecule/AppLinks";
import FooterText from "../atom/FooterText";

const Footer = () => {
  return (
    <footer className="bg-gray-900 text-white py-6">
      <div className="max-w-6xl mx-auto flex flex-col md:flex-row flex-wrap justify-between items-center px-6 gap-y-6">
        
        <div className="text-center md:text-left max-w-sm">
          <h1 className="text-2xl font-bold">Sea Ventures</h1>
          <FooterText footertext="Discover the perfect beach getaway. 🌊 Your adventure begins with us." />
        </div>
        <FooterLinks />
        <AppLinks />
      </div>

      <hr className="border-gray-700 my-4 mx-6" />
      <div className="text-center">
        <FooterText footertext="© 2025 Sea Ventures. All rights reserved." />
      </div>
    </footer>
  );
};

export default Footer;
