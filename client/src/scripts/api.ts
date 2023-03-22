import axios from "axios";


export async function getTodos(e: any): Promise<void> {
    axios
        .post('/api/create', {
            title: e.title,
            description: e.description,
        })
        .then(function (response) {
            console.log(response);
        })
        .catch(function (error) {
            console.log(error);
        });
}