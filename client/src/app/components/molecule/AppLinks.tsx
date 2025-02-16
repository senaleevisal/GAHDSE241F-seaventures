/* eslint-disable @next/next/no-img-element */
import React from "react";

const AppLinks = () => {
  return (
    <div className="text-center md:text-left">
      <h3 className="text-white font-bold">Download Our App</h3>
      <ul className="mt-2 space-y-2">
        <li>
          <a href="#" className="text-gray-400 hover:text-white flex items-center justify-center md:justify-start">
            Google Play
            <img src="footer_img/playstore.png" alt="Google Play" className="w-10 h-10 ml-2" />
          </a>
        </li>
        <li>
          <a href="#" className="text-gray-400 hover:text-white flex items-center justify-center md:justify-start">
            App Store
            <img src="footer_img/appstore.png" alt="App Store" className="w-10 h-10 ml-2" />
          </a>
        </li>
      </ul>
    </div>
  );
};

export default AppLinks;
