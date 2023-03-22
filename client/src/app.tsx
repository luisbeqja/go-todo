import { useEffect, useState } from 'preact/hooks';
import axios from 'axios';

export function App() {
  const [count, setCount] = useState(0);

  useEffect(() => {
/*     axios
      .post('/api/create', {
        title: 'test',
        description: 'description',
      })
      .then(function (response) {
        console.log(response);
      })
      .catch(function (error) {
        console.log(error);
      }); */
  }, []);

  return (
    <>
      <p class="read-the-docs">{count}</p>
    </>
  );
}
