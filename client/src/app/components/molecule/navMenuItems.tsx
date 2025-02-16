import React from 'react'
import MenuItemText from '../atom/menuItemText'

const navMenuItems = ({ menuOpen, navItems }: { menuOpen: boolean, navItems: { name: string, url: string }[] }) => {
  return (
    <ul className={`absolute lg:static bg-gray-900 lg:bg-transparent lg:flex lg:items-center lg:gap-6 
      ${menuOpen ? "top-16 left-0 w-full block" : "hidden"} 
      w-full lg:w-auto mt-2 lg:mt-0 p-4 lg:p-0 transition-all duration-300`}>
    {navItems.map((item, index) => (
      <MenuItemText key={index} menuItemText={item.name} isActive={false} itemUrl={item.url} />
    ))}
  </ul>
  )
}

export default navMenuItems