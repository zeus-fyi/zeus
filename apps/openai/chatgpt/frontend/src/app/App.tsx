import './App.css';

import {BrowserRouter, Route, Routes} from "react-router-dom";
import {ChatGPTPage} from "../chatgpt/ChatGPTWrapper";


export const App = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<ChatGPTPage/>}/>
            </Routes>
        </BrowserRouter>
    );
}

