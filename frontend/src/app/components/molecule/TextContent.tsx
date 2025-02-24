import React from "react";
import Heading from "../atom/heroSectionHeading";
import Paragraph from "../atom/landingPageParagraph";
import Button from "../atom/ButtonText";

const TextContent = () => {
  return (
    <div className="relative z-10 text-center text-white">
      <Heading />
      <Paragraph />
      <Button buttonText={"Get Started"}/>
    </div>
  );
};

export default TextContent;

