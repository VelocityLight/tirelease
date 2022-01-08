import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Example from "./pages/example";
import DataBoard from "./pages/databoard";
import Triage from "./pages/triage";
import Assignments from "./pages/assignments";
import AboutCI from "./pages/aboutci";
import Release from "./pages/release";
import RecentOpen from "./pages/open";
import RecentClose from "./pages/close";


const MyRoutes = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Example />} />
        <Route path="/open" element={<RecentOpen />} />
        <Route path="/close" element={<RecentClose />} />
        <Route path="/home" element={<Example />} />
        <Route path="/example" element={<Example />} />
        <Route path="/release" element={<Release />} />
        <Route path="/home/databoard" element={<DataBoard />} />
        <Route path="/home/triage" element={<Triage />} />
        <Route path="/assignments" element={<Assignments />} />
        <Route path="/home/aboutci" element={<AboutCI />} />
      </Routes>
    </BrowserRouter>
  );
};

export default MyRoutes;
