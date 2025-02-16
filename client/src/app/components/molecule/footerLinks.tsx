import React from "react";
import PageLinks from "../atom/PageLinks";

const links = [
  { name: "Home", href: "/" },
  { name: "Search", href: "/search" },
  { name: "About Us", href: "/about" },
];

const footerLinks = () => {
  return (
    <div className="text-center md:text-left">
      <h3 className="text-white font-bold">Quick Links</h3>
      <ul className="mt-2 space-y-2">
        {links.map((link, index) => (
          <PageLinks key={index} itemText={link.name} itemUrl={link.href} isActive={false} />
        ))}
      </ul>
    </div>
  );
};

export default footerLinks;
