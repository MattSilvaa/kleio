import React from 'react'
import {NavBar} from './components/navbar/navBar'
import {Divider} from '@mui/material'
import {SearchDropdowns} from './components/searchdropdowns/searchDropdowns'
import {Chart} from './components/chart/chart'

export const App = () => {
    return (
        <>
            <NavBar/>
            <SearchDropdowns/>
            <Divider/>
            <Chart/>
        </>
    )
}
