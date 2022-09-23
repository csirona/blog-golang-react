import logo from './logo.svg';
import './App.css';
import {NavBar} from "./components/NavBar"
import {Hero} from "./components/Hero"
import {Footer} from "./components/Footer"

var myHeaders = new Headers();
myHeaders.append("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiIxIiwiZXhwIjoxNjY0NTA1NTU2fQ.6tUzet29faLv5Rg2yPxw3M6JB7z3ZEcRwPwcPrw-RqQ");
myHeaders.append("Content-Type", "application/json");


var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  //body: raw,
  redirect: 'follow'
};

fetch("http://localhost:8000/api/signin", requestOptions)
  .then(response => response.text())
  .then(result => console.log(result))
  .catch(error => console.log('error', error));

function App() {
  return (
    <>
    <NavBar />
    <Hero />
    <Footer />
    </>
  );
}

export default App;
