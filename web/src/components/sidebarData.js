import React from 'react';
import { AiFillHome } from 'react-icons/ai';
import { BsPersonFill, BsPeopleFill, BsTerminal } from 'react-icons/bs';

export const SidebarData = [
  {
    title: 'Home',
    path: '/Home',
    icon: <AiFillHome />,
    cName: 'nav-text'
  },
  {
    title: 'Sign-In',
    path: '/SignIn',
    icon: <BsPersonFill />,
    cName: 'nav-text'
  },
  {
    title: 'Users Page',
    path: '/UsersPage',
    icon: <BsPeopleFill />,
    cName: 'nav-text'
  },
  {
    title: 'Etsy Auth Page',
    path: '/EtsyAuth',
    icon: <BsTerminal />,
    cName: 'nav-text'
  }
]