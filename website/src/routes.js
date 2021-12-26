import React from "react";
import {BrowserRouter, Routes, Route} from "react-router-dom"
import Example from "./pages/example";
import DataBoard from "./pages/databoard";
import Triage from "./pages/triage";

const MyRoutes = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path = "/page/example" element = {<Example/>} />
                <Route path = "/page/databoard" element = {<DataBoard/>} />
                <Route path = "/page/triage" element = {<Triage/>} />
            </Routes>
        </BrowserRouter>
    )
};

export default MyRoutes;