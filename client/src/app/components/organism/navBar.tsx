import React from 'react'
import NavContainer from '../atom/navContainer'
import LogoText from '../atom/logoText';
import NavMenuItems from '../molecule/navMenuItems';


const navItems = [
    { name: 'Home', url: '/' },
    { name: 'About', url: '/about'},
    { name: 'Contact', url: '/contact' },

];

const navBar = () => {

  
  return (
    <NavContainer>

        <LogoText logoname="SeaVentures" />
        <NavMenuItems menuOpen={false} navItems={navItems} />

    </NavContainer>
  )
}

export default navBar