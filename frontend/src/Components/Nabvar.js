import Container from 'react-bootstrap/Container';
import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import quiniela from '../images/futbol.png';
function NavBar() {
    return (
        <Navbar bg="dark" variant="dark">
        <Container>
          <Navbar.Brand href="/">
            <img
            src={quiniela}
              alt=" "
              width="60"
              height="60"
            />{' '}
            Usactar World Cup 2022
          </Navbar.Brand>
          <Nav className="me-auto">
              <Nav.Link href="/Live">🚨LIVE 🚨</Nav.Link>
              <Nav.Link href="/Logs">🗄️LOGS🗄️</Nav.Link>
            </Nav>
        </Container>
      </Navbar>
    );
  }

  export default NavBar;