import React from 'react'
import {Autocomplete} from '../autocomplete/autocomplete'
import {jobLocationSearchOptions, jobTitleSearchOptions} from './types'
import {Container} from '@mui/material'

// type Props = {}

const autocompleteStyle = {
    width: '50%'
}
export const SearchDropdowns = () => {

    return <Container
        sx={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            gap: '10%',
            direction: 'row',
            padding: '2%',
            width: '100%'
        }}>
        <Autocomplete id={'job-title'} options={jobTitleSearchOptions} labelTitle={'Job Title'} sx={autocompleteStyle}
                      fullWidth/>
        <Autocomplete id={'location'} options={jobLocationSearchOptions} labelTitle={'Location (US Only)'}
                      sx={autocompleteStyle} fullWidth/>
    </Container>
}


