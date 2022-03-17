import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Example from "./pages/example";
import DataBoard from "./pages/databoard";
// import Triage from "./pages/triage";
import Assignments from "./pages/assignments";
import TiFC from "./pages/tifc";
import Release from "./pages/release";
import RecentOpen from "./pages/open";
import RecentClose from "./pages/close";
import VersionPage from "./pages/version";
import AffectTriage from "./pages/affects";
import PickTriage from "./pages/pick";

const MyRoutes = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<RecentOpen />} />
        <Route path="/home" element={<RecentOpen />} />
        <Route path="/home/example" element={<Example />} />
        <Route path="/home/open" element={<RecentOpen />} />
        <Route path="/home/affection" element={<AffectTriage />} />
        <Route path="/home/cherrypick" element={<PickTriage />} />
        <Route path="/home/close" element={<RecentClose />} />
        <Route path="/home/databoard" element={<DataBoard />} />
        {/* <Route path="/home/triage" element={<Triage />} /> */}
        <Route path="/home/assignments" element={<Assignments />} />
        <Route path="/home/tifc" element={<TiFC />} />
        <Route path="/home/version" element={<VersionPage />} />
        <Route path="/home/triage" element={<Release />} />
      </Routes>
    </BrowserRouter>
  );
};

export default MyRoutes;
