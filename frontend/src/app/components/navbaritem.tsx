'use client';

import React, { useState } from 'react'



const NavbarItem = () => {
    const [menuOpen, setMenuOpen] = useState(false);
  
    const navItems = [
      { name: "Home", url: "#" },
      { name: "Search", url: "#" },
      { name: "About Us", url: "#" },
    ];
  
    return (
      <nav className="block w-full max-w-screen-lg px-4 py-2 mx-auto bg-white shadow-md rounded-md lg:px-8 lg:py-3 mt-10">
        <div className="container flex items-center justify-between mx-auto text-slate-800">
          <a href="#" className="text-base font-semibold text-slate-800">
            SeaVentures
          </a>
  
          <button
            className="lg:hidden"
            onClick={() => setMenuOpen(!menuOpen)}
          >
            <svg xmlns="http://www.w3.org/2000/svg" className="w-6 h-6" fill="none" stroke="currentColor" strokeWidth="2">
              <path strokeLinecap="round" strokeLinejoin="round" d="M4 6h16M4 12h16M4 18h16"></path>
            </svg>
          </button>
  
          {/* Navigation Menu */}
          <ul className={`lg:flex lg:items-center lg:gap-6 ${menuOpen ? "block" : "hidden"} w-full lg:w-auto mt-2 lg:mt-0`}>
            {navItems.map((item, index) => (
              <li key={index} className="p-1 text-sm text-slate-600">
                <a href={item.url} className="flex items-center">
                  {item.name}
                </a>
              </li>
            ))}
          </ul>
        </div>
      </nav>
    );
  };

export default NavbarItem