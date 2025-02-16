import React from "react";

const MenuItemText = ({ menuItemText, isActive, itemUrl }: { menuItemText: string; isActive: boolean; itemUrl: string }) => {
  return (
    <li
      className={`p-2 text-lg transition-all duration-300 rounded-md transform 
        ${isActive ? "scale-110 text-blue-500" : "text-gray-200"}
        hover:scale-110 hover:text-blue-500 cursor-pointer`}
    >
      <a href={itemUrl} className="flex items-center px-3 py-4 font-semibold">
        {menuItemText}
      </a>
    </li>
  );
};

export default MenuItemText;
