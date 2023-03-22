import { useEffect, useState } from 'preact/hooks';
import axios from 'axios';
import { getTodos } from './scripts/api';

export function App() {
  const [inputs, setInputs] = useState<any>({});

  const handleChange = (event: any) => {
    const name = event.target.name;
    const value = event.target.value;
    setInputs((values: any) => ({ ...values, [name]: value }));
  };

  const handleSubmit = (event: any) => {
    event.preventDefault();

    getTodos(inputs);
  };

  return (
    <>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          name="title"
          value={inputs.title || ''}
          onChange={handleChange}
          placeholder={'title'}
        />
        <input
          type="text"
          name="description"
          value={inputs.description || ''}
          onChange={handleChange}
          placeholder={'description'}
        />
        <input type="submit" />
      </form>
    </>
  );
}
