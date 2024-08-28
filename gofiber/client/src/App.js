import { Container, Row, Col} from "react-bootstrap";

import axios from "axios";
import { useState, useEffect } from "react";

import "./App.css";

function App() {
  const [apiData, setApiData] = useState(false)
  useEffect(() => {
    const fetchData = async () => {
      try {
        const apiUrl = process.env.REACT_APP_API_ROOT;
        const response = await axios.get(apiUrl);

        if(response.status === 200){
          if(response?.data.statusText === "Ok"){
            setApiData(response?.data?.blog_records);
          }
        }

      } catch (error) {
        console.log(error.response)
      }
    };

    fetchData();
    return () => {
      
    }
  }, []);
  
  console.log(apiData);
  return (
    <Container>
      <Row>
        <Col XS="12" clasName="py-2">
        <h1 className="text-center mt-5">React App with Go lang</h1>
        </Col>

          {apiData && (
            apiData.map((record, index) => (
              <Col key={index} XS="4" className="p-3 box">
                <div className="title">{ record.title }</div>
                <div>{ record.post }</div>
              </Col>
            )
            ))
          }

      </Row>
    </Container>
  );
}

export default App;
