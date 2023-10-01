/// <reference lib="dom" />
/// <reference lib="dom.iterable" />
import {createRoot} from 'react-dom/client'
import {App} from './app'
import React from 'react'

// Importing html file
// require("./index.html")

const container = document.getElementById('root')
if (container) {
    const root = createRoot(container)
    root.render(<App/>)
}