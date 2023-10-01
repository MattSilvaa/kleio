import {Autocomplete as MuiAutocomplete, TextField} from '@mui/material'
import {SxProps} from '@mui/system'
import React from 'react'

type Props = {
    id?: string
    options: string[]
    labelTitle?: string
    sx?: SxProps
    fullWidth?: boolean
}

export const Autocomplete = (props: Props) => {
    const {id, options, labelTitle = '', sx, fullWidth} = props

    return <MuiAutocomplete disablePortal id={id} options={options} sx={sx} fullWidth
                            renderInput={(params) => <TextField {...params} label={labelTitle}/>}/>
}