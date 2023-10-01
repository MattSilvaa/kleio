import React from 'react'
import {LineChart} from '@mui/x-charts'
import {Container} from '@mui/material'


export const Chart = () => {
    return <Container sx={{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        gap: '10%',
        direction: 'row',
        padding: '2%',
    }}>
        <LineChart xAxis={[{data: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]}]}
                   series={[
                       {
                           data: [2, 3, 5.5, 8.5, 1.5, 5, 1, 4, 3, 8],
                           showMark: ({index}) => index % 2 === 0,
                       },
                   ]}
                   width={1200}
                   height={1000}
        />
    </Container>
}