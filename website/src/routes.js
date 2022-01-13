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
        <Route path="/" element={<RecentOpen />} />
        <Route path="/home" element={<RecentOpen />} />
        <Route path="/home/example" element={<Example />} />
        <Route path="/home/open" element={<RecentOpen />} />
        <Route path="/home/close" element={<RecentClose />} />
        <Route path="/home/release" element={<Release />} />
        <Route path="/home/databoard" element={<DataBoard />} />
        <Route path="/home/triage" element={<Triage />} />
        <Route path="/home/assignments" element={<Assignments />} />
        <Route path="/home/aboutci" element={<AboutCI />} />
      </Routes>
    </BrowserRouter>
  );
};

export default MyRoutes;
