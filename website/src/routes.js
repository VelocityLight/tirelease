import React from "react";
import {BrowserRouter, Routes, Route} from "react-router-dom"
import Example from "./pages/example";
import DataBoard from "./pages/databoard";

const MyRoutes = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path = "/" element = {<Example/>} />
                <Route path = "/databoard" element = {<DataBoard/>} />
            </Routes>
        </BrowserRouter>
    )
};

export default MyRoutes;