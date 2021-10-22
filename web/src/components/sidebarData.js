import React from 'react';
import { AiFillHome, AiFillCreditCard } from 'react-icons/ai';
import { BsPersonFill, BsPeopleFill } from 'react-icons/bs';

export const SidebarData = [
    {
        title: ' Home',
        path: '/Home',
        icon: <AiFillHome />,
        cName: 'nav-text'
    },
    {
        title: ' Sign-In',
        path: '/SignIn',
        icon: <BsPersonFill />,
        cName: 'nav-text'
    },
    {
        title: ' Users Page',
        path: '/UsersPage',
        icon: <BsPeopleFill />,
        cName: 'nav-text'
    },
    {
        title: ' Orders',
        path: '/Orders',
        icon: <AiFillCreditCard />,
        cName: 'nav-text'
    }
]