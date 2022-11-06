
import './App.css';
import NavBar from './Components/Nabvar';
import 'bootstrap/dist/css/bootstrap.min.css';
import "react-widgets/styles.css";
import {
  BrowserRouter as Router,
  Routes,
  Route
} from "react-router-dom";
import Live from './Components/Live';
import Logs from './Components/Logs';
function App() {
  return (
    <Router>
    <div >
    <Routes>
      <Route path="/" element={<NavBar/>} />
      <Route path="/Live" element={<Live/>} />
      <Route path="/Logs" element={<Logs/>} />
    </Routes>
    </div>
    </Router>
  );
}

export default App;
