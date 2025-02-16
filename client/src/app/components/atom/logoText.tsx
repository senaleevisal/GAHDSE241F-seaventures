import React from 'react'

const logoText = ({ logoname }: { logoname: string }) => {


  return (

    <a href="#" className="text-xl font-bold text-white hover:text-blue-400 transition duration-300">
    {logoname}
  </a>
  )
}

export default logoText