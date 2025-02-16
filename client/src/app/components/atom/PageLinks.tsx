import React from "react";

const PageLinks = ({ itemText, isActive, itemUrl }: { itemText: string; isActive: boolean; itemUrl: string }) => {
  return (
    <li className={`p-1 text-sm ${isActive ? "text-slate-800" : "text-slate-400"}`}>
      <a href={itemUrl} className="flex items-center">
        {itemText}
      </a>
    </li>
  );
};

export default PageLinks;
