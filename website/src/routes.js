import React from "react";
import {BrowserRouter, Routes, Route} from "react-router-dom"
import Example from "./pages/example";
import DataBoard from "./pages/databoard";
import Triage from "./pages/triage";
import AboutCI from "./pages/aboutci";

const MyRoutes = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path = "/" element = {<Example/>} />
                <Route path = "/home" element = {<Example/>} />
                <Route path = "/home/example" element = {<Example/>} />
                <Route path = "/home/databoard" element = {<DataBoard/>} />
                <Route path = "/home/triage" element = {<Triage/>} />
                <Route path = "/home/aboutci" element = {<AboutCI/>} />
            </Routes>
        </BrowserRouter>
    )
};

export default MyRoutes;