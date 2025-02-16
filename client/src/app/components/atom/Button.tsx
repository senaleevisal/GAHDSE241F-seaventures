import React from "react";

const Button = ({ buttonText }: { buttonText: String }) => {
  return (
    <button className="mt-6 px-6 py-3 text-lg font-semibold text-white bg-blue-500 rounded-lg shadow-md hover:bg-blue-600 transition">
      {buttonText}
    </button>
  );
};

export default Button;
