import React from 'react';

const navContainer = ({ children }: { children: React.ReactNode }) => {
  return (
    <nav className="w-full px-6 py-3 bg-gray-900 shadow-lg">
      <div className="container flex items-center justify-between mx-auto">
        {children}
      </div>
    </nav>
  );
};

export default navContainer;
